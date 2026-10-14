package main

import (
	"fmt"
	"strings"
)

func renderFixtures(b *strings.Builder, m manifest, summaries []fixtureSummary) {
	b.WriteString("## Fixtures\n\n")
	b.WriteString("| Contract | DSL | IR | OpenAPI | Ops | v2 ops | Generated paths |\n")
	b.WriteString("| --- | --- | --- | --- | --- | --- | --- |\n")
	for index, summary := range summaries {
		ref := m.Fixtures[index]
		fmt.Fprintf(b, "| `%s` | `%s` | `%s` | `%s` | `%d` | `%d` | `%d` |\n",
			summary.ContractID,
			ref.DSL,
			ref.IR,
			ref.OpenAPI,
			summary.OperationCount,
			summary.V2OperationCount,
			summary.GeneratedPathCount,
		)
	}
	b.WriteString("\n")
}

func totalOperations(summaries []fixtureSummary) int {
	total := 0
	for _, summary := range summaries {
		total += summary.OperationCount
	}
	return total
}
