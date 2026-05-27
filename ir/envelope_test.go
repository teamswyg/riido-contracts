package ir

import (
	"testing"
	"time"
)

func validRunScopeEvent() CanonicalEvent {
	return CanonicalEvent{
		EventID:               "ev_1",
		OccurredAt:            time.Date(2026, 5, 19, 12, 0, 0, 0, time.UTC),
		EventSchemaVersion:    1,
		Scope:                 EventScopeRun,
		Type:                  EventTextDelta,
		ActorKind:             ActorAgent,
		ActorID:               "run_1",
		RiidoDaemonVersion:    "0.1.0",
		PolicyBundleVersion:   "pb-1",
		TaskID:                "task_1",
		RunID:                 "run_1",
		RuntimeID:             "rt_1",
		CapabilityFingerprint: "fp_abc",
		ProviderKind:          "claude",
		ProtocolKind:          "claude-stream-json",
		ProviderVersion:       "2.1.128",
		AdapterID:             "claude-stream-json",
		AdapterVersion:        "0.1.0",
		ProtocolVersion:       "stream-json-v1",
		NativeConfigVersion:   "nc_xyz",
	}
}

func TestValidateEnvelope_RunScopeValid(t *testing.T) {
	e := validRunScopeEvent()
	if v := ValidateEnvelope(e); len(v) != 0 {
		t.Fatalf("valid RunScope event flagged: %+v", v)
	}
}

func TestValidateEnvelope_RunScopeMissingRuntimeID(t *testing.T) {
	e := validRunScopeEvent()
	e.RuntimeID = ""
	v := ValidateEnvelope(e)
	if !hasViolation(v, "MISSING_FIELD", "RuntimeID") {
		t.Fatalf("expected MISSING_FIELD RuntimeID, got %+v", v)
	}
}

func TestValidateEnvelope_TaskScopeNoRuntimeRequired(t *testing.T) {
	e := CanonicalEvent{
		EventID:             "ev_2",
		OccurredAt:          time.Date(2026, 5, 19, 12, 0, 0, 0, time.UTC),
		EventSchemaVersion:  1,
		Scope:               EventScopeTask,
		Type:                EventTaskQueued,
		ActorKind:           ActorDaemon,
		ActorID:             "daemon-1",
		RiidoDaemonVersion:  "0.1.0",
		PolicyBundleVersion: "pb-1",
		TaskID:              "task_1",
		FSMVersion:          1, // TaskQueued is a transition event
	}
	if v := ValidateEnvelope(e); len(v) != 0 {
		t.Fatalf("valid TaskScope event flagged: %+v", v)
	}
}

func TestValidateEnvelope_TaskScopeForbidsRuntimeID(t *testing.T) {
	e := CanonicalEvent{
		EventID:             "ev_3",
		OccurredAt:          time.Date(2026, 5, 19, 12, 0, 0, 0, time.UTC),
		EventSchemaVersion:  1,
		Scope:               EventScopeTask,
		Type:                EventTaskCreated,
		ActorKind:           ActorDaemon,
		ActorID:             "daemon-1",
		RiidoDaemonVersion:  "0.1.0",
		PolicyBundleVersion: "pb-1",
		TaskID:              "task_1",
		RuntimeID:           "rt_X", // forbidden in TaskScope
		FSMVersion:          1,
	}
	v := ValidateEnvelope(e)
	if !hasViolation(v, "FORBIDDEN_FIELD", "RuntimeID") {
		t.Fatalf("expected FORBIDDEN_FIELD RuntimeID for TaskScope, got %+v", v)
	}
}

func TestValidateEnvelope_FakePlaceholderBanned(t *testing.T) {
	for _, sentinel := range []string{"unknown", "UNKNOWN", "none", "pending", "tbd", "n/a", " - "} {
		t.Run(sentinel, func(t *testing.T) {
			e := validRunScopeEvent()
			e.RuntimeID = sentinel
			v := ValidateEnvelope(e)
			if !hasViolation(v, "FAKE_PLACEHOLDER", "RuntimeID") {
				t.Fatalf("expected FAKE_PLACEHOLDER for RuntimeID=%q, got %+v", sentinel, v)
			}
		})
	}
}

