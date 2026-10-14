package main

import (
	"fmt"
	"strings"
)

func renderRepoDependencies(b *strings.Builder, deps []repoDependency) {
	b.WriteString("## Repository Dependencies\n\n")
	b.WriteString("| ID | From | To | Facts | Local scope |\n")
	b.WriteString("| --- | --- | --- | --- | --- |\n")
	for _, dep := range deps {
		fmt.Fprintf(
			b,
			"| `%s` | `%s` | `%s` | %s | %s |\n",
			markdownCell(dep.ID),
			markdownCell(dep.FromRepo),
			markdownCell(dep.ToRepo),
			markdownCell(strings.Join(dep.FactIDs, ", ")),
			markdownCell(dep.LocalScope),
		)
	}
	b.WriteString("\n")
}

func renderLoop(b *strings.Builder) {
	b.WriteString("## Loop Gates\n\n")
	b.WriteString("- Top-down: product/design evidence -> owning contracts SSOT -> ")
	b.WriteString("generated API/IR projection -> downstream harnesses.\n")
	b.WriteString("- Bottom-up: implementation or operations evidence -> local SSOT ")
	b.WriteString("finding -> owning contracts SSOT when meaning changes -> regeneration.\n")
}
