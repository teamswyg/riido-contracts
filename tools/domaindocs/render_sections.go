package main

import (
	"fmt"
	"strings"
)

func renderChanges(b *strings.Builder, m manifest) {
	for _, change := range m.Changes {
		fmt.Fprintf(b, "## %s %s\n\n", cell(change.Task), cell(change.Verb))
		for _, item := range change.Items {
			fmt.Fprintf(b, "- %s\n", cell(item))
		}
		b.WriteString("\n")
	}
}

func renderExternalBoundaries(b *strings.Builder, m manifest) {
	b.WriteString("## External Boundaries\n\n")
	for _, boundary := range m.ExternalBoundaries {
		fmt.Fprintf(b, "- %s\n", cell(boundary))
	}
	b.WriteString("\n")
}
