package main

import (
	"fmt"
	"strings"
)

func renderDoc(m manifest, summaries []fixtureSummary) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# %s\n\n%s\n\n", m.Title, generatedNotice)
	fmt.Fprintf(&b, "> Riido task: %s\n\n", m.RiidoTask)
	b.WriteString("Executable SSOT: API fixture DSL/IR/OpenAPI files listed below.\n\n")
	fmt.Fprintf(&b, "%s\n\n", m.Summary)
	renderSummary(&b, m, summaries)
	renderFixtures(&b, m, summaries)
	renderBullets(&b, "Projection Invariants", m.Invariants)
	renderBullets(&b, "Generated Client Delivery", m.DeliveryRules)
	renderBullets(&b, "Boundary", m.Boundaries)
	renderLoop(&b, m.Loop)
	return b.String()
}

func renderSummary(b *strings.Builder, m manifest, summaries []fixtureSummary) {
	fmt.Fprintf(b, "- Evidence artifact: `%s`\n", m.EvidenceArtifact)
	fmt.Fprintf(b, "- Fixtures: `%d`\n", len(summaries))
	fmt.Fprintf(b, "- Operations: `%d`\n", totalOperations(summaries))
	fmt.Fprintf(b, "- Required generated paths: `%d`\n\n", len(m.RequiredGeneratedPaths))
}
