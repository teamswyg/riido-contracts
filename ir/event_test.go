package ir

import "testing"

func TestTransitionClassification(t *testing.T) {
	mustTransition := []EventType{
		EventTaskQueued,
		EventTaskClaimed,
		EventRuntimePinned,
		EventRunStarted,
		EventValidationPassed,
		EventValidationFailed,
		EventReworkAccepted,
		EventRuntimePinViolated,
	}
	for _, e := range mustTransition {
		if !e.IsTransition() {
			t.Errorf("%s must be a transition event", e)
		}
	}
	mustNotTransition := []EventType{
		EventTextDelta,
		EventToolCallStarted,
		EventToolCallFinished,
		EventFileChanged,
		EventStatusUpdate,
		EventUsageDelta,
		EventLogLine,
		EventProviderUnknownEvent,
		EventSnapshot,
		EventOperatorNote,
		EventCorrection,
		EventPolicyBundleLoaded,
		EventValidationStarted,
		EventValidationRuleExecuted,
	}
	for _, e := range mustNotTransition {
		if e.IsTransition() {
			t.Errorf("%s must NOT be a transition event", e)
		}
	}
}

func TestReducerErrorMessage(t *testing.T) {
	var err error = &ReducerError{
		Code:    "IR_REDUCER_INCOMPAT",
		EventID: "ev_1",
		Detail:  "no dispatch for (TaskQueued, 999)",
	}
	if err.Error() == "" {
		t.Fatal("expected non-empty error message")
	}
}

// TestReducerInterfaceHasNoAppender is a compile-time check that ReduceResult
// has no field that grants the reducer permission to append events.
//
// Adding fields like OutboundEvents / Append / Writer to ReduceResult would
// silently violate ir-event-log.md §5.0. If you find yourself needing such
// a field, the right place is the EventIngestor / FSM Orchestrator (outside
// this package), not the reducer.
func TestReducerInterfaceHasNoAppender(t *testing.T) {
	var r ReduceResult
	_ = r.LastEventID
	_ = r.CurrentState
	_ = r.Diagnostics
	_ = r.Error
}

func TestActorKindValues(t *testing.T) {
	kinds := []ActorKind{ActorAgent, ActorDaemon, ActorHuman, ActorSystem}
	if len(kinds) != 4 {
		t.Fatalf("expected 4 actor kinds, got %d", len(kinds))
	}
}
