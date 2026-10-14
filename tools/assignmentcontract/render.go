package main

import (
	"fmt"
	"strings"
)

func renderDoc(m manifest, c contract) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# %s\n\n%s\n\n", m.Title, generatedNotice)
	fmt.Fprintf(&b, "> Riido task: %s\n\n", m.RiidoTask)
	fmt.Fprintf(&b, "Executable SSOT: [`%s`](../../%s).\n\n", m.Contract, m.Contract)
	fmt.Fprintf(&b, "%s\n\n", m.Summary)
	renderSummary(&b, m, c)
	renderBullets(&b, "Responsibility", m.Responsibilities)
	renderBullets(&b, "Boundary", m.Boundaries)
	renderContractTables(&b, c)
	renderBullets(&b, "Invariants", m.Invariants)
	renderLoop(&b, m.Loop)
	return b.String()
}

func renderSummary(b *strings.Builder, m manifest, c contract) {
	fmt.Fprintf(b, "- Evidence artifact: `%s`\n", m.EvidenceArtifact)
	fmt.Fprintf(b, "- Contract schema: `%s`\n", c.SchemaVersion)
	fmt.Fprintf(b, "- Service schema: `%s`\n", c.ServiceSchemaVersion)
	fmt.Fprintf(b, "- States / poll actions / task events: `%d / %d / %d`\n",
		len(c.AssignmentStates), len(c.PollActions), len(c.TaskEvents))
	b.WriteString("\n")
}
