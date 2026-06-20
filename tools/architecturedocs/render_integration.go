package main

import (
	"fmt"
	"strings"
)

func renderIntegrationDoc(m manifest) string {
	var b strings.Builder
	b.WriteString("# Contracts Integration Matrix\n\n")
	b.WriteString(generatedMarker + "\n\n")
	fmt.Fprintf(&b, "> Riido task: %s\n\n", cell(m.RiidoTask))
	b.WriteString("This repo has no external runtime dependencies. Integration means downstream compatibility, not live infrastructure.\n\n")
	renderPublicGates(&b, m)
	renderDownstreamGates(&b, m)
	renderLocalCommands(&b, m)
	return strings.TrimRight(b.String(), "\n") + "\n"
}

func renderPublicGates(b *strings.Builder, m manifest) {
	b.WriteString("## Public Gates\n\n")
	b.WriteString("| Surface | Verification | External dependency |\n| --- | --- | --- |\n")
	for _, gate := range m.PublicGates {
		fmt.Fprintf(b, "| %s | %s | %s |\n",
			cell(gate.Surface), cell(gate.Verification), cell(gate.ExternalDependency))
	}
	b.WriteString("\n")
}

func renderDownstreamGates(b *strings.Builder, m manifest) {
	b.WriteString("## Downstream Gates\n\n")
	b.WriteString("| Consumer | Expected gate |\n| --- | --- |\n")
	for _, gate := range m.DownstreamGates {
		fmt.Fprintf(b, "| `%s` | %s |\n", cell(gate.Consumer), cell(gate.ExpectedGate))
	}
	b.WriteString("\n")
	b.WriteString("Downstream gates must run in their owning repositories. This repo only proves the contract package is internally consistent and tag-ready.\n\n")
}

func renderLocalCommands(b *strings.Builder, m manifest) {
	b.WriteString("## Local Commands\n\n```bash\n")
	for _, command := range m.LocalCommands {
		fmt.Fprintf(b, "%s\n", command)
	}
	b.WriteString("```\n")
}