func TestValidateEnvelope_UnknownScopeRejected(t *testing.T) {
	e := validRunScopeEvent()
	e.Scope = "bogus-scope"
	v := ValidateEnvelope(e)
	if !hasViolation(v, "UNKNOWN_SCOPE", "Scope") {
		t.Fatalf("expected UNKNOWN_SCOPE, got %+v", v)
	}
}

func TestValidateEnvelope_SystemScopeNoTask(t *testing.T) {
	e := CanonicalEvent{
		EventID:             "ev_sys",
		OccurredAt:          time.Date(2026, 5, 19, 12, 0, 0, 0, time.UTC),
		EventSchemaVersion:  1,
		Scope:               EventScopeSystem,
		Type:                EventPolicyBundleLoaded,
		ActorKind:           ActorSystem,
		RiidoDaemonVersion:  "0.1.0",
		PolicyBundleVersion: "pb-2",
	}
	if v := ValidateEnvelope(e); len(v) != 0 {
		t.Fatalf("valid SystemScope event flagged: %+v", v)
	}
}

func TestValidateEnvelope_RunScopeTransitionRequiresFSMVersion(t *testing.T) {
	e := validRunScopeEvent()
	e.Type = EventTaskClaimed // transition event
	e.FSMVersion = 0
	v := ValidateEnvelope(e)
	if !hasViolation(v, "INVALID_FSMVERSION", "FSMVersion") {
		t.Fatalf("expected INVALID_FSMVERSION for transition in RunScope, got %+v", v)
	}
}

func TestValidateEnvelope_NonTransitionAllowsZeroFSMVersion(t *testing.T) {
	e := validRunScopeEvent() // Type=EventTextDelta (non-transition)
	e.FSMVersion = 0
	if v := ValidateEnvelope(e); len(v) != 0 {
		t.Fatalf("non-transition event with FSMVersion=0 should be valid, got %+v", v)
	}
}

func hasViolation(vs []EnvelopeViolation, code, field string) bool {
	for _, v := range vs {
		if v.Code == code && v.Field == field {
			return true
		}
	}
	return false
}

// #27i CapabilityFingerprint / NativeConfigVersion boundary tests.

func TestValidateEnvelope_PreExecuteRunScopeAllowsMissingNCV(t *testing.T) {
	// TaskClaimed is the canonical example: runtime identity has just been
	// pinned, but workspace prep hasn't happened — NCV does not yet exist.
	e := validRunScopeEvent()
	e.Type = EventTaskClaimed
	e.FSMVersion = 1 // transition event
	e.NativeConfigVersion = ""
	if v := ValidateEnvelope(e); len(v) != 0 {
		t.Fatalf("pre-execute RunScope TaskClaimed without NCV should be valid, got %+v", v)
	}

	for _, evt := range []EventType{
		EventWorkdirPreparing,
		EventWorkdirCreated,
		EventRuntimePinned,
		EventRuntimeHandshakeOK,
		EventBlockerRaised,
		EventTaskCancelled,
		EventTaskTimedOut,
		EventReworkAccepted,
		EventTaskFailed,
	} {
		t.Run(string(evt), func(t *testing.T) {
			ev := validRunScopeEvent()
			ev.Type = evt
			ev.NativeConfigVersion = ""
			// Some are transition events: provide FSMVersion to isolate the NCV check.
			if evt.IsTransition() {
				ev.FSMVersion = 1
			}
			v := ValidateEnvelope(ev)
			if hasViolation(v, "MISSING_FIELD", "NativeConfigVersion") {
				t.Fatalf("%s is pre-execute RunScope — must allow empty NCV. Got: %+v", evt, v)
			}
		})
	}
}

