package ir

import "strings"

// NativeConfigRequirement classifies an EventType by how it treats
// NativeConfigVersion (NCV) in RunScope contexts.
//
// Source of truth: ir-event-log.md §3.0.1 / §3.0.2,
// ir-schema-versioning.md §1.5.3.1 / §1.5.3.2.
type NativeConfigRequirement int

const (
	// NativeConfigForbidden — this EventType is never RunScope-typical;
	// NCV must not be present regardless of context. (For non-RunScope
	// events the scope rule independently forbids NCV; this classification
	// is additional defense-in-depth and a taxonomy completeness marker.)
	NativeConfigForbidden NativeConfigRequirement = iota

	// NativeConfigOptionalPreExecute — RunScope event that occurs before
	// native config materialization; NCV may legitimately be absent.
	// Envelope alone can validate.
	NativeConfigOptionalPreExecute

	// NativeConfigRequired — RunScope event that occurs after native
	// config has been injected; NCV must be present. Envelope alone can validate.
	NativeConfigRequired

	// NativeConfigPhaseDependent — RunScope event that may occur in
	// either pre-execute OR execution-bound phase (e.g., TaskFailed can
	// fail during Preparing OR during Validating). Envelope alone cannot
	// decide; callers should use ValidateEnvelopeWithRunContext.
	NativeConfigPhaseDependent
)

// preExecuteOnlyEvents — strictly pre-execute RunScope.
var preExecuteOnlyEvents = map[EventType]struct{}{
	EventTaskClaimed:        {},
	EventWorkdirPreparing:   {},
	EventWorkdirCreated:     {},
	EventRuntimePinned:      {},
	EventRuntimeHandshakeOK: {},
}

// phaseDependentEvents — may occur either pre-execute OR execution-bound.
// NCV requirement depends on whether the run has crossed NativeConfigInjected.
var phaseDependentEvents = map[EventType]struct{}{
	EventBlockerRaised:          {},
	EventBlockerResolved:        {},
	EventBlockerResolvedRequeue: {},
	EventTaskFailed:             {},
	EventTaskCancelled:          {},
	EventTaskTimedOut:           {},
	EventRuntimePinViolated:     {},
	EventReworkAccepted:         {},
}

// nonRunScopeEvents — EventTypes that should never be RunScope; NCV always forbidden.
// (The scope rule independently forbids NCV in non-RunScope events; this is
// additional defense-in-depth.)
var nonRunScopeEvents = map[EventType]struct{}{
	// Cat A (TaskScope-only)
	EventTaskCreated: {},
	EventTaskQueued:  {},
	// Cat B (RuntimeScope)
	EventRuntimeRegistered:         {},
	EventRuntimeRejected:           {},
	EventRuntimeFingerprintChanged: {},
	EventCapabilityReevaluated:     {},
	EventLeaseInvalidated:          {},
	// Cat F (SystemScope-typical)
	EventPolicyBundleLoaded:   {},
	EventPolicyBundleSwitched: {},
	// Cat G (System/Runtime-scope upgrade signals)
	EventUpgradeDetected:          {},
	EventUpgradePolicyReevaluated: {},
	EventDrainStarted:             {},
	EventDrainTimedOut:            {},
}

// NativeConfigRequirementOf returns the static classification for an EventType.
// Default (unclassified) is NativeConfigRequired — execution-bound RunScope.
func NativeConfigRequirementOf(t EventType) NativeConfigRequirement {
	if _, ok := preExecuteOnlyEvents[t]; ok {
		return NativeConfigOptionalPreExecute
	}
	if _, ok := phaseDependentEvents[t]; ok {
		return NativeConfigPhaseDependent
	}
	if _, ok := nonRunScopeEvents[t]; ok {
		return NativeConfigForbidden
	}
	return NativeConfigRequired
}

// RunContext carries the dynamic facts needed to resolve a PhaseDependent
// EventType's NCV requirement. Owned by the orchestrator / EventIngestor /
// reducer — see ir-schema-versioning.md §1.5.3.2 / ir-event-log.md §3.0.2.
type RunContext struct {
	// NativeConfigEstablished is true once a NativeConfigInjected event
	// (or any later ExecutionBoundOnly event) has been appended for the
	// same RunID. Once true, PhaseDependent events for that run require NCV.
	NativeConfigEstablished bool
}

// EnvelopeViolation describes a single rule violation found by ValidateEnvelope.
type EnvelopeViolation struct {
	Code   string // "MISSING_FIELD" | "FORBIDDEN_FIELD" | "FAKE_PLACEHOLDER" | "UNKNOWN_SCOPE" | "INVALID_FSMVERSION"
	Field  string
	Detail string
}

// fakePlaceholders are sentinel strings that MUST NOT be used to fill an
// identifier field that is conceptually absent.
//
// See ir-schema-versioning.md §1.5.5: when an identifier is not yet decided
// (e.g., RuntimeID before claim), the correct behavior is to emit the event
// under a lower EventScope where that field is legitimately absent, NOT to
// stuff a sentinel value.
var fakePlaceholders = map[string]struct{}{
	"unknown": {},
	"none":    {},
	"pending": {},
	"n/a":     {},
	"na":      {},
	"tbd":     {},
	"-":       {},
}

