package ir

import "testing"

// TestReducerInterfaceHasNoAppender is a compile-time check that ReduceResult
// has no field that grants the reducer permission to append events.
//
// Adding fields like OutboundEvents / Append / Writer to ReduceResult would
// silently violate ir-event-log.md §5.0. If you find yourself needing such
// a field, the right place is the EventIngestor / FSM Orchestrator.
func TestReducerInterfaceHasNoAppender(t *testing.T) {
	var r ReduceResult
	_ = r.LastEventID
	_ = r.CurrentState
	_ = r.Diagnostics
	_ = r.Error
}
