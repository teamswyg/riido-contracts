package main

import (
	"bytes"
	"fmt"
)

func writeTransitionValidate(b *bytes.Buffer, model transitionModel) {
	fmt.Fprintf(
		b,
		"func %s(from %s, to %s",
		model.Spec.Validate,
		typeRef(model.From, model.Spec.Package),
		typeRef(model.To, model.Spec.Package),
	)
	if model.Spec.EventEnum != "" {
		fmt.Fprintf(b, ", trigger %s", typeRef(model.Event, model.Spec.Package))
	}
	fmt.Fprintln(b, ") bool {")
	if model.Spec.AllowSame {
		fmt.Fprintln(b, "\tif from == to && from.IsKnown() {")
		fmt.Fprintln(b, "\t\treturn true")
		fmt.Fprintln(b, "\t}")
	}
	fmt.Fprintf(b, "\treturn %s[%s", transitionPrivateName(model.Spec, "Allowed"), model.Spec.Name)
	fmt.Fprint(b, "{From: from, To: to")
	if model.Spec.EventEnum != "" {
		fmt.Fprint(b, ", Trigger: trigger")
	}
	fmt.Fprintln(b, "}]")
	fmt.Fprintln(b, "}")
}
