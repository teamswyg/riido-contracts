package ir

import "testing"

func TestTransitionClassification(t *testing.T) {
	for _, event := range transitionEvents() {
		if !event.IsTransition() {
			t.Errorf("%s must be a transition event", event)
		}
	}
	for _, event := range nonTransitionEvents() {
		if event.IsTransition() {
			t.Errorf("%s must NOT be a transition event", event)
		}
	}
}

func transitionEvents() []EventType {
	return []EventType{
		EventTaskQueued, EventTaskClaimed, EventRuntimePinned, EventRunStarted,
		EventValidationPassed, EventValidationFailed, EventReworkAccepted,
		EventRuntimePinViolated,
	}
}

func nonTransitionEvents() []EventType {
	return []EventType{
		EventTextDelta, EventToolCallStarted, EventToolCallFinished,
		EventFileChanged, EventStatusUpdate, EventUsageDelta, EventLogLine,
		EventProviderUnknownEvent, EventSnapshot, EventOperatorNote,
		EventCorrection, EventPolicyBundleLoaded, EventValidationStarted,
		EventValidationRuleExecuted,
	}
}
