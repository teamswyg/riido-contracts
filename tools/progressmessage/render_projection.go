package main

import (
	"fmt"
	"strings"
)

func renderProjection(b *strings.Builder, m docManifest) {
	b.WriteString("## Projection\n\n```text\n")
	for index, step := range m.Projection {
		if index == 0 {
			fmt.Fprintf(b, "%s\n", step)
			continue
		}
		fmt.Fprintf(b, "  -> %s\n", step)
	}
	b.WriteString("```\n\nProvider prompts should prefer:\n\n")
	fmt.Fprintf(b, "```text\n%s\n```\n\n", m.ExamplePayload)
}
