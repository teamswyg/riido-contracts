package main

import (
	"fmt"
	"strings"
)

func renderRules(b *strings.Builder, rules []string) {
	b.WriteString("## Rules\n\n")
	for _, rule := range rules {
		fmt.Fprintf(b, "- %s\n", rule)
	}
	b.WriteString("\n")
}
