package main

import (
	"fmt"
	"strings"
)

func renderInlineList(b *strings.Builder, title string, values []string) {
	fmt.Fprintf(b, "- %s: ", title)
	for index, value := range values {
		if index > 0 {
			fmt.Fprint(b, ", ")
		}
		fmt.Fprintf(b, "`%s`", value)
	}
	fmt.Fprintln(b)
}

func renderDependencyPhrases(b *strings.Builder, model model) {
	fmt.Fprintln(b, "## Dependency Phrases")
	fmt.Fprintln(b)
	for _, phrase := range model.DependencyPhrases {
		fmt.Fprintf(b, "- %s\n", phrase)
	}
	fmt.Fprintln(b)
}