func isFakePlaceholder(v string) bool {
	if v == "" {
		return false // empty means absent — that's fine; scope rules handle it
	}
	_, ok := fakePlaceholders[strings.ToLower(strings.TrimSpace(v))]
	return ok
}

// ValidateEnvelope checks that a CanonicalEvent satisfies the scope-aware
// envelope rules from ir-schema-versioning.md §1.5.
//
// Returns an empty slice if the event is valid. Non-empty slice means the
// event MUST be rejected at the ingest boundary.
//
// SKELETON: this is intentionally minimal — production ingest will likely
// extend it with EventType-specific Payload schema validation, schemaVersion
// dispatch, and inter-event ordering checks.
func ValidateEnvelope(e CanonicalEvent) []EnvelopeViolation {
	var v []EnvelopeViolation

	// Scope must be one of the declared values.
	if !e.Scope.IsValid() {
		v = append(v, EnvelopeViolation{
			Code:   "UNKNOWN_SCOPE",
			Field:  "Scope",
			Detail: string(e.Scope),
		})
		// without a valid scope we cannot run the per-scope rules
		return v
	}

	// Common envelope — required for every scope.
	v = append(v, requireNonEmpty("EventID", e.EventID)...)
	if e.OccurredAt.IsZero() {
		v = append(v, EnvelopeViolation{Code: "MISSING_FIELD", Field: "OccurredAt"})
	}
	if e.EventSchemaVersion <= 0 {
		v = append(v, EnvelopeViolation{Code: "MISSING_FIELD", Field: "EventSchemaVersion", Detail: "must be >= 1"})
	}
	v = append(v, requireNonEmpty("Type", string(e.Type))...)
	v = append(v, requireNonEmpty("ActorKind", string(e.ActorKind))...)
	// ActorID may be empty for ActorSystem; otherwise we require it.
	if e.ActorKind != ActorSystem && e.ActorID == "" {
		v = append(v, EnvelopeViolation{Code: "MISSING_FIELD", Field: "ActorID", Detail: "required unless ActorKind=system"})
	}
	v = append(v, requireNonEmpty("RiidoDaemonVersion", e.RiidoDaemonVersion)...)
	v = append(v, requireNonEmpty("PolicyBundleVersion", e.PolicyBundleVersion)...)

	// Fake-placeholder ban on identity fields, regardless of scope.
	for _, f := range []struct{ name, val string }{
		{"RuntimeID", e.RuntimeID},
		{"CapabilityFingerprint", e.CapabilityFingerprint},
		{"NativeConfigVersion", e.NativeConfigVersion},
		{"ProviderVersion", e.ProviderVersion},
		{"AdapterID", e.AdapterID},
		{"AdapterVersion", e.AdapterVersion},
		{"ProtocolVersion", e.ProtocolVersion},
		{"TaskID", e.TaskID},
		{"RunID", e.RunID},
	} {
		if isFakePlaceholder(f.val) {
			v = append(v, EnvelopeViolation{
				Code:   "FAKE_PLACEHOLDER",
				Field:  f.name,
				Detail: f.val,
			})
		}
	}

	// Per-scope rules.
	switch e.Scope {
	case EventScopeSystem:
		v = append(v, forbid(e, []string{
			"TaskID", "RunID", "RuntimeID", "CapabilityFingerprint",
			"ProviderKind", "ProtocolKind", "ProviderVersion",
			"AdapterID", "AdapterVersion", "ProtocolVersion",
			"NativeConfigVersion",
		})...)
		if e.FSMVersion != 0 {
			v = append(v, EnvelopeViolation{Code: "FORBIDDEN_FIELD", Field: "FSMVersion", Detail: "FSMVersion must be 0 for SystemScope"})
		}
	case EventScopeRuntime:
		v = append(v, requireNonEmpty("RuntimeID", e.RuntimeID)...)
		v = append(v, forbid(e, []string{
			"TaskID", "RunID", "NativeConfigVersion",
		})...)
		if e.FSMVersion != 0 {
			v = append(v, EnvelopeViolation{Code: "FORBIDDEN_FIELD", Field: "FSMVersion", Detail: "FSMVersion must be 0 for RuntimeScope"})
		}
	case EventScopeTask:
		v = append(v, requireNonEmpty("TaskID", e.TaskID)...)
		v = append(v, forbid(e, []string{
			"RunID", "RuntimeID", "CapabilityFingerprint",
			"ProviderKind", "ProtocolKind", "ProviderVersion",
			"AdapterID", "AdapterVersion", "ProtocolVersion",
			"NativeConfigVersion",
		})...)
		if e.Type.IsTransition() && e.FSMVersion <= 0 {
			v = append(v, EnvelopeViolation{
				Code:   "INVALID_FSMVERSION",
				Field:  "FSMVersion",
				Detail: "TaskScope transition events require FSMVersion >= 1",
			})
		}
	case EventScopeRun:
		// Tier 1: run identity (always required for RunScope).
		// Tier 2: runtime capability (always required for RunScope).
		for _, f := range []struct{ name, val string }{
			{"TaskID", e.TaskID},
			{"RunID", e.RunID},
			{"RuntimeID", e.RuntimeID},
			{"CapabilityFingerprint", e.CapabilityFingerprint},
			{"ProviderKind", e.ProviderKind},
			{"ProtocolKind", e.ProtocolKind},
			{"ProviderVersion", e.ProviderVersion},
			{"AdapterID", e.AdapterID},
			{"AdapterVersion", e.AdapterVersion},
			{"ProtocolVersion", e.ProtocolVersion},
		} {
			if f.val == "" {
				v = append(v, EnvelopeViolation{Code: "MISSING_FIELD", Field: f.name, Detail: "required for RunScope"})
			}
		}
		// Tier 3: execution context. NCV requirement is classified by
		// NativeConfigRequirementOf — envelope-alone handles
		// Forbidden / OptionalPreExecute / Required. PhaseDependent is
		// deferred to ValidateEnvelopeWithRunContext.
		switch NativeConfigRequirementOf(e.Type) {
		case NativeConfigForbidden:
			v = append(v, EnvelopeViolation{
				Code:   "FORBIDDEN_FIELD",
				Field:  "Scope",
				Detail: "EventType " + string(e.Type) + " is not allowed in RunScope",
			})
		case NativeConfigRequired:
			if e.NativeConfigVersion == "" {
				v = append(v, EnvelopeViolation{
					Code:   "MISSING_FIELD",
					Field:  "NativeConfigVersion",
					Detail: "required for execution-bound RunScope event " + string(e.Type),
				})
			}
		case NativeConfigOptionalPreExecute:
			// NCV may be empty or present — both envelope-valid.
		case NativeConfigPhaseDependent:
			// Envelope alone cannot decide. The orchestrator must call
			// ValidateEnvelopeWithRunContext with the run's phase fact.
		}
		if e.Type.IsTransition() && e.FSMVersion <= 0 {
			v = append(v, EnvelopeViolation{
				Code:   "INVALID_FSMVERSION",
				Field:  "FSMVersion",
				Detail: "RunScope transition events require FSMVersion >= 1",
			})
		}
	}

	return v
}

