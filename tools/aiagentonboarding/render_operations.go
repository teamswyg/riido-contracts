package main

import (
	"fmt"
	"strings"
)

func renderOperations(b *strings.Builder, title string, ops []operation) {
	fmt.Fprintf(b, "## %s\n\n", title)
	fmt.Fprintln(b, "| Operation | Method | Path | Response |")
	fmt.Fprintln(b, "| --- | --- | --- | --- |")
	for _, op := range ops {
		response := ""
		if op.Response != nil {
			response = fmt.Sprintf("%d %s", op.Response.Status, op.Response.Ref)
		}
		fmt.Fprintf(b, "| `%s` | `%s` | `%s` | `%s` |\n",
			op.OperationID, op.Method, op.Path, response)
	}
	fmt.Fprintln(b)
}
