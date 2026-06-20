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
	fmt.Fprintf(&b, "Executable SSOT: `%s` plus generated `task` FSM SPI.\n\n", m.SourceDSL)
	fmt.Fprintf(&b, "%s\n\n", m.Summary)
	renderSummary(&b, model)
	renderStates(&b, model.States)
	renderTransitions(&b, model.Transitions)
	renderBullets(&b, "Invariant Anchors", m.Invariants)
	renderBullets(&b, "Responsibility Boundary", m.Responsibilities)
	renderLoop(&b, m.Loop)
	return b.String()
}

func renderSummary(b *strings.Builder, model model) {
	fmt.Fprintf(b, "- Evidence artifact: `%s`; FSM schema version: `%d`\n", model.Manifest.EvidenceArtifact, model.FSMSchema)
	fmt.Fprintf(b, "- States: `%d`; transitions: `%d`; start: `%s`; terminal: `%s`\n\n",
		len(model.States), model.TransitionCount,
		strings.Join(model.StartStates, "`, `"), strings.Join(model.TerminalStates, "`, `"))
}
