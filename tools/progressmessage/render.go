package main

import (
	"fmt"
	"strings"

	"github.com/teamswyg/riido-contracts/progressmessage"
)

func renderDoc(
	m docManifest,
	dsl progressmessage.DSLDocument,
	ir progressmessage.IRDocument,
) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# %s\n\n%s\n\n", m.Title, generatedNotice)
	fmt.Fprintf(&b, "> Riido task: %s\n\n", m.RiidoTask)
	fmt.Fprintf(&b, "Executable SSOT: [`%s`](../../%s).\n\n", m.DSL, m.DSL)
	fmt.Fprintf(&b, "%s\n\n", m.Summary)
	renderSummary(&b, m, dsl, ir)
	renderRules(&b, m.Rules)
	renderProjection(&b, m)
	renderMessages(&b, ir.Messages)
	renderLoop(&b, m.Loop)
	return b.String()
}

func renderSummary(
	b *strings.Builder,
	m docManifest,
	dsl progressmessage.DSLDocument,
	ir progressmessage.IRDocument,
) {
	counts := usageCounts(ir)
	fmt.Fprintf(b, "- Status: `%s`\n", status(ir))
	fmt.Fprintf(b, "- Evidence artifact: `%s`\n", m.EvidenceArtifact)
	fmt.Fprintf(b, "- Contract: `%s`\n", dsl.ContractID)
	fmt.Fprintf(b, "- Messages: `%d/%d`\n", len(ir.Messages), ir.MaxMessages)
	fmt.Fprintf(b, "- Required/active/reserved: `%d/%d/%d`\n\n",
		counts[progressmessage.UsageRequired],
		counts[progressmessage.UsageActive],
		counts[progressmessage.UsageReserved],
	)
}
