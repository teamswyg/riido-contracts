package main

import (
	"fmt"

	"github.com/teamswyg/riido-contracts/assignment"
	"github.com/teamswyg/riido-contracts/task"
)

func taskStateCodesFromConsts(names []string) ([]task.TaskStateCode, error) {
	out := make([]task.TaskStateCode, 0, len(names))
	for _, name := range names {
		code, ok := taskStateCodeFromConst(name)
		if !ok {
			return nil, fmt.Errorf("unknown task state const %s", name)
		}
		out = append(out, code)
	}
	return out, nil
}

func taskStateCodeFromConst(name string) (task.TaskStateCode, bool) {
	for _, code := range task.AllTaskStateCodes() {
		if taskStateConstName(code) == name {
			return code, true
		}
	}
	return task.TaskStateCodeUnknown, false
}

func taskStateConstName(code task.TaskStateCode) string {
	return "State" + pascalCodeSuffix(code.String())
}

func assignmentStateCodesFromConsts(names []string) ([]assignment.AssignmentStateCode, error) {
	out := make([]assignment.AssignmentStateCode, 0, len(names))
	for _, name := range names {
		code, ok := assignmentStateCodeFromConst(name)
		if !ok {
			return nil, fmt.Errorf("unknown assignment state const %s", name)
		}
		out = append(out, code)
	}
	return out, nil
}

func assignmentStateCodeFromConst(name string) (assignment.AssignmentStateCode, bool) {
	for _, code := range assignment.AllAssignmentStateCodes() {
		if assignmentStateConstName(code) == name {
			return code, true
		}
	}
	return assignment.AssignmentStateCodeUnknown, false
}

func assignmentStateConstName(code assignment.AssignmentStateCode) string {
	return "Assignment" + pascalCodeSuffix(code.String())
}
