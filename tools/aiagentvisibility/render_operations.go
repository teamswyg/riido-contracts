package main

import (
	"fmt"
	"strings"
)

func renderOperations(b *strings.Builder, operations []operation) {
	b.WriteString("## Operations\n\n")
	for _, op := range operations {
		fmt.Fprintf(b, "- `%s` `%s %s` -> `%d %s`; policy `%s`; scenarios `%d`\n",
			op.OperationID, op.Method, op.Path, responseStatus(op), responseName(op),
			op.RBACPolicy, len(op.Scenarios))
	}
	b.WriteString("\n")
}

func renderEnum(b *strings.Builder, enum enumSpec) {
	b.WriteString("## Visibility Enum\n\n")
	values := make([]string, 0, len(enum.Values))
	for _, value := range enum.Values {
		values = append(values, value.Value)
	}
	fmt.Fprintf(b, "- `%s`: `%s`\n\n", enum.Name, strings.Join(values, "`, `"))
}
