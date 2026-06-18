package main

import (
	"bytes"
	"fmt"

	"github.com/teamswyg/riido-contracts/assignment"
)

func writeAssignmentFSMPoints(b *bytes.Buffer, model assignmentFSMModel) {
	writeAssignmentStateSetVar(b, "assignmentFSMStartStateSet", model.StartStates)
	fmt.Fprintln(b)
	writeAssignmentStateSetVar(b, "assignmentFSMEndStateSet", model.EndStates)
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func (fsm generatedAssignmentFSM) PointKind(state AssignmentStateCode) AssignmentFSMPointKind {")
	fmt.Fprintln(b, "\tswitch {")
	fmt.Fprintln(b, "\tcase fsm.IsStartState(state):")
	fmt.Fprintln(b, "\t\treturn AssignmentFSMPointStart")
	fmt.Fprintln(b, "\tcase fsm.IsEndState(state):")
	fmt.Fprintln(b, "\t\treturn AssignmentFSMPointEnd")
	fmt.Fprintln(b, "\tcase state.IsKnown():")
	fmt.Fprintln(b, "\t\treturn AssignmentFSMPointIntermediate")
	fmt.Fprintln(b, "\tdefault:")
	fmt.Fprintln(b, "\t\treturn AssignmentFSMPointUnknown")
	fmt.Fprintln(b, "\t}")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func (generatedAssignmentFSM) IsStartState(state AssignmentStateCode) bool {")
	fmt.Fprintln(b, "\treturn assignmentFSMStartStateSet[state]")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func (generatedAssignmentFSM) IsEndState(state AssignmentStateCode) bool {")
	fmt.Fprintln(b, "\treturn assignmentFSMEndStateSet[state]")
	fmt.Fprintln(b, "}")
}

func writeAssignmentStateSetVar(b *bytes.Buffer, name string, states []assignment.AssignmentStateCode) {
	fmt.Fprintf(b, "var %s = map[AssignmentStateCode]bool{\n", name)
	for _, state := range states {
		fmt.Fprintf(b, "\t%s: true,\n", assignmentStateCodeRef(state))
	}
	fmt.Fprintln(b, "}")
}
