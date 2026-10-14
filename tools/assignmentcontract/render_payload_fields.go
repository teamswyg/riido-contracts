package main

import (
	"fmt"
	"strings"
)

func renderPayloadFields(b *strings.Builder, fields []payloadField) {
	b.WriteString("## Assignment Payload Fields\n\n")
	b.WriteString("| Field | Source | Snapshot | Required | Consumer |\n")
	b.WriteString("| --- | --- | --- | --- | --- |\n")
	for _, field := range fields {
		fmt.Fprintf(b, "| `%s` | `%s` | `%s` | `%t` | %s |\n",
			field.Name,
			field.Source,
			field.Snapshot,
			field.Required,
			cell(field.Consumer),
		)
	}
	b.WriteString("\n")
}
