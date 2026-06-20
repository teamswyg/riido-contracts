package main

import (
	"fmt"
	"strings"
)

func renderStates(b *strings.Builder, states []stateRow) {
	b.WriteString("## States\n\n")
	b.WriteString("| State | Point |\n| --- | --- |\n")
	for _, state := range states {
		fmt.Fprintf(b, "| `%s` | %s |\n", state.Name, state.Kind)
	}
	b.WriteString("\n")
}
