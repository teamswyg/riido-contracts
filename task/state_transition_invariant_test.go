package task

import "testing"

func TestInvariant5_NoTransitionFromTerminal(t *testing.T) {
	for _, tr := range LegalTransitions() {
		if tr.From.IsTerminal() {
			t.Errorf("invariant 5 violated: terminal %s -> %s via %s", tr.From, tr.To, tr.Trigger)
		}
	}
}
