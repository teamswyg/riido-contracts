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
	renderOperations(&b, model.Operations)
	renderSchemas(&b, m.RequiredSchemaFields)
	renderInlineList(&b, "Policies", m.RequiredPolicies)
	renderEnums(&b, model.Enums)
	renderInlineList(&b, "Invariant anchors", m.Invariants)
	renderLoop(&b, m.Loop)
	return b.String()
}

func renderSummary(b *strings.Builder, model model) {
	fmt.Fprintf(b, "- Evidence artifact: `%s`\n", model.Manifest.EvidenceArtifact)
	fmt.Fprintf(b, "- Operations: `%d`; schemas: `%d`; policies: `%d`; enums: `%d`\n",
		len(model.Operations), len(model.Schemas), len(model.Policies), len(model.Enums))
	fmt.Fprintf(b, "- Scenarios: `%d`; editability stream variant: `%s`\n",
		model.ScenarioCount, model.Manifest.RequiredStreamVariant)
	fmt.Fprintf(b, "- DSL/IR match: `%t`; OpenAPI match: `%t`; stream pass: `%t`\n\n",
		model.DSLIRMatch, model.OpenAPIMatch, model.StreamPass)
}
