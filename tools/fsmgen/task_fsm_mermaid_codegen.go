package main

import (
	"bytes"
	"fmt"
)

func writeTaskFSMMermaid(b *bytes.Buffer, model taskFSMModel) {
	fmt.Fprintf(b, "const TaskFSMMermaid = `%s`\n\n", taskMermaid(model.StartStates, model.EndStates))
	fmt.Fprintln(b, "func (generatedTaskFSM) Mermaid() string {")
	fmt.Fprintln(b, "\treturn TaskFSMMermaid")
	fmt.Fprintln(b, "}")
}
