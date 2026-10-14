package main

import (
	"fmt"
	"strings"
)

func renderEnums(b *strings.Builder, enums []enumSpec) {
	fmt.Fprintln(b, "## Enums")
	fmt.Fprintln(b)
	for _, enum := range enums {
		fmt.Fprintf(b, "- `%s`: ", enum.Name)
		for i, value := range enumValues(enum) {
			if i > 0 {
				fmt.Fprint(b, ", ")
			}
			fmt.Fprintf(b, "`%s`", value)
		}
		fmt.Fprintln(b)
	}
	fmt.Fprintln(b)
}

func renderLoop(b *strings.Builder, loop evidenceLoop) {
	fmt.Fprintln(b, "## Evidence Loop")
	fmt.Fprintln(b)
	fmt.Fprintf(b, "- Observe: %s\n", loop.Observation)
	fmt.Fprintf(b, "- Hypothesis: %s\n", loop.Hypothesis)
	fmt.Fprintf(b, "- Execute: %s\n", loop.Execute)
	fmt.Fprintf(b, "- Evaluate: %s\n", loop.Evaluate)
	fmt.Fprintf(b, "- Retrospective: %s\n", loop.Retrospective)
}
