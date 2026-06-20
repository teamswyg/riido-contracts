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
	renderOperations(&b, "Onboarding operations", model.OnboardingOperations)
	renderOperations(&b, "Direct configuration creates", model.DirectCreateOperations)
	renderInlineList(&b, "Fixture schema fields", model.FixtureFields)
	renderInlineList(&b, "Create request fields", model.CreateRequestFields)
	renderFixtureRows(&b, m.FixtureRows)
	renderList(&b, "Invariant anchors", m.Invariants)
	renderInlineList(&b, "No-diff route fragments", m.NoDiffPathFragments)
	renderLoop(&b, m.Loop)
	return b.String()
}

func renderSummary(b *strings.Builder, model model) {
	fmt.Fprintf(b, "- Evidence artifact: `%s`\n", model.Manifest.EvidenceArtifact)
	fmt.Fprintf(b, "- Operations: `%d`; onboarding: `%d`; direct create: `%d`\n",
		len(model.Operations), len(model.OnboardingOperations), len(model.DirectCreateOperations))
	fmt.Fprintf(b, "- Fixture fields: `%d`; create request fields: `%d`; scenarios: `%d`\n",
		len(model.FixtureFields), len(model.CreateRequestFields), model.ScenarioCount)
	fmt.Fprintf(b, "- DSL/IR match: `%t`; OpenAPI match: `%t`; no-diff paths clean: `%t`\n\n",
		model.DSLIRMatch, model.OpenAPIMatch, model.NoDiffPathsClean)
}
