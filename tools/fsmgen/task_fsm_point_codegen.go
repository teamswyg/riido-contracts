package main

import (
	"bytes"
	"fmt"

	"github.com/teamswyg/riido-contracts/task"
)

func writeTaskFSMPoints(b *bytes.Buffer, model taskFSMModel) {
	writeTaskStateSetVar(b, "taskFSMStartStateSet", model.StartStates)
	fmt.Fprintln(b)
	writeTaskStateSetVar(b, "taskFSMEndStateSet", model.EndStates)
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func (fsm generatedTaskFSM) PointKind(state TaskStateCode) TaskFSMPointKind {")
	fmt.Fprintln(b, "\tswitch {")
	fmt.Fprintln(b, "\tcase fsm.IsStartState(state):")
	fmt.Fprintln(b, "\t\treturn TaskFSMPointStart")
	fmt.Fprintln(b, "\tcase fsm.IsEndState(state):")
	fmt.Fprintln(b, "\t\treturn TaskFSMPointEnd")
	fmt.Fprintln(b, "\tcase state.IsKnown():")
	fmt.Fprintln(b, "\t\treturn TaskFSMPointIntermediate")
	fmt.Fprintln(b, "\tdefault:")
	fmt.Fprintln(b, "\t\treturn TaskFSMPointUnknown")
	fmt.Fprintln(b, "\t}")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func (generatedTaskFSM) IsStartState(state TaskStateCode) bool {")
	fmt.Fprintln(b, "\treturn taskFSMStartStateSet[state]")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func (generatedTaskFSM) IsEndState(state TaskStateCode) bool {")
	fmt.Fprintln(b, "\treturn taskFSMEndStateSet[state]")
	fmt.Fprintln(b, "}")
}

func writeTaskStateSetVar(b *bytes.Buffer, name string, states []task.TaskStateCode) {
	fmt.Fprintf(b, "var %s = map[TaskStateCode]bool{\n", name)
	for _, state := range states {
		fmt.Fprintf(b, "\t%s: true,\n", taskStateCodeRef(state))
	}
	fmt.Fprintln(b, "}")
}
