package ir

import "testing"

func TestValidateEnvelope_PhaseDependentEnvelopeAllowsEither(t *testing.T) {
	for _, evt := range phaseDependentNativeConfigEvents() {
		t.Run(string(evt)+"/no-ncv", func(t *testing.T) {
			ev := transitionEvent(validRunScopeEvent(), evt)
			ev.NativeConfigVersion = ""
			v := ValidateEnvelope(ev)
			if hasViolation(v, "MISSING_FIELD", "NativeConfigVersion") {
				t.Fatalf("PhaseDependent %s without NCV should be valid. Violations: %+v", evt, v)
			}
		})
		t.Run(string(evt)+"/with-ncv", func(t *testing.T) {
			ev := transitionEvent(validRunScopeEvent(), evt)
			ev.NativeConfigVersion = "nc_xyz"
			assertNoEnvelopeViolations(t, ev)
		})
	}
}

func TestValidateEnvelope_TaskClaimedNeverRequiresNCV(t *testing.T) {
	ev := transitionEvent(validRunScopeEvent(), EventTaskClaimed)
	ev.NativeConfigVersion = ""
	v := ValidateEnvelope(ev)
	if hasViolation(v, "MISSING_FIELD", "NativeConfigVersion") {
		t.Fatalf("TaskClaimed must never require NCV at envelope. Violations: %+v", v)
	}
}

func phaseDependentNativeConfigEvents() []EventType {
	return []EventType{
		EventTaskFailed,
		EventTaskCancelled,
		EventTaskTimedOut,
		EventBlockerRaised,
		EventBlockerResolved,
		EventBlockerResolvedRequeue,
		EventRuntimePinViolated,
		EventReworkAccepted,
	}
}
