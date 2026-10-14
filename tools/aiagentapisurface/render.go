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
	renderTuples(&b, "V2-only operations", model.V2Only)
	renderList(&b, "ClientStreamEvent variants", model.StreamVariants)
	renderList(&b, "Codegen rules", m.CodegenRules)
	renderList(&b, "Invariant anchors", m.Invariants)
	renderLoop(&b, m.Loop)
	return b.String()
}

func renderSummary(b *strings.Builder, model model) {
	fmt.Fprintf(b, "- Evidence artifact: `%s`\n", model.Manifest.EvidenceArtifact)
	fmt.Fprintf(b, "- Operations: `%d`; v1: `%d`; v2: `%d`; v2-only: `%d`\n",
		len(model.Operations), model.V1Count, model.V2Count, len(model.V2Only))
	fmt.Fprintf(b, "- OpenAPI paths: `%d`; OpenAPI operations: `%d`\n",
		model.OpenAPIPathCount, model.OpenAPIOpCount)
	fmt.Fprintf(b, "- DSL/IR match: `%t`; IR/OpenAPI match: `%t`; v2 covers v1: `%t`\n\n",
		model.DSLIRMatch, model.IROpenAPIMatch, model.V2CoversV1)
	fmt.Fprintln(b, "- V2 workspace prefix: `/v2/client/workspaces/{workspace_id}/ai-agent/...`")
	fmt.Fprintln(b)
}
