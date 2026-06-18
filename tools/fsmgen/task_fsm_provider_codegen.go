package main

import (
	"bytes"
	"fmt"
)

func writeTaskFSMProvider(b *bytes.Buffer, model taskFSMModel) {
	fmt.Fprintln(b, "func GeneratedTaskFSM() TaskFSM {")
	fmt.Fprintln(b, "\treturn generatedTaskFSM{}")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func GeneratedTaskFSMServiceProvider() TaskFSMServiceProvider {")
	fmt.Fprintln(b, "\treturn generatedTaskFSMServiceProvider{}")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func (generatedTaskFSMServiceProvider) TaskFSM() TaskFSM {")
	fmt.Fprintln(b, "\treturn generatedTaskFSM{}")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func (generatedTaskFSM) Name() string {")
	fmt.Fprintf(b, "\treturn %q\n", model.Meta.ReadmeSection)
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func (generatedTaskFSM) TypeUnion() TaskFSMTypeUnion {")
	fmt.Fprintf(b, "\treturn TaskFSMTypeUnion%s\n", model.Meta.TypeUnion)
	fmt.Fprintln(b, "}")
}
