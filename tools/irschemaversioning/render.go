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
	fmt.Fprintf(&b, "Executable SSOT: `%s` package validators plus this manifest.\n\n", m.Package)
	fmt.Fprintf(&b, "%s\n\n", m.Summary)
	renderSummary(&b, model)
	renderScopes(&b, model.ScopeRules)
	renderFields(&b, model)
	renderBullets(&b, "Invariant Anchors", m.Invariants)
	renderLoop(&b, m.Loop)
	return b.String()
}

func renderSummary(b *strings.Builder, model model) {
	fmt.Fprintf(b, "- Evidence artifact: `%s`\n", model.Manifest.EvidenceArtifact)
	fmt.Fprintf(b, "- CanonicalEvent fields: `%d`; scopes: `%d`; validation entrypoints: `%d`\n",
		model.CanonicalEventFields, model.EventScopeCount, model.ValidateEntrypoints)
	fmt.Fprintf(b, "- Common static required fields: `%d`; ActorID conditional checks: `%d`\n",
		model.CommonRequiredCount, model.ActorIDConditional)
	fmt.Fprintf(b, "- Fake placeholder fields: `%d`; values: `%d`; violation codes: `%d`\n\n",
		model.FakePlaceholderFields, model.FakePlaceholderValues, model.ViolationCodeCount)
}