func requireNonEmpty(field, val string) []EnvelopeViolation {
	if val == "" {
		return []EnvelopeViolation{{Code: "MISSING_FIELD", Field: field}}
	}
	return nil
}

func forbid(e CanonicalEvent, fields []string) []EnvelopeViolation {
	var v []EnvelopeViolation
	for _, name := range fields {
		val := fieldValue(e, name)
		if val != "" {
			v = append(v, EnvelopeViolation{
				Code:   "FORBIDDEN_FIELD",
				Field:  name,
				Detail: "must be absent for this scope",
			})
		}
	}
	return v
}

// ValidateEnvelopeWithRunContext is the dynamic-context companion to
// ValidateEnvelope. It performs the same static checks first, then resolves
// PhaseDependent NativeConfigVersion requirements using the provided run
// context.
//
// Callers: EventIngestor / FSM Orchestrator / reducer — any layer that has
// the run's IR log accessible and can compute whether NativeConfigInjected
// (or a later ExecutionBoundOnly event) has been appended for the same RunID.
//
// Source of truth: ir-schema-versioning.md §1.5.3.2, ir-event-log.md §3.0.2.
func ValidateEnvelopeWithRunContext(e CanonicalEvent, ctx RunContext) []EnvelopeViolation {
	v := ValidateEnvelope(e)
	if e.Scope != EventScopeRun {
		return v
	}
	if NativeConfigRequirementOf(e.Type) != NativeConfigPhaseDependent {
		return v
	}
	if ctx.NativeConfigEstablished && e.NativeConfigVersion == "" {
		v = append(v, EnvelopeViolation{
			Code:   "MISSING_FIELD",
			Field:  "NativeConfigVersion",
			Detail: "PhaseDependent event " + string(e.Type) + " requires NCV after run has crossed NativeConfigInjected",
		})
	}
	return v
}

func fieldValue(e CanonicalEvent, name string) string {
	switch name {
	case "TaskID":
		return e.TaskID
	case "RunID":
		return e.RunID
	case "RuntimeID":
		return e.RuntimeID
	case "CapabilityFingerprint":
		return e.CapabilityFingerprint
	case "ProviderKind":
		return e.ProviderKind
	case "ProtocolKind":
		return e.ProtocolKind
	case "ProviderVersion":
		return e.ProviderVersion
	case "AdapterID":
		return e.AdapterID
	case "AdapterVersion":
		return e.AdapterVersion
	case "ProtocolVersion":
		return e.ProtocolVersion
	case "NativeConfigVersion":
		return e.NativeConfigVersion
	}
	return ""
}
