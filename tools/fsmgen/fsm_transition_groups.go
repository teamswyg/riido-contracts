package main

import (
	"sort"

	"github.com/teamswyg/riido-contracts/assignment"
	"github.com/teamswyg/riido-contracts/ir"
	"github.com/teamswyg/riido-contracts/task"
)

func taskTransitionsByFromAndTrigger() map[task.TaskStateCode]map[ir.EventTypeCode][]task.TaskStateCode {
	groups := map[task.TaskStateCode]map[ir.EventTypeCode][]task.TaskStateCode{}
	for _, transition := range task.LegalTransitionCodes() {
		if _, ok := groups[transition.From]; !ok {
			groups[transition.From] = map[ir.EventTypeCode][]task.TaskStateCode{}
		}
		groups[transition.From][transition.Trigger] = append(groups[transition.From][transition.Trigger], transition.To)
	}
	for _, byTrigger := range groups {
		for trigger := range byTrigger {
			sort.SliceStable(byTrigger[trigger], func(i, j int) bool {
				return byTrigger[trigger][i] < byTrigger[trigger][j]
			})
		}
	}
	return groups
}

func assignmentTransitionsByFrom() map[assignment.AssignmentStateCode][]assignment.AssignmentStateCode {
	groups := map[assignment.AssignmentStateCode][]assignment.AssignmentStateCode{}
	for _, transition := range assignment.AssignmentTransitionCodes() {
		groups[transition.From] = append(groups[transition.From], transition.To)
	}
	for from := range groups {
		sort.SliceStable(groups[from], func(i, j int) bool {
			return groups[from][i] < groups[from][j]
		})
	}
	return groups
}

func sortedTaskTriggers(groups map[ir.EventTypeCode][]task.TaskStateCode) []ir.EventTypeCode {
	triggers := make([]ir.EventTypeCode, 0, len(groups))
	for trigger := range groups {
		triggers = append(triggers, trigger)
	}
	sort.SliceStable(triggers, func(i, j int) bool {
		return triggers[i] < triggers[j]
	})
	return triggers
}
