package main

import (
	"fmt"
	"strings"
)

func renderManifest(m manifest) string {
	var b strings.Builder
	b.WriteString("# Contracts Context Map\n\n")
	b.WriteString(generatedMarker + "\n\n")
	fmt.Fprintf(&b, "> Riido task: %s\n\n", cell(m.RiidoTask))
	fmt.Fprintf(&b, "This file is the public context map for `%s`.\n\n", cell(m.Module))
	b.WriteString("## Role\n\n")
	fmt.Fprintf(&b, "%s\n\n", cell(m.Role))
	renderOwnedContexts(&b, m)
	renderNonOwnedContexts(&b, m)
	fmt.Fprintf(&b, "%s\n\n", cell(m.BoundaryRule))
	renderDirectionRules(&b, m)
	renderSSOTLinks(&b, m)
	return strings.TrimRight(b.String(), "\n") + "\n"
}

func renderOwnedContexts(b *strings.Builder, m manifest) {
	b.WriteString("## Owned Contexts\n\n")
	b.WriteString("| Context | Package | Responsibility |\n| --- | --- | --- |\n")
	for _, ctx := range m.OwnedContexts {
		fmt.Fprintf(b, "| %s | `%s` | %s |\n",
			cell(ctx.Context), cell(ctx.Package), cell(ctx.Responsibility))
	}
	b.WriteString("\n")
}

func renderNonOwnedContexts(b *strings.Builder, m manifest) {
	b.WriteString("## Non-Owned Contexts\n\n")
	b.WriteString("| Context | Owner | Boundary |\n| --- | --- | --- |\n")
	for _, ctx := range m.NonOwnedContexts {
		fmt.Fprintf(b, "| %s | `%s` | %s |\n",
			cell(ctx.Context), cell(ctx.Owner), cell(ctx.Boundary))
	}
	b.WriteString("\n")
}
