package main

import (
	"bytes"
	"fmt"
)

func writeTransitionAllowed(b *bytes.Buffer, model transitionModel) {
	fmt.Fprintf(
		b,
		"var %s = map[%s]bool{\n",
		transitionPrivateName(model.Spec, "Allowed"),
		model.Spec.Name,
	)
	for _, entry := range model.Spec.Entries {
		writeTransitionAllowedEntry(b, model, entry)
	}
	fmt.Fprintln(b, "}")
}

func writeTransitionAllowedEntry(
	b *bytes.Buffer,
	model transitionModel,
	entry transitionEntry,
) {
	fmt.Fprintf(
		b,
		"\t{From: %s, To: %s",
		codeRef(model.From, model.Spec.Package, entry.From),
		codeRef(model.To, model.Spec.Package, entry.To),
	)
	if model.Spec.EventEnum != "" {
		fmt.Fprintf(b, ", Trigger: %s", codeRef(model.Event, model.Spec.Package, entry.Event))
	}
	fmt.Fprintln(b, "}: true,")
}
