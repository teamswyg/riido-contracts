package main

import (
	"fmt"
	"strings"
)

func renderSource(b *strings.Builder, m manifest) {
	b.WriteString("## Source\n\n")
	b.WriteString("| Field | Value |\n")
	b.WriteString("| --- | --- |\n")
	fmt.Fprintf(b, "| Figma file | `%s` |\n", cell(m.Figma.FileName))
	fmt.Fprintf(b, "| File key | `%s` |\n", cell(m.Figma.FileKey))
	fmt.Fprintf(b, "| Primary page | `%s` %s |\n", cell(m.Figma.PageID), cell(m.Figma.PageName))
	fmt.Fprintf(b, "| Inspected at | `%s` |\n", cell(m.Figma.InspectedAt))
	fmt.Fprintf(b, "| Inspection authority | %s |\n", cell(m.InspectionMethod.Authority))
	fmt.Fprintf(b, "| Stabilized by | %s |\n", cell(strings.Join(m.StabilizedBy, ", ")))
	b.WriteString("\n")
	b.WriteString("The executable manifest's `stabilized_by` list is the downstream projection mirror source.\n\n")
	b.WriteString("Inspection uses `figma.root.children`, `await figma.setCurrentPageAsync(page)`, ")
	b.WriteString("and `page.children.length`. Metadata XML/read tools are supporting evidence only ")
	b.WriteString("because lazy/unloaded page reads must not redefine page-level child counts.\n\n")
}

func renderPolicy(b *strings.Builder, policy coveragePolicy) {
	b.WriteString("## Coverage Rule\n\n")
	fmt.Fprintf(b, "- Summary: %s\n", cell(policy.Summary))
	fmt.Fprintf(b, "- Top-down: %s\n", cell(policy.TopDown))
	fmt.Fprintf(b, "- Bottom-up: %s\n\n", cell(policy.BottomUp))
}
