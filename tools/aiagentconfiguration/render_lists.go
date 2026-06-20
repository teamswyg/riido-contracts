package main

import (
	"fmt"
	"strings"
)

func renderInlineList(b *strings.Builder, title string, values []string) {
	fmt.Fprintf(b, "## %s\n\n", title)
	for i, value := range values {
		if i > 0 {
			fmt.Fprint(b, ", ")
		}
		fmt.Fprintf(b, "`%s`", value)
	}
	fmt.Fprintln(b)
	fmt.Fprintln(b)
}

func renderSchemas(b *strings.Builder, schemas []schemaExpectation) {
	fmt.Fprintln(b, "## Schema Field Sets")
	fmt.Fprintln(b)
	for _, schema := range schemas {
		fmt.Fprintf(b, "- `%s`: `%d` fields\n", schema.Schema, len(schema.Fields))
	}
	fmt.Fprintln(b)
}