// ExecutionBoundOnly events: NCV is required at envelope level.
// RuntimePinViolated is INTENTIONALLY excluded — it's PhaseDependent now (#27j).
func TestValidateEnvelope_ExecutionBoundOnlyRequiresNCV(t *testing.T) {
	for _, evt := range []EventType{
		EventRunStarted,
		EventNativeConfigInjected,
		EventTextDelta,
		EventToolCallStarted,
		EventFileChanged,
		EventValidationStarted,
		EventValidationPassed,
		EventReviewRequested,
		EventHumanApproved,
		EventConfigTemplateReinjected,
	} {
		t.Run(string(evt), func(t *testing.T) {
			ev := validRunScopeEvent()
			ev.Type = evt
			ev.NativeConfigVersion = ""
			if evt.IsTransition() {
				ev.FSMVersion = 1
			}
			v := ValidateEnvelope(ev)
			if !hasViolation(v, "MISSING_FIELD", "NativeConfigVersion") {
				t.Fatalf("%s is ExecutionBoundOnly RunScope — must require NCV. Violations: %+v", evt, v)
			}
		})
	}
}

// PhaseDependent events: envelope-alone allows NCV to be absent OR present.
func TestValidateEnvelope_PhaseDependentEnvelopeAllowsEither(t *testing.T) {
	for _, evt := range []EventType{
		EventTaskFailed,
		EventTaskCancelled,
		EventTaskTimedOut,
		EventBlockerRaised,
		EventBlockerResolved,
		EventBlockerResolvedRequeue,
		EventRuntimePinViolated,
		EventReworkAccepted,
	} {
		t.Run(string(evt)+"/no-ncv", func(t *testing.T) {
			ev := validRunScopeEvent()
			ev.Type = evt
			ev.NativeConfigVersion = ""
			if evt.IsTransition() {
				ev.FSMVersion = 1
			}
			v := ValidateEnvelope(ev)
			if hasViolation(v, "MISSING_FIELD", "NativeConfigVersion") {
				t.Fatalf("PhaseDependent %s without NCV should be envelope-valid. Violations: %+v", evt, v)
			}
		})
		t.Run(string(evt)+"/with-ncv", func(t *testing.T) {
			ev := validRunScopeEvent()
			ev.Type = evt
			ev.NativeConfigVersion = "nc_xyz"
			if evt.IsTransition() {
				ev.FSMVersion = 1
			}
			v := ValidateEnvelope(ev)
			if len(v) != 0 {
				t.Fatalf("PhaseDependent %s with NCV should be envelope-valid. Violations: %+v", evt, v)
			}
		})
	}
}

// Dynamic check via ValidateEnvelopeWithRunContext:
// PhaseDependent NCV becomes required once NativeConfigEstablished=true.
func TestValidateEnvelopeWithRunContext_PhaseDependent(t *testing.T) {
	cases := []struct {
		name                    string
		evt                     EventType
		nativeConfigEstablished bool
		ncvSet                  bool
		wantNCVMissing          bool
	}{
		{"TaskFailed before NCV established", EventTaskFailed, false, false, false},
		{"TaskFailed after NCV established without NCV", EventTaskFailed, true, false, true},
		{"TaskFailed after NCV established with NCV", EventTaskFailed, true, true, false},
		{"RuntimePinViolated before NCV established", EventRuntimePinViolated, false, false, false},
		{"RuntimePinViolated after NCV established without NCV", EventRuntimePinViolated, true, false, true},
		{"BlockerRaised before NCV established", EventBlockerRaised, false, false, false},
		{"BlockerRaised after NCV established without NCV", EventBlockerRaised, true, false, true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ev := validRunScopeEvent()
			ev.Type = c.evt
			if c.evt.IsTransition() {
				ev.FSMVersion = 1
			}
			if !c.ncvSet {
				ev.NativeConfigVersion = ""
			}
			v := ValidateEnvelopeWithRunContext(ev, RunContext{NativeConfigEstablished: c.nativeConfigEstablished})
			got := hasViolation(v, "MISSING_FIELD", "NativeConfigVersion")
			if got != c.wantNCVMissing {
				t.Fatalf("%s: want NCV missing violation=%v, got %v. All violations: %+v", c.name, c.wantNCVMissing, got, v)
			}
		})
	}
}

