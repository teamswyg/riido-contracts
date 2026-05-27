// Package task owns the C1 Task Lifecycle domain: TaskState enum, legal
// transition matrix, terminal definition, and the FSM invariants enumerated
// in docs/20-domain/task-lifecycle.md.
//
// What this package does NOT own:
//   - EventType catalog and payload schema → ir (C2).
//   - "Which event is a transition event" classification → ir.
//   - Reducer behavior / dispatch / unknown-pair handling → ir.
//
// Dependency direction: task → ir (read-only). ir does NOT depend on task.
package task

// FSMSchemaVersion is the current version of the TaskState transition
// matrix owned by docs/20-domain/task-lifecycle.md.
const FSMSchemaVersion = 1

// TaskState is one of the 15 lifecycle states defined in
// docs/20-domain/task-lifecycle.md §2.
type TaskState string

const (
	StateCreated      TaskState = "Created"
	StateQueued       TaskState = "Queued"
	StateClaimed      TaskState = "Claimed"
	StatePreparing    TaskState = "Preparing"
	StateRunning      TaskState = "Running"
	StateNeedsInput   TaskState = "NeedsInput"
	StateBlocked      TaskState = "Blocked"
	StateValidating   TaskState = "Validating"
	StatePatchReady   TaskState = "PatchReady"
	StateHumanReview  TaskState = "HumanReview"
	StateReworkQueued TaskState = "ReworkQueued"
	StateCompleted    TaskState = "Completed"
	StateFailed       TaskState = "Failed"
	StateCancelled    TaskState = "Cancelled"
	StateTimedOut     TaskState = "TimedOut"
)

// AllStates returns the 15 task states in declaration order.
func AllStates() []TaskState {
	return []TaskState{
		StateCreated, StateQueued, StateClaimed, StatePreparing,
		StateRunning, StateNeedsInput, StateBlocked, StateValidating,
		StatePatchReady, StateHumanReview, StateReworkQueued,
		StateCompleted, StateFailed, StateCancelled, StateTimedOut,
	}
}

// IsTerminal reports whether s is one of the four terminal states:
// Completed, Failed, Cancelled, TimedOut.
//
// Invariant 5 (task-lifecycle.md §5): no transition can originate from a
// terminal state.
func (s TaskState) IsTerminal() bool {
	switch s {
	case StateCompleted, StateFailed, StateCancelled, StateTimedOut:
		return true
	default:
		return false
	}
}

// IsActive reports whether s is an active state (Running or Validating).
// During active states, the (RuntimeID, CapabilityFingerprint) pinning
// invariant is enforced — see invariant 3 in task-lifecycle.md §5.
func (s TaskState) IsActive() bool {
	return s == StateRunning || s == StateValidating
}
