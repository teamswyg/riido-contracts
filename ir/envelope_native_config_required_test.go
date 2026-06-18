package ir

import "testing"

func TestValidateEnvelope_ExecutionBoundOnlyRequiresNCV(t *testing.T) {
	for _, evt := range executionBoundNativeConfigEvents() {
		t.Run(string(evt), func(t *testing.T) {
			ev := transitionEvent(validRunScopeEvent(), evt)
			ev.NativeConfigVersion = ""
			v := ValidateEnvelope(ev)
			if !hasViolation(v, "MISSING_FIELD", "NativeConfigVersion") {
				t.Fatalf("%s must require NCV. Violations: %+v", evt, v)
			}
		})
	}
}

func TestValidateEnvelope_TextDeltaAlwaysRequiresNCV(t *testing.T) {
	ev := validRunScopeEvent()
	ev.Type = EventTextDelta
	ev.NativeConfigVersion = ""
	v := ValidateEnvelope(ev)
	if !hasViolation(v, "MISSING_FIELD", "NativeConfigVersion") {
		t.Fatalf("TextDelta must always require NCV. Violations: %+v", v)
	}
}

func executionBoundNativeConfigEvents() []EventType {
	return []EventType{
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
	}
}
