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
	fmt.Fprintf(&b, "Executable SSOT: `%s` package plus this manifest.\n\n", m.Package)
	fmt.Fprintf(&b, "%s\n\n", m.Summary)
	renderSummary(&b, model)
	renderVocabulary(&b, model)
	renderProtocols(&b, model.Protocols)
	renderBullets(&b, "Invariant Anchors", m.Invariants)
	renderLoop(&b, m.Loop)
	return b.String()
}

func renderSummary(b *strings.Builder, model model) {
	fmt.Fprintf(b, "- Evidence artifact: `%s`\n", model.Manifest.EvidenceArtifact)
	fmt.Fprintf(b, "- ProviderCapability fields: `%d`; fingerprint input fields: `%d`\n",
		model.ProviderCapabilityFields, model.FingerprintInputFields)
	fmt.Fprintf(b, "- Protocol kinds: `%d`; critical arg sets: `%d`\n\n",
		len(model.Protocols), model.CriticalArgSetCount)
}
