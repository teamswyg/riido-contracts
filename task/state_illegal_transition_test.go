package task

import (
	"testing"

	"github.com/teamswyg/riido-contracts/ir"
)

func TestIllegalTransitionsRejected(t *testing.T) {
	for _, c := range illegalTransitionCases() {
		t.Run(c.name, func(t *testing.T) {
			got := ValidateTransition(c.from, c.to, c.trigger)
			if got != c.expectedLegal {
				t.Errorf("%s: %v (want %v) - %s", c.name, got, c.expectedLegal, c.invariantMessage)
			}
		})
	}
}

type illegalTransitionCase struct {
	name             string
	from, to         TaskState
	trigger          ir.EventType
	expectedLegal    bool
	invariantMessage string
}

func illegalTransitionCases() []illegalTransitionCase {
	return []illegalTransitionCase{
		{"skip queued", StateCreated, StateRunning, ir.EventRunStarted, false, "Created cannot jump to Running"},
		{"skip validating", StateRunning, StatePatchReady, ir.EventValidationPassed, false, "Running must go through Validating"},
		{"from terminal", StateCompleted, StateRunning, ir.EventRunStarted, false, "no transition originates from a terminal state"},
	}
}
