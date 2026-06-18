package main

import (
	"bytes"
	"fmt"
)

func writeTaskFSMStates(b *bytes.Buffer, model taskFSMModel) {
	writeTaskStateSliceVar(b, "taskFSMStartStates", model.StartStates)
	fmt.Fprintln(b)
	writeTaskStateSliceVar(b, "taskFSMEndStates", model.EndStates)
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func (generatedTaskFSM) States() []TaskStateCode {")
	fmt.Fprintln(b, "\treturn AllTaskStateCodes()")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func (generatedTaskFSM) StartStates() []TaskStateCode {")
	fmt.Fprintln(b, "\treturn append([]TaskStateCode(nil), taskFSMStartStates...)")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func (generatedTaskFSM) EndStates() []TaskStateCode {")
	fmt.Fprintln(b, "\treturn append([]TaskStateCode(nil), taskFSMEndStates...)")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func (generatedTaskFSM) TerminalStates() []TaskStateCode {")
	fmt.Fprintln(b, "\treturn append([]TaskStateCode(nil), taskFSMEndStates...)")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func (generatedTaskFSM) Transitions() []TaskTransitionCode {")
	fmt.Fprintln(b, "\treturn LegalTransitionCodes()")
	fmt.Fprintln(b, "}")
}
