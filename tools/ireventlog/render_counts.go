package main

import (
	"fmt"
	"strings"
)

func renderNativeConfig(b *strings.Builder, counts nativeConfigCounts) {
	b.WriteString("## Native Config Classification\n\n")
	b.WriteString("| Class | EventTypes |\n| --- | ---: |\n")
	fmt.Fprintf(b, "| forbidden | `%d` |\n", counts.Forbidden)
	fmt.Fprintf(b, "| pre-execute | `%d` |\n", counts.PreExecute)
	fmt.Fprintf(b, "| required | `%d` |\n", counts.Required)
	fmt.Fprintf(b, "| phase-dependent | `%d` |\n\n", counts.PhaseDependent)
}

func renderReducer(b *strings.Builder, model model) {
	b.WriteString("## Reducer Surface\n\n")
	fmt.Fprintf(b, "- CanonicalEvent fields: `%d`\n", model.CanonicalEventFields)
	fmt.Fprintf(b, "- ReduceResult fields: `%s`\n\n",
		strings.Join(model.ReduceResultFieldNames, "`, `"))
}
