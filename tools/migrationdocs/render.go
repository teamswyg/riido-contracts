package main

import (
	"fmt"
	"strings"
)

func renderManifest(m manifest) string {
	var b strings.Builder
	b.WriteString("# Riido Contracts Migration Plan\n\n")
	b.WriteString(generatedMarker + "\n\n")
	fmt.Fprintf(&b, "> Riido task: %s\n\n", cell(m.RiidoTask))
	renderGoal(&b, m)
	renderPromotionRule(&b, m)
	renderCandidateContracts(&b, m)
	renderRepositoryBoundaries(&b, m)
	renderVersioning(&b, m)
	renderMigrationOrder(&b, m)
	renderMigrationSlices(&b, m)
	renderValidationGates(&b, m)
	renderWorkMap(&b, m)
	return strings.TrimRight(b.String(), "\n") + "\n"
}

func renderGoal(b *strings.Builder, m manifest) {
	b.WriteString("## Goal\n\n")
	fmt.Fprintf(b, "%s\n\n", m.Goal)
}

func renderPromotionRule(b *strings.Builder, m manifest) {
	b.WriteString("## Promotion Rule\n\n")
	b.WriteString("Move a fact into `riido-contracts` only when all conditions are true:\n\n")
	writeNumbered(b, m.PromotionRule.Conditions)
	fmt.Fprintf(b, "%s\n\n", m.PromotionRule.Fallback)
}
