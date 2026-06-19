package main

import (
	"fmt"
	"strings"
)

func renderPages(b *strings.Builder, pages []page) {
	b.WriteString("## Page Registry\n\n")
	b.WriteString("| Figma page | Name | Top-level children |\n")
	b.WriteString("| --- | --- | ---: |\n")
	for _, page := range pages {
		fmt.Fprintf(b, "| `%s` | %s | %d |\n", cell(page.NodeID), cell(page.Name), page.ChildCount)
	}
	b.WriteString("\n")
}

func renderCoverageSummary(b *strings.Builder, entries []coverageEntry) {
	counts := countStatuses(entries)
	b.WriteString("## Coverage Summary\n\n")
	b.WriteString("| Status | Count |\n")
	b.WriteString("| --- | ---: |\n")
	for _, status := range sortedKeys(counts) {
		fmt.Fprintf(b, "| `%s` | %d |\n", cell(status), counts[status])
	}
	b.WriteString("\n")
}
