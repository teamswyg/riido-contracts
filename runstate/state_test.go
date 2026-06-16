package runstate

import "testing"

func TestRunStateRoundTrip(t *testing.T) {
	for _, state := range AllStates() {
		code := state.Code()
		if !code.IsKnown() {
			t.Fatalf("%s code is unknown", state)
		}
		if got := code.RunState(); got != state {
			t.Fatalf("%s round trip = %s", state, got)
		}
		if got := ParseRunStateCode(string(state)); got != code {
			t.Fatalf("ParseRunStateCode(%q) = %s, want %s", state, got, code)
		}
	}
}

func TestRunStateTerminalPredicate(t *testing.T) {
	for _, state := range []RunState{StateCompleted, StateFailed, StateCancelled, StateTimedOut, StateIdleStopped} {
		if !state.IsTerminal() {
			t.Fatalf("%s must be terminal", state)
		}
	}
	if StateRunning.IsTerminal() {
		t.Fatal("running must not be terminal")
	}
}
