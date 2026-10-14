package main

import (
	"fmt"
	"strings"
)

func renderStates(b *strings.Builder, states []state) {
	b.WriteString("## Assignment States\n\n")
	b.WriteString("| State | Agent active | Terminal | Transitions |\n")
	b.WriteString("| --- | --- | --- | --- |\n")
	for _, state := range states {
		fmt.Fprintf(b, "| `%s` | `%t` | `%t` | %s |\n",
			state.Value,
			state.AgentActive,
			state.Terminal,
			cell(strings.Join(state.Transitions, ", ")),
		)
	}
	b.WriteString("\n")
}
