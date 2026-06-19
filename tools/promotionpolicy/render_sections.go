package main

import (
	"fmt"
	"strings"
)

func renderPromotionRule(b *strings.Builder, m manifest) {
	b.WriteString("## Promotion Rule\n\n")
	b.WriteString("Promote a fact only when all conditions are true:\n\n")
	for i, condition := range m.PromotionConditions {
		fmt.Fprintf(b, "%d. %s\n", i+1, cell(condition))
	}
	fmt.Fprintf(b, "\n%s\n\n", cell(m.SingleRuntimeRule))
}

func renderVersioning(b *strings.Builder, m manifest) {
	b.WriteString("## Versioning\n\n")
	b.WriteString("The module uses Git tags. Package schema constants are independent from the Go module tag:\n\n")
	for _, axis := range m.SchemaVersionAxes {
		fmt.Fprintf(b, "- `%s` %s.\n", cell(axis.Axis), cell(axis.Rule))
	}
	fmt.Fprintf(b, "\n%s\n\n", cell(m.ModuleTagRule))
}
