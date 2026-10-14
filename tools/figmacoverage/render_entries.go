package main

import (
	"fmt"
	"strings"
)

func renderEntries(b *strings.Builder, title string, entries []coverageEntry) {
	b.WriteString("## " + title + "\n\n")
	if len(entries) == 0 {
		b.WriteString("_None._\n\n")
		return
	}
	b.WriteString("| Node | Name | Status | Owners | Generated paths |\n")
	b.WriteString("| --- | --- | --- | --- | --- |\n")
	for _, entry := range entries {
		fmt.Fprintf(
			b,
			"| `%s` | %s | `%s` | %s | %s |\n",
			cell(entry.NodeID),
			cell(entry.Name),
			cell(entry.CoverageStatus),
			cell(strings.Join(entry.OwnerRepos, ", ")),
			cell(strings.Join(entry.GeneratedPaths, ", ")),
		)
	}
	b.WriteString("\n")
}

func renderEntryDetails(b *strings.Builder, entries []coverageEntry) {
	b.WriteString("### Coverage Entry Details\n\n")
	for _, entry := range entries {
		fmt.Fprintf(b, "- `%s` %s", cell(entry.NodeID), cell(entry.Name))
		if entry.AbsorbedByTopLevelNodeID != "" {
			fmt.Fprintf(b, " absorbed by `%s`", cell(entry.AbsorbedByTopLevelNodeID))
		}
		b.WriteString("\n")
		for _, fact := range entry.CoveredFacts {
			fmt.Fprintf(b, "  - %s\n", cell(fact))
		}
	}
	b.WriteString("\n")
}
