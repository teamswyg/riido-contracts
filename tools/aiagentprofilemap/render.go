package main

import (
	"fmt"
	"strings"
)

func renderDoc(model model) string {
	var b strings.Builder
	m := model.Manifest
	fmt.Fprintf(&b, "# %s\n\n%s\n\n", m.Title, generatedNotice)
	fmt.Fprintf(&b, "> Riido task: %s\n\n", m.RiidoTask)
	fmt.Fprintf(&b, "Executable SSOT: `%s`, `%s`, and `%s`.\n\n",
		m.DSLFixture, m.IRFixture, m.OpenAPIFixture)
	fmt.Fprintf(&b, "%s\n\n", m.Summary)
	renderSummary(&b, model)
	renderOperation(&b, model.Operation)
	renderInlineList(&b, "Read-model rules", model.Policy.Rules)
	renderSchemas(&b, m.RequiredSchemaFields)
	renderInlineList(&b, "Invariant anchors", m.Invariants)
	renderLoop(&b, m.Loop)
	return b.String()
}

func renderSummary(b *strings.Builder, model model) {
	fmt.Fprintf(b, "- Evidence artifact: `%s`\n", model.Manifest.EvidenceArtifact)
	fmt.Fprintf(b, "- Operations: `1`; schemas: `%d`; policy rules: `%d`\n",
		len(model.Schemas), len(model.Policy.Rules))
	fmt.Fprintf(b, "- Scenarios: `%d`\n", model.ScenarioCount)
	fmt.Fprintf(b, "- DSL/IR match: `%t`; OpenAPI match: `%t`; map shape pass: `%t`\n\n",
		model.DSLIRMatch, model.OpenAPIMatch, model.MapShapePass)
}
