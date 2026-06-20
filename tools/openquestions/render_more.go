package main

import (
	"fmt"
	"strings"
)

func renderLoop(b *strings.Builder, loop evidenceLoop) {
	b.WriteString("## Evidence Loop\n\n")
	b.WriteString("| Step | Evidence |\n| --- | --- |\n")
	fmt.Fprintf(b, "| Observe | %s |\n", cell(loop.Observation))
	fmt.Fprintf(b, "| Hypothesis | %s |\n", cell(loop.Hypothesis))
	fmt.Fprintf(b, "| Execute | %s |\n", cell(loop.Execute))
	fmt.Fprintf(b, "| Evaluate | %s |\n", cell(loop.Evaluate))
	fmt.Fprintf(b, "| Retrospective | %s |\n", cell(loop.Retrospective))
}

func cell(value string) string {
	value = strings.ReplaceAll(value, "|", "\\|")
	value = strings.ReplaceAll(value, "\n", " ")
	if strings.TrimSpace(value) == "" {
		return "-"
	}
	return value
}
