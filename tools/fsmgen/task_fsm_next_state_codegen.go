package main

import (
	"bytes"
	"fmt"

	"github.com/teamswyg/riido-contracts/ir"
	"github.com/teamswyg/riido-contracts/task"
)

func writeTaskFSMNextStateTable(b *bytes.Buffer, _ taskFSMModel) {
	fmt.Fprintln(b, "type taskFSMNextStateKey struct {")
	fmt.Fprintln(b, "\tFrom    TaskStateCode")
	fmt.Fprintln(b, "\tTrigger ir.EventTypeCode")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintln(b, "var taskFSMNextStates = map[taskFSMNextStateKey][]TaskStateCode{")
	for _, row := range taskFSMNextStateRows() {
		fmt.Fprintf(b, "\t{From: %s, Trigger: %s}: ", taskStateCodeRef(row.From), eventTypeCodeRef(row.Trigger))
		writeInlineTaskStateSlice(b, row.Next)
	}
	fmt.Fprintln(b, "}")
}

func writeTaskFSMNextStates(b *bytes.Buffer, _ taskFSMModel) {
	fmt.Fprintln(b, "func (generatedTaskFSM) CanTransition(from TaskStateCode, to TaskStateCode, trigger ir.EventTypeCode) bool {")
	fmt.Fprintln(b, "\treturn ValidateTransitionCode(from, to, trigger)")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func (generatedTaskFSM) NextStates(from TaskStateCode, trigger ir.EventTypeCode) []TaskStateCode {")
	fmt.Fprintln(b, "\tstates := taskFSMNextStates[taskFSMNextStateKey{From: from, Trigger: trigger}]")
	fmt.Fprintln(b, "\treturn append([]TaskStateCode(nil), states...)")
	fmt.Fprintln(b, "}")
}

type taskFSMNextStateRow struct {
	From    task.TaskStateCode
	Trigger ir.EventTypeCode
	Next    []task.TaskStateCode
}

func taskFSMNextStateRows() []taskFSMNextStateRow {
	groups := taskTransitionsByFromAndTrigger()
	rows := []taskFSMNextStateRow{}
	for _, from := range task.AllTaskStateCodes() {
		triggers := sortedTaskTriggers(groups[from])
		for _, trigger := range triggers {
			rows = append(rows, taskFSMNextStateRow{
				From:    from,
				Trigger: trigger,
				Next:    groups[from][trigger],
			})
		}
	}
	return rows
}

func writeInlineTaskStateSlice(b *bytes.Buffer, states []task.TaskStateCode) {
	fmt.Fprint(b, "[]TaskStateCode{")
	for index, state := range states {
		if index > 0 {
			fmt.Fprint(b, ", ")
		}
		fmt.Fprint(b, taskStateCodeRef(state))
	}
	fmt.Fprintln(b, "},")
}
