package main

import (
	"fmt"
	"strings"
)

func renderManifest(m manifest) string {
	var b strings.Builder
	b.WriteString("# Domain Contract Docs\n\n")
	b.WriteString(generatedMarker + "\n\n")
	fmt.Fprintf(&b, "Executable SSOT: [`%s`](README.riido.json)\n\n", defaultManifest)
	fmt.Fprintf(&b, "%s\n\n", cell(m.Summary))
	renderArchitectureLinks(&b, m)
	renderChanges(&b, m)
	renderExternalBoundaries(&b, m)
	fmt.Fprintf(&b, "%s are tracked in [`%s`](%s).\n",
		cell(m.OpenQuestions.Label), cell(m.OpenQuestions.Path), cell(m.OpenQuestions.Path))
	return strings.TrimRight(b.String(), "\n") + "\n"
}

func renderArchitectureLinks(b *strings.Builder, m manifest) {
	b.WriteString("## Architecture Links\n\n")
	for _, link := range m.ArchitectureLinks {
		fmt.Fprintf(b, "- %s: [`%s`](%s)\n", cell(link.Label), cell(link.Path), cell(link.Path))
	}
	b.WriteString("\n")
}
