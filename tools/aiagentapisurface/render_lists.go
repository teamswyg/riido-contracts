package main

import (
	"fmt"
	"strings"
)

func renderList(b *strings.Builder, title string, values []string) {
	fmt.Fprintf(b, "## %s\n\n", title)
	for _, value := range values {
		fmt.Fprintf(b, "- `%s`\n", value)
	}
	fmt.Fprintln(b)
}

func renderTuples(b *strings.Builder, title string, values []operationTuple) {
	fmt.Fprintf(b, "## %s\n\n", title)
	for _, value := range values {
		fmt.Fprintf(b, "- `%s %s` (`%s`)\n", value.Method, value.Path, value.OperationID)
	}
	fmt.Fprintln(b)
}
