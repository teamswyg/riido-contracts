package task

import (
	"testing"

	"github.com/teamswyg/riido-contracts/ir"
)

func TestHappyPathReachable(t *testing.T) {
	for _, step := range happyPathSteps() {
		if !ValidateTransition(step.from, step.to, step.trigger) {
			t.Errorf("happy-path step %s -(%s)-> %s rejected by matrix", step.from, step.trigger, step.to)
		}
	}
}

type happyPathStep struct {
	from, to TaskState
	trigger  ir.EventType
}

func happyPathSteps() []happyPathStep {
	return []happyPathStep{
		{StateCreated, StateQueued, ir.EventTaskQueued},
		{StateQueued, StateClaimed, ir.EventTaskClaimed},
		{StateClaimed, StatePreparing, ir.EventWorkdirPreparing},
		{StatePreparing, StateRunning, ir.EventRunStarted},
		{StateRunning, StateValidating, ir.EventRunReportedDone},
		{StateValidating, StatePatchReady, ir.EventValidationPassed},
		{StatePatchReady, StateHumanReview, ir.EventReviewRequested},
		{StateHumanReview, StateCompleted, ir.EventHumanApproved},
	}
}
