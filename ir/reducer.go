package ir

// Reducer is the pure-function contract over an IR event log.
//
// PURITY INVARIANT from the generated IR Event Log reader:
// Reducer implementations MUST:
//   - read events only;
//   - perform NO I/O;
//   - have NO writer / appender / event-log mutation dependency;
//   - return only data (ReduceResult).
//
// Decisions like "this incompatible event should trigger BlockerRaised" are
// made OUTSIDE the reducer by an EventIngestor / FSM Orchestrator that
// inspects ReduceResult.Error and appends new events under server actor
// attribution (during live ingest only — never during replay).
type Reducer interface {
	Reduce(events []CanonicalEvent) ReduceResult
}

// ReduceResult is the pure output of a Reducer.
//
// It never contains "events to append" — that authority lives outside
// the reducer (see PURITY INVARIANT above).
type ReduceResult struct {
	// LastEventID is the EventID of the last successfully reduced event.
	// Empty when zero events have been processed.
	LastEventID string

	// CurrentState is the derived TaskState as a string.
	// Kept as string to avoid an ir → task package cycle; the task package
	// (or orchestrator) converts back to its TaskState type at the call site.
	CurrentState string

	// Diagnostics are non-fatal observations the orchestrator may forward
	// to operators or telemetry. The reducer remains pure even when these
	// are non-empty.
	Diagnostics []ReducerDiagnostic

	// Error is non-nil when the reducer cannot proceed (unknown
	// (EventType, EventSchemaVersion) pair, out-of-order transition, etc.).
	// In REPLAY contexts: orchestrator MUST NOT mutate the log.
	// In LIVE INGEST contexts: orchestrator MAY append a BlockerRaised
	// event under server actor attribution.
	Error *ReducerError
}

// ReducerError is a fatal reducer outcome. It is data, not a side-effect.
//
// Defined codes:
//
//	IR_REDUCER_INCOMPAT     — unknown (EventType, EventSchemaVersion) pair
//	OUT_OF_ORDER_TRANSITION — transition sequence violates monotonicity
//	INVARIANT_VIOLATION     — generic invariant breach
type ReducerError struct {
	Code    string
	EventID string
	Detail  string
}

func (e *ReducerError) Error() string {
	if e == nil {
		return ""
	}
	return "ir reducer: " + e.Code + " on event " + e.EventID + ": " + e.Detail
}

// ReducerDiagnostic is a non-fatal observation surfaced by a reducer.
type ReducerDiagnostic struct {
	Code    string
	EventID string
	Detail  string
}
