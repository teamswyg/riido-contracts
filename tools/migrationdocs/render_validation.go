package main

import (
	"fmt"
	"strings"
)

func renderValidationGates(b *strings.Builder, m manifest) {
	b.WriteString("## Validation Gates\n\n")
	b.WriteString("Required for this repository:\n\n")
	writeCommandBlock(b, m.ValidationGates.RequiredCommands)
	b.WriteString("Architecture-doc migration PRs must also pass:\n\n")
	writeCommandBlock(b, m.ValidationGates.ArchitectureCommands)
	b.WriteString("When contract fixtures are added, public CI must also verify:\n\n")
	writeBullets(b, m.ValidationGates.FixtureChecks)
}

func writeCommandBlock(b *strings.Builder, commands []string) {
	b.WriteString("```bash\n")
	for _, command := range commands {
		fmt.Fprintf(b, "%s\n", command)
	}
	b.WriteString("```\n\n")
}
