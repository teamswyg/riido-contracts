package main

import (
	"fmt"
	"strings"
)

func renderManifestLoopInventory(b *strings.Builder, report scanReport) {
	b.WriteString("## Manifest Loop Inventory\n\n")
	fmt.Fprintf(b, "- Complete manifest loops: `%d`\n", report.ManifestLoops.Complete)
	fmt.Fprintf(b, "- Missing manifest loops: `%d`\n\n", report.ManifestLoops.Missing)
	b.WriteString("| Group | Missing loops | Sample paths |\n| --- | ---: | --- |\n")
	if len(report.ManifestLoops.MissingGroups) == 0 {
		b.WriteString("| None | 0 | - |\n\n")
		return
	}
	for _, group := range report.ManifestLoops.MissingGroups {
		fmt.Fprintf(b, "| `%s` | %d | %s |\n", group.Group, group.Count, manifestLoopSampleText(report, group.Group))
	}
	b.WriteString("\n")
}

func manifestLoopSampleText(report scanReport, group string) string {
	for _, sample := range report.ManifestLoops.MissingSamples {
		if sample.Group == group {
			return renderSamplePaths(sample.Paths)
		}
	}
	return "None"
}
