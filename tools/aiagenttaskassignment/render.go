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
	renderInlineList(&b, "Forbidden request fields", m.ForbiddenRequestFields)
	renderInlineList(&b, "Non-target path fragments", m.NoDiffPathFragments)
	renderInlineList(&b, "Invariant anchors", m.Invariants)
	renderLoop(&b, m.Loop)
	return b.String()
}

func renderSummary(b *strings.Builder, model model) {
	fmt.Fprintf(b, "- Evidence artifact: `%s`\n", model.Manifest.EvidenceArtifact)
	fmt.Fprintf(b, "- Operations: `%d`; schemas: `%d`; policies: `%d`\n",
		len(model.Operations), len(model.Schemas), len(model.Policies))
	fmt.Fprintf(b, "- Scenarios: `%d`\n", model.ScenarioCount)
	fmt.Fprintf(b, "- DSL/IR match: `%t`; OpenAPI match: `%t`\n",
		model.DSLIRMatch, model.OpenAPIMatch)
	fmt.Fprintf(b, "- Forbidden request fields absent: `%t`; non-target paths absent: `%t`\n\n",
		model.ForbiddenFieldsAbsent, model.NoDiffPathsAbsent)
}
