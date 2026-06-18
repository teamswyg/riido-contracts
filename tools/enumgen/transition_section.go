package main

import (
	"bytes"
	"fmt"
)

func formatTransitionSection(
	name string,
	model transitionModel,
	write func(*bytes.Buffer, transitionModel),
) ([]byte, error) {
	var b bytes.Buffer
	writeHeader(&b, model.Spec.Package)
	writeTransitionImports(&b, model)
	write(&b, model)
	return formatSource(name, b.Bytes())
}

func writeTransitionImports(b *bytes.Buffer, model transitionModel) {
	imports := transitionImports(model.Spec, model.From, model.To, model.Event)
	if len(imports) == 0 {
		return
	}
	fmt.Fprintln(b, "import (")
	for _, imp := range imports {
		fmt.Fprintf(b, "\t%q\n", imp)
	}
	fmt.Fprintln(b, ")")
	fmt.Fprintln(b)
}
