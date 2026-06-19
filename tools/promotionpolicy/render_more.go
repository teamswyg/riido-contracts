package main

import (
	"fmt"
	"strings"
)

func renderRuntimeTags(b *strings.Builder, m manifest) {
	b.WriteString("## Runtime Release Tags\n\n")
	fmt.Fprintf(b, "%s\n\n", cell(m.RuntimeTagRule))
	b.WriteString("| Runtime tag | Meaning |\n| --- | --- |\n")
	for _, tag := range m.RuntimeTagModel {
		fmt.Fprintf(b, "| `%s` | %s |\n", cell(tag.Pattern), cell(tag.Meaning))
	}
	b.WriteString("\n")
}

func renderBreakingRules(b *strings.Builder, m manifest) {
	b.WriteString("## Breaking Change Rules\n\n")
	b.WriteString("Breaking changes require an explicit migration slice:\n\n")
	for _, rule := range m.BreakingChangeRules {
		fmt.Fprintf(b, "- %s\n", cell(rule))
	}
	fmt.Fprintf(b, "\n%s\n\n", cell(m.AdditiveChangeRule))
}

func renderDownstreamRule(b *strings.Builder, m manifest) {
	b.WriteString("## Downstream Import Rule\n\n")
	fmt.Fprintf(b, "%s\n", cell(m.DownstreamImportRule))
}
