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
	fmt.Fprintf(&b, "Executable SSOT: `%s` plus `%s`.\n\n", m.Package, defaultManifestPath)
	fmt.Fprintf(&b, "%s\n\n", m.Summary)
	renderSummary(&b, model)
	renderModel(&b, model)
	renderDependencyPhrases(&b, model)
	renderInlineList(&b, "AgentRuntimeBinding fields", model.BindingFields)
	renderInlineList(&b, "Invariant anchors", m.Invariants)
	fmt.Fprintln(&b)
	renderLoop(&b, m.Loop)
	return b.String()
}

func renderSummary(b *strings.Builder, model model) {
	fmt.Fprintf(b, "- Evidence artifact: `%s`\n", model.Manifest.EvidenceArtifact)
	fmt.Fprintf(b, "- API policy: `%s`; dependency fact: `%s`\n",
		model.Manifest.APIPolicyID, model.Manifest.DependencyFactID)
	fmt.Fprintf(b, "- Daemon refresh: runtime snapshot every `%d` seconds; stale after `%d` seconds\n\n",
		model.SnapshotInterval, model.RuntimeStaleAfter)
}

func renderModel(b *strings.Builder, model model) {
	fmt.Fprintln(b, "## Model")
	fmt.Fprintln(b)
	renderInlineList(b, "Principals", model.PrincipalKinds)
	renderInlineList(b, "Daemon headers", model.DaemonHeaders)
	renderInlineList(b, "Client headers", model.ClientHeaders)
	renderInlineList(b, "Ownership", model.OwnershipEdges)
	renderInlineList(b, "Binding sources", model.BindingSources)
	renderInlineList(b, "Hard exclusions", model.ExcludedFallbacks)
	renderInlineList(b, "Secret non-exposure sinks", model.SecretSinks)
	fmt.Fprintln(b)
}
