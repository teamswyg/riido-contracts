package main

import (
	"fmt"
	"strings"
)

func renderManifestLoopInventory(b *strings.Builder, report scanReport) {
	b.WriteString("## Manifest Loop Inventory\n\n")
	fmt.Fprintf(b, "- Complete manifest loops: `%d`\n", report.ManifestLoops.Complete)
	fmt.Fprintf(b, "- Direct manifest loops: `%d`\n", report.ManifestLoops.Direct)
	fmt.Fprintf(b, "- Delegated manifest loops: `%d`\n", report.ManifestLoops.Delegated)
	fmt.Fprintf(b, "- Missing manifest loops: `%d`\n", report.ManifestLoops.Missing)
	fmt.Fprintf(b, "- Missing loop budget: `%d`\n\n", report.ManifestLoopBudget.MaxMissing)
	b.WriteString("| Group | Missing loops | Budget | Sample paths |\n| --- | ---: | ---: | --- |\n")
	if len(report.ManifestLoops.MissingGroups) == 0 {
		b.WriteString("| None | 0 | 0 | - |\n\n")
		return
	}
	for _, group := range report.ManifestLoops.MissingGroups {
		fmt.Fprintf(b, "| `%s` | %d | %d | %s |\n", group.Group, group.Count, manifestLoopGroupBudget(report, group.Group), manifestLoopSampleText(report, group.Group))
	}
	b.WriteString("\n")
}

func manifestLoopGroupBudget(report scanReport, group string) int {
	return report.ManifestLoopBudget.MaxMissingByGroup[group]
}

func manifestLoopSampleText(report scanReport, group string) string {
	for _, sample := range report.ManifestLoops.MissingSamples {
		if sample.Group == group {
			return renderSamplePaths(sample.Paths)
		}
	}
	return "None"
}