// Static envelope: TextDelta without NCV is always invalid.
func TestValidateEnvelope_TextDeltaAlwaysRequiresNCV(t *testing.T) {
	ev := validRunScopeEvent()
	ev.Type = EventTextDelta
	ev.NativeConfigVersion = ""
	v := ValidateEnvelope(ev)
	if !hasViolation(v, "MISSING_FIELD", "NativeConfigVersion") {
		t.Fatalf("TextDelta must always require NCV. Violations: %+v", v)
	}
}

// Static envelope: TaskClaimed without NCV is always valid.
func TestValidateEnvelope_TaskClaimedNeverRequiresNCV(t *testing.T) {
	ev := validRunScopeEvent()
	ev.Type = EventTaskClaimed
	ev.FSMVersion = 1
	ev.NativeConfigVersion = ""
	v := ValidateEnvelope(ev)
	if hasViolation(v, "MISSING_FIELD", "NativeConfigVersion") {
		t.Fatalf("TaskClaimed must never require NCV at envelope. Violations: %+v", v)
	}
}

// Non-RunScope events flagged FORBIDDEN if appearing in RunScope.
func TestValidateEnvelope_NonRunScopeEventInRunScopeIsForbidden(t *testing.T) {
	ev := validRunScopeEvent()
	ev.Type = EventTaskQueued // classified NonRunScope
	ev.FSMVersion = 1
	v := ValidateEnvelope(ev)
	if !hasViolation(v, "FORBIDDEN_FIELD", "Scope") {
		t.Fatalf("TaskQueued classified as NonRunScope; RunScope occurrence should be flagged. Violations: %+v", v)
	}
}

// Classifier spot checks.
func TestNativeConfigRequirementOf(t *testing.T) {
	cases := []struct {
		evt  EventType
		want NativeConfigRequirement
	}{
		{EventTaskClaimed, NativeConfigOptionalPreExecute},
		{EventWorkdirPreparing, NativeConfigOptionalPreExecute},
		{EventRuntimePinned, NativeConfigOptionalPreExecute},
		{EventRuntimeHandshakeOK, NativeConfigOptionalPreExecute},

		{EventBlockerRaised, NativeConfigPhaseDependent},
		{EventTaskFailed, NativeConfigPhaseDependent},
		{EventTaskCancelled, NativeConfigPhaseDependent},
		{EventTaskTimedOut, NativeConfigPhaseDependent},
		{EventRuntimePinViolated, NativeConfigPhaseDependent},
		{EventReworkAccepted, NativeConfigPhaseDependent},

		{EventNativeConfigInjected, NativeConfigRequired},
		{EventRunStarted, NativeConfigRequired},
		{EventTextDelta, NativeConfigRequired},
		{EventValidationPassed, NativeConfigRequired},
		{EventHumanApproved, NativeConfigRequired},

		{EventTaskCreated, NativeConfigForbidden},
		{EventTaskQueued, NativeConfigForbidden},
		{EventRuntimeRegistered, NativeConfigForbidden},
		{EventPolicyBundleLoaded, NativeConfigForbidden},
		{EventUpgradeDetected, NativeConfigForbidden},
	}
	for _, c := range cases {
		t.Run(string(c.evt), func(t *testing.T) {
			if got := NativeConfigRequirementOf(c.evt); got != c.want {
				t.Errorf("NativeConfigRequirementOf(%s) = %v, want %v", c.evt, got, c.want)
			}
		})
	}
}
