package main

import (
	"fmt"
	"strings"
)

func renderOperations(b *strings.Builder, operations []operation) {
	b.WriteString("## Operations\n\n")
	for _, op := range operations {
		fmt.Fprintf(b, "- `%s` `%s %s` -> `%d %s`",
			op.OperationID, op.Method, op.Path, responseStatus(op), responseName(op))
		if requestName(op) != "" {
			fmt.Fprintf(b, "; request `%s`", requestName(op))
		}
		fmt.Fprintf(b, "; scenarios `%d`\n", len(op.Scenarios))
	}
	b.WriteString("\n")
}
