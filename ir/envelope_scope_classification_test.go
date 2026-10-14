package ir

import "testing"

func TestValidateEnvelope_NonRunScopeEventInRunScopeIsForbidden(t *testing.T) {
	ev := transitionEvent(validRunScopeEvent(), EventTaskQueued)
	v := ValidateEnvelope(ev)
	if !hasViolation(v, "FORBIDDEN_FIELD", "Scope") {
		t.Fatalf("TaskQueued in RunScope should be flagged. Violations: %+v", v)
	}
}
