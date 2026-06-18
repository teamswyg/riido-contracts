package task

import "testing"

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
