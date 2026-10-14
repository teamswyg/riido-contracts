package main

import (
	"fmt"
	"strings"
)

func renderRepositoryBoundaries(b *strings.Builder, m manifest) {
	b.WriteString("## Repository Boundaries\n\n")
	b.WriteString("`riido-contracts` may contain:\n\n")
	writeBullets(b, m.RepositoryBoundaries.MayContain)
	b.WriteString("`riido-contracts` must not contain:\n\n")
	writeBullets(b, m.RepositoryBoundaries.MustNotContain)
}

func renderMigrationOrder(b *strings.Builder, m manifest) {
	b.WriteString("## Migration Order\n\n")
	writeNumbered(b, m.MigrationOrder)
}

func renderMigrationSlices(b *strings.Builder, m manifest) {
	b.WriteString("## Current Migration Slices\n\n")
	for _, slice := range m.MigrationSlices {
		fmt.Fprintf(b, "### %s\n\n", slice.Title)
		for _, line := range slice.Intro {
			fmt.Fprintf(b, "%s\n", line)
		}
		if len(slice.Intro) > 0 {
			b.WriteString("\n")
		}
		if len(slice.Does) > 0 {
			b.WriteString("This slice does:\n\n")
			writeBullets(b, slice.Does)
		}
		if slice.DoesNot != "" {
			fmt.Fprintf(b, "%s\n\n", slice.DoesNot)
		}
	}
}
