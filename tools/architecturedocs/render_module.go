package main

import (
	"fmt"
	"strings"
)

func renderModuleDoc(m manifest) string {
	var b strings.Builder
	b.WriteString("# Contracts Module Decomposition\n\n")
	b.WriteString(generatedMarker + "\n\n")
	fmt.Fprintf(&b, "> Riido task: %s\n\n", cell(m.RiidoTask))
	fmt.Fprintf(&b, "%s\n\n", cell(m.Summary))
	b.WriteString("## Packages\n\n")
	b.WriteString("| Package/path | Role | Must not own |\n| --- | --- | --- |\n")
	for _, p := range m.Packages {
		fmt.Fprintf(&b, "| %s | %s | %s |\n", p.Display, cell(p.Role), cell(p.MustNotOwn))
	}
	b.WriteString("\n## Dependency Rules\n\n")
	fmt.Fprintf(&b, "%s\n\n", cell(m.PackageRules.Dependency))
	fmt.Fprintf(&b, "%s\n\n", cell(m.PackageRules.Forbidden))
	b.WriteString("`go list -m all` must return only this module.\n\n")
	renderContractShape(&b, m)
	b.WriteString("## Tag Boundary\n\n")
	fmt.Fprintf(&b, "%s\n", cell(m.PackageRules.TagBoundary))
	return strings.TrimRight(b.String(), "\n") + "\n"
}

func renderContractShape(b *strings.Builder, m manifest) {
	b.WriteString("## Contract Shape\n\n")
	b.WriteString("A contract package may contain:\n\n")
	for _, item := range m.ContractShape.Allowed {
		fmt.Fprintf(b, "- %s\n", cell(item))
	}
	b.WriteString("\nA contract package must not contain:\n\n")
	for _, item := range m.ContractShape.Forbidden {
		fmt.Fprintf(b, "- %s\n", cell(item))
	}
	b.WriteString("\n")
}
