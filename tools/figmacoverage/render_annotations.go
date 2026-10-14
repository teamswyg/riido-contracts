package main

import (
	"fmt"
	"strings"
)

func renderAnnotationPolicy(b *strings.Builder, policy annotationContentPolicy) {
	b.WriteString("## API Generated Annotation Content Policy\n\n")
	for _, line := range policy.LabelFormat {
		fmt.Fprintf(b, "- %s\n", cell(line))
	}
	fmt.Fprintf(b, "- Rule: %s\n", cell(policy.Rule))
	b.WriteString("- Transport terms: `text/event-stream`, non-stream `GET`, non-`GET`.\n")
	for _, retired := range policy.RetiredCategories {
		fmt.Fprintf(b, "- Retired `%s` `%s`: retired with zero annotations; %s\n",
			cell(retired.CategoryID), cell(retired.CategoryLabel), cell(retired.ToolLimitation))
	}
	b.WriteString("\n")
}

func renderAnnotations(b *strings.Builder, annotations []annotationInventory) {
	b.WriteString("## API Generated Annotation Inventory\n\n")
	b.WriteString("| Figma path | Canonical path | v2 counterpart | Kind | Count | UI area | Background |\n")
	b.WriteString("| --- | --- | --- | --- | ---: | --- | --- |\n")
	for _, item := range annotations {
		fmt.Fprintf(
			b,
			"| `%s` | `%s` | `%s` | `%s` | %d | %s | %s |\n",
			cell(item.FigmaGeneratedPath),
			cell(item.CanonicalGeneratedPath),
			cell("v2."+item.CanonicalGeneratedPath),
			cell(item.OperationKind),
			item.AnnotationCount,
			cell(item.UIArea),
			cell(item.Background),
		)
	}
	b.WriteString("\n")
}

func renderResolvedAnnotations(b *strings.Builder, annotations []annotation) {
	b.WriteString("## Resolved API Generated Annotation Nodes\n\n")
	for _, annotation := range annotations {
		fmt.Fprintf(b, "- `%s`: `%s` -> `%s` (`%s`) %s\n",
			cell(annotation.NodeID), cell(annotation.FigmaGeneratedPath),
			cell(annotation.CanonicalGeneratedPath), cell(annotation.CategoryLabel), cell(annotation.Resolution))
		if strings.Contains(annotation.FigmaLabel, "작업중") {
			b.WriteString("  - 상세내용은 작업중입니다\n")
		}
	}
	b.WriteString("\n")
}
