package main

import (
	"fmt"
	"strings"
)

func renderLoop(b *strings.Builder, loop evidenceLoop) {
	b.WriteString("## Evidence Loop\n\n")
	rows := []struct{ name, value string }{
		{"Observe", loop.Observation},
		{"Hypothesis", loop.Hypothesis},
		{"Execute", loop.Execute},
		{"Evaluate", loop.Evaluate},
		{"Retrospective", loop.Retrospective},
	}
	for _, row := range rows {
		fmt.Fprintf(b, "- %s: %s\n", row.name, row.value)
	}
}
