package main

import (
	"bytes"
	"fmt"
)

func writeAssignmentFSMMermaid(b *bytes.Buffer, model assignmentFSMModel) {
	fmt.Fprintf(b, "const AssignmentFSMMermaid = `%s`\n\n", assignmentMermaid(model.StartStates, model.EndStates))
	fmt.Fprintln(b, "func (generatedAssignmentFSM) Mermaid() string {")
	fmt.Fprintln(b, "\treturn AssignmentFSMMermaid")
	fmt.Fprintln(b, "}")
}
