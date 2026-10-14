package main

import (
	"fmt"
	"strings"
)

func renderDirectionRules(b *strings.Builder, m manifest) {
	b.WriteString("## Direction Rules\n\n")
	for _, rule := range m.DirectionRules {
		fmt.Fprintf(b, "- %s\n", cell(rule))
	}
	b.WriteString("\n")
}

func renderSSOTLinks(b *strings.Builder, m manifest) {
	b.WriteString("## SSOT Links\n\n")
	for _, link := range m.SSOTLinks {
		fmt.Fprintf(b, "- %s: [`%s`](%s)\n", cell(link.Label), cell(link.Path), cell(link.Path))
	}
}
