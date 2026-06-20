package task

import "github.com/teamswyg/riido-contracts/ir"

// Transition encodes one legal (from, to, trigger) triple from the generated
// Task Lifecycle FSM.
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

// LegalTransitions returns a defensive copy of the matrix.
func LegalTransitions() []Transition {
	codes := LegalTransitionCodes()
	out := make([]Transition, len(codes))
	for index, code := range codes {
		out[index] = Transition{
			From:    code.From.TaskState(),
			To:      code.To.TaskState(),
			Trigger: code.Trigger.EventType(),
		}
	}
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
	return ValidateTransitionCode(from.Code(), to.Code(), trigger.Code())
}
