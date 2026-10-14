package main

import (
	"fmt"
	"strings"
)

func renderScopes(b *strings.Builder, rules []scopeRule) {
	b.WriteString("## Scope Rules\n\n")
	b.WriteString("| Scope | Required | Forbidden | Conditional |\n| --- | ---: | ---: | ---: |\n")
	for _, rule := range rules {
		fmt.Fprintf(b, "| `%s` | `%d` | `%d` | `%d` |\n",
			rule.Scope, rule.Required, rule.Forbidden, rule.Conditional)
	}
	b.WriteString("\n")
}
