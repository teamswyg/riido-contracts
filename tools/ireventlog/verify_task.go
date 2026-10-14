package main

import (
	"fmt"

	"github.com/teamswyg/riido-contracts/task"
)

func verifyTaskTriggers() error {
	for _, transition := range task.LegalTransitionCodes() {
		if !transition.Trigger.IsTransition() {
			return fmt.Errorf("task FSM trigger %s is not an IR transition", transition.Trigger)
		}
	}
	return nil
}
