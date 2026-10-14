package ir

import "testing"

func TestValidateEnvelope_PreExecuteRunScopeAllowsMissingNCV(t *testing.T) {
	e := transitionEvent(validRunScopeEvent(), EventTaskClaimed)
	e.NativeConfigVersion = ""
	assertNoEnvelopeViolations(t, e)

	for _, evt := range preExecuteNativeConfigEvents() {
		t.Run(string(evt), func(t *testing.T) {
			ev := transitionEvent(validRunScopeEvent(), evt)
			ev.NativeConfigVersion = ""
			v := ValidateEnvelope(ev)
			if hasViolation(v, "MISSING_FIELD", "NativeConfigVersion") {
				t.Fatalf("%s must allow empty NCV. Got: %+v", evt, v)
			}
		})
	}
}

func preExecuteNativeConfigEvents() []EventType {
	return []EventType{
		EventWorkdirPreparing,
		EventWorkdirCreated,
		EventRuntimePinned,
		EventRuntimeHandshakeOK,
		EventBlockerRaised,
		EventTaskCancelled,
		EventTaskTimedOut,
		EventReworkAccepted,
		EventTaskFailed,
	}
}
