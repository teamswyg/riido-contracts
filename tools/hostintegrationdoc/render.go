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
	renderInlineList(&b, "Distribution channels", model.DistributionChannels)
	renderInlineList(&b, "Store-managed channels", model.StoreManagedChannels)
	renderInlineList(&b, "Provider routing statuses", model.ProviderStatuses)
	renderInlineList(&b, "Non-owned surfaces", model.NonOwnedSurfaces)
	renderInlineList(&b, "Invariant anchors", m.Invariants)
	renderLoop(&b, m.Loop)
	return b.String()
}

func renderSummary(b *strings.Builder, model model) {
	fmt.Fprintf(b, "- Evidence artifact: `%s`\n", model.Manifest.EvidenceArtifact)
	fmt.Fprintf(b, "- Distribution valid: `%t`; provider routing valid: `%t`\n",
		model.DistributionValid, model.ProviderRoutingValid)
	fmt.Fprintf(b, "- Store-managed classification exclusive: `%t`\n\n",
		model.StoreManagedExclusive)
}
