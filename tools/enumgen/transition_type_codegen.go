package main

import (
	"bytes"
	"fmt"
)

func writeTransitionTypes(b *bytes.Buffer, model transitionModel) {
	fmt.Fprintf(b, "type %s struct {\n", model.Spec.Name)
	fmt.Fprintf(b, "\tFrom %s\n", typeRef(model.From, model.Spec.Package))
	fmt.Fprintf(b, "\tTo %s\n", typeRef(model.To, model.Spec.Package))
	if model.Spec.EventEnum != "" {
		fmt.Fprintf(b, "\tTrigger %s\n", typeRef(model.Event, model.Spec.Package))
	}
	fmt.Fprintln(b, "}")
}
