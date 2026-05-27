package ir

import "time"

// CanonicalEvent is the append-only IR record.
//
// Mandatory field set is SCOPE-DEPENDENT per docs/20-domain/ir-schema-versioning.md §1.5:
// the four EventScopes (System/Runtime/Task/Run) each have different field
// requirements. The struct lists every possible field; ValidateEnvelope (see
// envelope.go) enforces the scope-specific present/absent rules.
//
// Common envelope fields (required for every scope):
//
//	EventID, OccurredAt, EventSchemaVersion, Scope, Type,
//	ActorKind, ActorID, RiidoDaemonVersion, PolicyBundleVersion,
//	Payload, Unknown.
//
// Scope-additional fields (required only for the indicated scope):
//
//	SystemScope:  none.
//	RuntimeScope: RuntimeID (CapabilityFingerprint allowed once capability detected).
//	TaskScope:    TaskID.
//	RunScope:     TaskID, RunID, RuntimeID, CapabilityFingerprint,
//	              ProviderKind, ProtocolKind, ProviderVersion, AdapterID,
//	              AdapterVersion, ProtocolVersion, NativeConfigVersion.
//
// FSMVersion is required only for transition events (Type.IsTransition() == true)
// in TaskScope or RunScope.
//
// NOTE: a single ambiguous identifier (merging runtime slot + vendor family +
// protocol selection) is intentionally absent. Use RuntimeID for the runtime
// slot, ProviderKind for the vendor family, and ProtocolKind for the
// adapter/protocol selection — three distinct concepts.
// See docs/20-domain/provider-capability.md §0 invariant 1.
type CanonicalEvent struct {
	// common envelope
	EventID             string
	OccurredAt          time.Time
	EventSchemaVersion  int
	Scope               EventScope
	Type                EventType
	ActorKind           ActorKind // server-decided (ir-event-log.md §9)
	ActorID             string    // server-decided
	RiidoDaemonVersion  string
	PolicyBundleVersion string
	Payload             map[string]any
	Unknown             map[string]any

	// scope-conditional identity fields (presence governed by ValidateEnvelope)
	TaskID                string
	RunID                 string
	RuntimeID             string
	CapabilityFingerprint string

	// kinds (categorical, not semver) — RunScope required, RuntimeScope optional
	ProviderKind string
	ProtocolKind string

	// versions / execution fingerprint — RunScope required
	ProviderVersion     string
	AdapterID           string
	AdapterVersion      string
	ProtocolVersion     string
	NativeConfigVersion string

	// FSM transition only (0 allowed only for non-transition events)
	FSMVersion int
}
