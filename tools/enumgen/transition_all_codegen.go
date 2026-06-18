package main

import (
	"bytes"
	"fmt"
)

func writeTransitionAll(b *bytes.Buffer, model transitionModel) {
	fmt.Fprintf(b, "var %s = []%s{\n", transitionPrivateName(model.Spec, "All"), model.Spec.Name)
	for _, entry := range model.Spec.Entries {
		writeTransitionEntry(b, model, entry)
	}
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintf(b, "func %s() []%s {\n", model.Spec.AllFunc, model.Spec.Name)
	fmt.Fprintf(
		b,
		"\treturn append([]%s(nil), %s...)\n",
		model.Spec.Name,
		transitionPrivateName(model.Spec, "All"),
	)
	fmt.Fprintln(b, "}")
}

func writeTransitionEntry(b *bytes.Buffer, model transitionModel, entry transitionEntry) {
	fmt.Fprintf(
		b,
		"\t{From: %s, To: %s",
		codeRef(model.From, model.Spec.Package, entry.From),
		codeRef(model.To, model.Spec.Package, entry.To),
	)
	if model.Spec.EventEnum != "" {
		fmt.Fprintf(b, ", Trigger: %s", codeRef(model.Event, model.Spec.Package, entry.Event))
	}
	fmt.Fprintln(b, "},")
}
