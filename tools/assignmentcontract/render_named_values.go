package main

import (
	"fmt"
	"strings"
)

func renderNamedValues(b *strings.Builder, title string, values []namedValue) {
	b.WriteString("## " + title + "\n\n")
	b.WriteString("| Name | Value |\n")
	b.WriteString("| --- | --- |\n")
	for _, value := range values {
		fmt.Fprintf(b, "| `%s` | `%s` |\n", value.Name, value.Value)
	}
	b.WriteString("\n")
}
