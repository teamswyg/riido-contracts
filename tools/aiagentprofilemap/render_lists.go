package main

import (
	"fmt"
	"strings"
)

func renderInlineList(b *strings.Builder, title string, values []string) {
	fmt.Fprintf(b, "## %s\n\n", title)
	for _, value := range values {
		fmt.Fprintf(b, "- %s\n", value)
	}
	b.WriteString("\n")
}

func renderSchemas(b *strings.Builder, schemas []schemaExpectation) {
	b.WriteString("## Schemas\n\n")
	for _, schema := range schemas {
		fmt.Fprintf(b, "- `%s`: `%s`\n", schema.Schema, strings.Join(schema.Fields, "`, `"))
	}
	b.WriteString("\n")
}

func renderLoop(b *strings.Builder, loop evidenceLoop) {
	b.WriteString("## Evidence Loop\n\n")
	fmt.Fprintf(b, "- Observe: %s\n", loop.Observation)
	fmt.Fprintf(b, "- Hypothesis: %s\n", loop.Hypothesis)
	fmt.Fprintf(b, "- Execute: %s\n", loop.Execute)
	fmt.Fprintf(b, "- Evaluate: %s\n", loop.Evaluate)
	fmt.Fprintf(b, "- Retrospective: %s\n", loop.Retrospective)
}
