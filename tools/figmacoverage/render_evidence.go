package main

import (
	"fmt"
	"strings"
)

func renderEvidence(b *strings.Builder, m manifest) {
	b.WriteString("## Evidence Controls\n\n")
	fmt.Fprintf(b, "- Verified evidence nodes: %d\n", len(m.VerifiedEvidenceNodes))
	fmt.Fprintf(b, "- Supporting tool limitations: %d\n", len(m.SupportingToolLimitations))
	fmt.Fprintf(b, "- API generated annotation rows: %d\n", len(m.APIGeneratedAnnotationInventory))
	fmt.Fprintf(b, "- Live API generated annotations: %d\n\n", m.APIAnnotationContentPolicy.LiveInspection.TotalAPIGeneratedAnnotations)
	b.WriteString("### API Generated Live Page Counts\n\n")
	for _, page := range m.APIAnnotationContentPolicy.LiveInspection.PageCounts {
		fmt.Fprintf(b, "- `%s` %s: riido=%d api_generated=%d missing_kind=%d missing_background=%d\n",
			cell(page.PageID), cell(page.PageName), page.RiidoAnnotationCount,
			page.APIGeneratedCount, page.MissingOperationKind, page.MissingBackground)
	}
	b.WriteString("\n")
	b.WriteString("### Supporting Tool Limitations\n\n")
	for _, limitation := range m.SupportingToolLimitations {
		fmt.Fprintf(b, "- `%s`\n", cell(limitation.ID))
		fmt.Fprintf(b, "  - Tool: %s\n", cell(limitation.Tool))
		fmt.Fprintf(b, "  - Observed: %s\n", cell(limitation.ObservedResult))
		fmt.Fprintf(b, "  - Authoritative: %s\n", cell(strings.Join(limitation.AuthoritativeResult, ", ")))
		fmt.Fprintf(b, "  - Rule: %s\n", cell(limitation.Rule))
		renderLimitationNote(b, limitation.ID)
	}
	b.WriteString("\n")
}

func renderLimitationNote(b *strings.Builder, id string) {
	switch id {
	case "figma-metadata-page-list-underreports-pages.v1":
		b.WriteString("  - Normalized: get_metadata without `nodeId`; pages `129:5215`, `42:3014`, `0:1`; must not remove `expected_pages`.\n")
	case "figma-headless-file-key-placeholder.v1":
		b.WriteString("  - Normalized: `figma.fileKey=headless`; `MUOd9lctoEHASUStN3vUuK`; authoritative file identity; must not overwrite `figma.file_key`.\n")
	case "figma-onboarding-page-load-timeout.v1":
		b.WriteString("  - Normalized: get_metadata(nodeId=42:3014) after 120s; `Wireframe - 온보딩`; `236:33845`; `236:33847`; six onboarding `riido.*` `API Generated`; must not rewrite `expected_pages`; onboarding generated paths unresolved; `page.children.length=84`; known captured inventory remains 83.\n")
	}
}
