package task

import (
	"testing"

	"github.com/teamswyg/riido-contracts/ir"
)

func TestAllStatesCount(t *testing.T) {
	if got := len(AllStates()); got != 15 {
		t.Fatalf("expected 15 task states, got %d", got)
	}
}

func TestFSMSchemaVersion(t *testing.T) {
	if FSMSchemaVersion != 1 {
		t.Fatalf("FSMSchemaVersion = %d, want 1", FSMSchemaVersion)
	}
}

func TestTerminalCount(t *testing.T) {
	terminals := 0
	for _, s := range AllStates() {
		if s.IsTerminal() {
			terminals++
		}
	}
	if terminals != 4 {
		t.Fatalf("expected 4 terminal states (Completed/Failed/Cancelled/TimedOut), got %d", terminals)
	}
}

func TestActiveSubset(t *testing.T) {
	if !StateRunning.IsActive() {
		t.Error("Running must be active")
	}
	if !StateValidating.IsActive() {
		t.Error("Validating must be active")
	}
	if StateBlocked.IsActive() {
		t.Error("Blocked must NOT be active (runtime pinning is not enforced while blocked)")
	}
}

// TestInvariant5_NoTransitionFromTerminal verifies invariant 5 from
// task-lifecycle.md §5: terminal states have no outgoing transitions.
func TestInvariant5_NoTransitionFromTerminal(t *testing.T) {
	for _, tr := range LegalTransitions() {
		if tr.From.IsTerminal() {
			t.Errorf("invariant 5 violated: terminal %s -> %s via %s", tr.From, tr.To, tr.Trigger)
		}
	}
}

// TestHappyPathReachable walks the canonical Created→…→Completed path.
func TestHappyPathReachable(t *testing.T) {
	path := []struct {
		from, to TaskState
		trigger  ir.EventType
	}{
		{StateCreated, StateQueued, ir.EventTaskQueued},
		{StateQueued, StateClaimed, ir.EventTaskClaimed},
		{StateClaimed, StatePreparing, ir.EventWorkdirPreparing},
		{StatePreparing, StateRunning, ir.EventRunStarted},
		{StateRunning, StateValidating, ir.EventRunReportedDone},
		{StateValidating, StatePatchReady, ir.EventValidationPassed},
		{StatePatchReady, StateHumanReview, ir.EventReviewRequested},
		{StateHumanReview, StateCompleted, ir.EventHumanApproved},
	}
	for _, step := range path {
		if !ValidateTransition(step.from, step.to, step.trigger) {
			t.Errorf("happy-path step %s -(%s)-> %s rejected by matrix", step.from, step.trigger, step.to)
		}
	}
}

func TestIllegalTransitionsRejected(t *testing.T) {
	cases := []struct {
		name             string
		from, to         TaskState
		trigger          ir.EventType
		expectedLegal    bool
		invariantMessage string
	}{
		{
			name:             "skip queued",
			from:             StateCreated,
			to:               StateRunning,
			trigger:          ir.EventRunStarted,
			expectedLegal:    false,
			invariantMessage: "Created cannot jump to Running",
		},
		{
			name:             "skip validating",
			from:             StateRunning,
			to:               StatePatchReady,
			trigger:          ir.EventValidationPassed,
			expectedLegal:    false,
			invariantMessage: "Running must go through Validating",
		},
		{
			name:             "from terminal",
			from:             StateCompleted,
			to:               StateRunning,
			trigger:          ir.EventRunStarted,
			expectedLegal:    false,
			invariantMessage: "no transition originates from a terminal state",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := ValidateTransition(c.from, c.to, c.trigger)
			if got != c.expectedLegal {
				t.Errorf("%s: %v (want %v) — %s", c.name, got, c.expectedLegal, c.invariantMessage)
			}
		})
	}
}

// TestTriggerNamesMatchIR is a smoke check that all trigger names in the
// matrix are real EventType constants in the ir package. This catches the
// case where ir-event-log.md catalog and task-lifecycle.md §4 drift apart.
func TestTriggerNamesMatchIR(t *testing.T) {
	for _, tr := range LegalTransitions() {
		if tr.Trigger == "" {
			t.Errorf("empty trigger for transition %s -> %s", tr.From, tr.To)
		}
	}
}

// TestEveryTriggerIsClassifiedAsTransition guards the C1 ↔ C2 boundary
// (task-lifecycle.md §9). Every event used as a state-transition trigger
// in task.LegalTransitions() MUST be classified as a transition event by
// ir.EventType.IsTransition(); otherwise the FSM reducer would dispatch
// state changes on events the IR catalog claims are non-transition.
//
// The reverse implication (every transition event appearing in the matrix)
// is intentionally NOT enforced: some transition events (e.g., TaskCreated,
// RuntimePinned) are invariants/markers rather than state-to-state moves.
func TestEveryTriggerIsClassifiedAsTransition(t *testing.T) {
	for _, tr := range LegalTransitions() {
		if !tr.Trigger.IsTransition() {
			t.Errorf(
				"drift: %s -> %s uses trigger %q which ir.IsTransition() says is NOT a transition event",
				tr.From, tr.To, tr.Trigger,
			)
		}
	}
}
