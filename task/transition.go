package task

import "github.com/teamswyg/riido-contracts/ir"

// Transition encodes one legal (from, to, trigger) triple from the matrix in
// docs/20-domain/task-lifecycle.md §3 + §4.
//
// The matrix only checks "is this state-to-state move structurally legal".
// It does NOT enforce the contextual invariants 2/3/4/6 (which need the full
// event log to evaluate: RuntimePinned ordering, validation precondition,
// new RunID issuance, etc.). Those checks live in the orchestrator layer.
type Transition struct {
	From    TaskState
	To      TaskState
	Trigger ir.EventType
}

// legalTransitions is the SSOT for state-to-state move legality.
// Source of truth: task-lifecycle.md §3 matrix + §4 trigger reference table.
var legalTransitions = []Transition{
	// happy path
	{StateCreated, StateQueued, ir.EventTaskQueued},
	{StateQueued, StateClaimed, ir.EventTaskClaimed},
	{StateClaimed, StatePreparing, ir.EventWorkdirPreparing},
	{StatePreparing, StateRunning, ir.EventRunStarted},
	{StateRunning, StateValidating, ir.EventRunReportedDone},
	{StateValidating, StatePatchReady, ir.EventValidationPassed},
	{StatePatchReady, StateHumanReview, ir.EventReviewRequested},
	{StateHumanReview, StateCompleted, ir.EventHumanApproved},
	// auto-approve branch
	{StatePatchReady, StateCompleted, ir.EventAutoApproved},
	// NeedsInput round-trip
	{StateRunning, StateNeedsInput, ir.EventInputRequested},
	{StateNeedsInput, StateRunning, ir.EventInputProvided},
	// Blocker round-trip
	{StatePreparing, StateBlocked, ir.EventBlockerRaised},
	{StateRunning, StateBlocked, ir.EventBlockerRaised},
	{StateBlocked, StateRunning, ir.EventBlockerResolved},
	{StateBlocked, StateQueued, ir.EventBlockerResolvedRequeue},
	// validation failure
	{StateValidating, StateFailed, ir.EventValidationFailed},
	// human rejection
	{StateHumanReview, StateReworkQueued, ir.EventHumanRejected}, // rework=true
	{StateHumanReview, StateCancelled, ir.EventHumanRejected},    // rework=false
	// rework re-entry (new RunID issued — invariant 6)
	{StateReworkQueued, StateQueued, ir.EventReworkAccepted},
	// runtime pin violation (invariant 3)
	{StateRunning, StateFailed, ir.EventRuntimePinViolated},
	{StateValidating, StateFailed, ir.EventRuntimePinViolated},
	// generic TaskFailed (broader sources)
	{StateClaimed, StateFailed, ir.EventTaskFailed},
	{StatePreparing, StateFailed, ir.EventTaskFailed},
	{StateRunning, StateFailed, ir.EventTaskFailed},
	{StateNeedsInput, StateFailed, ir.EventTaskFailed},
	{StateBlocked, StateFailed, ir.EventTaskFailed},
	{StateValidating, StateFailed, ir.EventTaskFailed},
	// TaskCancelled from non-terminal states (task-lifecycle.md §3.2)
	{StateCreated, StateCancelled, ir.EventTaskCancelled},
	{StateQueued, StateCancelled, ir.EventTaskCancelled},
	{StateClaimed, StateCancelled, ir.EventTaskCancelled},
	{StatePreparing, StateCancelled, ir.EventTaskCancelled},
	{StateRunning, StateCancelled, ir.EventTaskCancelled},
	{StateNeedsInput, StateCancelled, ir.EventTaskCancelled},
	{StateBlocked, StateCancelled, ir.EventTaskCancelled},
	{StateValidating, StateCancelled, ir.EventTaskCancelled},
	{StatePatchReady, StateCancelled, ir.EventTaskCancelled},
	{StateHumanReview, StateCancelled, ir.EventTaskCancelled},
	{StateReworkQueued, StateCancelled, ir.EventTaskCancelled},
	// TaskTimedOut from active subset (task-lifecycle.md §3.3)
	{StateRunning, StateTimedOut, ir.EventTaskTimedOut},
	{StateNeedsInput, StateTimedOut, ir.EventTaskTimedOut},
	{StateBlocked, StateTimedOut, ir.EventTaskTimedOut},
	{StateValidating, StateTimedOut, ir.EventTaskTimedOut},
	{StateHumanReview, StateTimedOut, ir.EventTaskTimedOut},
}

// LegalTransitions returns a defensive copy of the matrix.
func LegalTransitions() []Transition {
	out := make([]Transition, len(legalTransitions))
	copy(out, legalTransitions)
	return out
}

// ValidateTransition reports whether (from, to) under the given trigger is
// permitted by the matrix.
//
// Limitations (intentionally not enforced here):
//   - Invariant 2 (RuntimePinned must precede RunStarted) needs event-log context.
//   - Invariant 3 (pin violation → Failed) needs the actual fingerprints.
//   - Invariant 4 (Completed requires validation + approval) needs prior events.
//   - Invariant 6 (ReworkQueued issues a new RunID) needs RunID generation context.
//
// Those checks live in the orchestrator that calls into both ir (for events)
// and this package (for matrix lookup).
func ValidateTransition(from, to TaskState, trigger ir.EventType) bool {
	for _, t := range legalTransitions {
		if t.From == from && t.To == to && t.Trigger == trigger {
			return true
		}
	}
	return false
}
