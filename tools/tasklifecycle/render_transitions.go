package main

import (
	"fmt"
	"strings"
)

func renderTransitions(b *strings.Builder, groups []transitionGroup) {
	b.WriteString("## Transition Surface\n\n")
	b.WriteString("| From | Trigger to next state |\n| --- | --- |\n")
	for _, group := range groups {
		var edges []string
		for _, edge := range group.Edges {
			edges = append(edges, "`"+renderEdge(edge)+"`")
		}
		fmt.Fprintf(b, "| `%s` | %s |\n", group.From, strings.Join(edges, "<br>"))
	}
	b.WriteString("\n")
}
