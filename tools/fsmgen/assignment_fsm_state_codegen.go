package main

import (
	"bytes"
	"fmt"
)

func writeAssignmentFSMStates(b *bytes.Buffer, model assignmentFSMModel) {
	writeAssignmentStateSliceVar(b, "assignmentFSMStartStates", model.StartStates)
	fmt.Fprintln(b)
	writeAssignmentStateSliceVar(b, "assignmentFSMEndStates", model.EndStates)
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func (generatedAssignmentFSM) States() []AssignmentStateCode {")
	fmt.Fprintln(b, "\treturn AllAssignmentStateCodes()")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func (generatedAssignmentFSM) StartStates() []AssignmentStateCode {")
	fmt.Fprintln(b, "\treturn append([]AssignmentStateCode(nil), assignmentFSMStartStates...)")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func (generatedAssignmentFSM) EndStates() []AssignmentStateCode {")
	fmt.Fprintln(b, "\treturn append([]AssignmentStateCode(nil), assignmentFSMEndStates...)")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func (generatedAssignmentFSM) TerminalStates() []AssignmentStateCode {")
	fmt.Fprintln(b, "\treturn append([]AssignmentStateCode(nil), assignmentFSMEndStates...)")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func (generatedAssignmentFSM) Transitions() []AssignmentTransitionCode {")
	fmt.Fprintln(b, "\treturn AssignmentTransitionCodes()")
	fmt.Fprintln(b, "}")
}
