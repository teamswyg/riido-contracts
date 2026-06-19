package main

import (
	"fmt"
	"strings"
)

func renderFacts(b *strings.Builder, facts []fact) {
	b.WriteString("## Facts\n\n")
	b.WriteString("| ID | Fact | Owner | Downstreams |\n")
	b.WriteString("| --- | --- | --- | --- |\n")
	for _, f := range facts {
		fmt.Fprintf(
			b,
			"| `%s` | %s | %s | %s |\n",
			markdownCell(f.ID),
			markdownCell(f.Fact),
			renderOwner(f.Owner),
			markdownCell(joinDownstreamRepos(f.Downstreams)),
		)
	}
	b.WriteString("\n")
	for _, f := range facts {
		renderFactDetail(b, f)
	}
}

func renderFactDetail(b *strings.Builder, f fact) {
	fmt.Fprintf(b, "### `%s`\n\n", markdownCell(f.ID))
	fmt.Fprintf(b, "- Human phrase: %s\n", markdownCell(f.HumanDocPhrase))
	fmt.Fprintf(b, "- Owner: %s\n", renderOwner(f.Owner))
	b.WriteString("- Source refs:\n")
	for _, ref := range f.SourceRefs {
		fmt.Fprintf(b, "  - %s\n", renderSourceRef(ref))
	}
	b.WriteString("- Downstreams:\n")
	for _, downstream := range f.Downstreams {
		fmt.Fprintf(b, "  - `%s`: %s\n", markdownCell(downstream.Repo), markdownCell(downstream.LocalScope))
	}
	b.WriteString("\n")
}
