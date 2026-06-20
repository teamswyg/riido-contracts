package main

import (
	"fmt"
	"strings"
)

func renderCandidateContracts(b *strings.Builder, m manifest) {
	b.WriteString("## Candidate Contracts\n\n")
	b.WriteString("| Candidate | Source in private `riido_daemon` | Target decision |\n| --- | --- | --- |\n")
	for _, c := range m.CandidateContracts {
		fmt.Fprintf(b, "| %s | %s | %s |\n", cell(c.Candidate), cell(c.Source), cell(c.Decision))
	}
	b.WriteString("\n")
}

func renderVersioning(b *strings.Builder, m manifest) {
	b.WriteString("## Versioning\n\n")
	fmt.Fprintf(b, "%s\n\n", m.Versioning.Intro)
	b.WriteString("| Axis | Owner before split | Contract handling |\n| --- | --- | --- |\n")
	for _, axis := range m.Versioning.Axes {
		fmt.Fprintf(b, "| %s | %s | %s |\n",
			cell(axis.Axis), cell(axis.OwnerBeforeSplit), cell(axis.ContractHandling))
	}
	b.WriteString("\n")
}

func renderWorkMap(b *strings.Builder, m manifest) {
	b.WriteString("## Migration Work Map\n\n")
	b.WriteString("| Area | Riido task | Target repository |\n| --- | --- | --- |\n")
	for _, entry := range m.MigrationWorkMap {
		fmt.Fprintf(b, "| %s | %s | %s |\n",
			cell(entry.Area), cell(entry.RiidoTask), cell(entry.TargetRepository))
	}
}
