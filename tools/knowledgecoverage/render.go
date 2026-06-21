package main

import (
	"fmt"
	"strings"
)

func renderDoc(m manifest, report scanReport) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# %s\n\n%s\n\n", m.Title, generatedNotice)
	fmt.Fprintf(&b, "Executable SSOT: [`%s`](%s).\n\n", defaultManifest, defaultManifest)
	fmt.Fprintf(&b, "- Status: `%s`\n", status(report))
	fmt.Fprintf(&b, "- Evidence artifact: `%s`\n", m.EvidenceArtifact)
	fmt.Fprintf(&b, "- Workflow trigger paths: `%d`\n", len(m.WorkflowTriggerPaths))
	fmt.Fprintf(&b, "- Markdown readers scanned: `%d`\n", report.ScannedCount)
	fmt.Fprintf(&b, "- Generated readers: `%d`\n", report.GeneratedCount)
	fmt.Fprintf(&b, "- Executable readers: `%d`\n", report.ExecutableCount)
	fmt.Fprintf(&b, "- Adjacent manifests: `%d`\n", report.AdjacentCount)
	fmt.Fprintf(&b, "- Generated readers with adjacent manifests: `%d`\n", report.GeneratedAdjacentCount)
	fmt.Fprintf(&b, "- Executable readers with adjacent manifests: `%d`\n", report.ExecutableAdjacentCount)
	fmt.Fprintf(&b, "- Manual reader candidates: `%d`\n", report.ManualCount)
	fmt.Fprintf(&b, "- Manifest inventory: `%d`\n\n", report.ManifestInventory)
	renderManifestInventory(&b, report.ManifestInventoryByGroup)
	renderManifestInventorySamples(&b, report.ManifestInventorySamples)
	renderManifestLoopInventory(&b, report)
	renderManualSamples(&b, report.ManualSamples)
	renderLoop(&b, m.Loop)
	return b.String()
}

func renderManifestInventory(b *strings.Builder, groups []manifestGroupCount) {
	b.WriteString("## Manifest Inventory\n\n")
	b.WriteString("| Group | Count |\n| --- | ---: |\n")
	for _, item := range groups {
		fmt.Fprintf(b, "| `%s` | %d |\n", item.Group, item.Count)
	}
	b.WriteString("\n")
}

func renderManualSamples(b *strings.Builder, samples []docRecord) {
	b.WriteString("## Manual Reader Candidates\n\n")
	if len(samples) == 0 {
		b.WriteString("None.\n\n")
		return
	}
	for _, sample := range samples {
		fmt.Fprintf(b, "- `%s` (%d lines)\n", sample.Path, sample.Lines)
	}
	b.WriteString("\n")
}
