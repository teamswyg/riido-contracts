package ir

import "testing"

func TestValidateEnvelope_RunScopeTransitionRequiresFSMVersion(t *testing.T) {
	e := validRunScopeEvent()
	e.Type = EventTaskClaimed
	e.FSMVersion = 0
	v := ValidateEnvelope(e)
	if !hasViolation(v, "INVALID_FSMVERSION", "FSMVersion") {
		t.Fatalf("expected INVALID_FSMVERSION for transition in RunScope, got %+v", v)
	}
}

func TestValidateEnvelope_NonTransitionAllowsZeroFSMVersion(t *testing.T) {
	e := validRunScopeEvent()
	e.FSMVersion = 0
	assertNoEnvelopeViolations(t, e)
}
