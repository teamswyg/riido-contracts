package assignment

import "testing"

func TestAssignmentTransitionBDDScenarios(t *testing.T) {
	for _, tc := range transitionBDDCases() {
		t.Run(tc.name, func(t *testing.T) {
			if got := CanTransition(tc.from, tc.to); got != tc.want {
				t.Fatalf("%s: CanTransition(%q,%q) = %v, want %v", tc.acceptance, tc.from, tc.to, got, tc.want)
			}
		})
	}
}

type transitionBDDCase struct {
	name       string
	from, to   AssignmentState
	want       bool
	acceptance string
}

func transitionBDDCases() []transitionBDDCase {
	return []transitionBDDCase{
		{"daemon can lease queued work", AssignmentQueued, AssignmentLeased, true, "poll start can claim queued assignment"},
		{"daemon reports active assignment as running", AssignmentReady, AssignmentRunning, true, "daemon event can move ready assignment to running"},
		{"running completes after agent event", AssignmentRunning, AssignmentCompleted, true, "terminal success is reached only from running"},
		{"terminal state cannot restart", AssignmentCompleted, AssignmentRunning, false, "completed assignment is terminal"},
		{"queued cannot skip to completed", AssignmentQueued, AssignmentCompleted, false, "assignment must pass through daemon-active states before success"},
	}
}
