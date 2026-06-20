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
	fmt.Fprintf(&b, "Executable SSOT: `%s` plus generated `ir` and `task` Go contracts.\n\n", m.SourceDSL)
	fmt.Fprintf(&b, "%s\n\n", m.Summary)
	renderSummary(&b, model)
	renderTransitions(&b, model.TransitionEvents)
	renderNativeConfig(&b, model.NativeConfigCounts)
	renderReducer(&b, model)
	renderBullets(&b, "Invariant Anchors", m.Invariants)
	renderBullets(&b, "Adjacent Contracts", m.AdjacentContracts)
	renderBullets(&b, "Responsibility Boundary", m.Responsibilities)
	renderLoop(&b, m.Loop)
	return b.String()
}

func renderSummary(b *strings.Builder, model model) {
	fmt.Fprintf(b, "- Evidence artifact: `%s`\n", model.Manifest.EvidenceArtifact)
	fmt.Fprintf(b, "- EventTypes: `%d`; transition: `%d`; non-transition: `%d`\n",
		model.EventCount, model.TransitionCount, model.NonTransitionCount)
	fmt.Fprintf(b, "- Task FSM transitions: `%d`; unique IR triggers: `%d`\n\n",
		model.TaskFSMTransitionCount, model.TaskFSMTriggerCount)
}
