package main

import (
	"bytes"
	"fmt"
)

func writeAssignmentFSMProvider(b *bytes.Buffer, model assignmentFSMModel) {
	fmt.Fprintln(b, "func GeneratedAssignmentFSM() AssignmentFSM {")
	fmt.Fprintln(b, "\treturn generatedAssignmentFSM{}")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func GeneratedAssignmentFSMServiceProvider() AssignmentFSMServiceProvider {")
	fmt.Fprintln(b, "\treturn generatedAssignmentFSMServiceProvider{}")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func (generatedAssignmentFSMServiceProvider) AssignmentFSM() AssignmentFSM {")
	fmt.Fprintln(b, "\treturn generatedAssignmentFSM{}")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func (generatedAssignmentFSM) Name() string {")
	fmt.Fprintf(b, "\treturn %q\n", model.Meta.ReadmeSection)
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func (generatedAssignmentFSM) TypeUnion() AssignmentFSMTypeUnion {")
	fmt.Fprintf(b, "\treturn AssignmentFSMTypeUnion%s\n", model.Meta.TypeUnion)
	fmt.Fprintln(b, "}")
}
