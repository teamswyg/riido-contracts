package main

import (
	"github.com/teamswyg/riido-contracts/ir"
	"github.com/teamswyg/riido-contracts/task"
)

func taskFSMTriggerCount() int {
	seen := map[ir.EventTypeCode]bool{}
	for _, transition := range task.LegalTransitionCodes() {
		seen[transition.Trigger] = true
	}
	return len(seen)
}
