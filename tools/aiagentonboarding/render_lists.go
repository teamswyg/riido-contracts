package main

import (
	"fmt"
	"strings"
)

func renderList(b *strings.Builder, title string, values []string) {
	fmt.Fprintf(b, "## %s\n\n", title)
	for _, value := range values {
		fmt.Fprintf(b, "- `%s`\n", value)
	}
	fmt.Fprintln(b)
}

func renderInlineList(b *strings.Builder, title string, values []string) {
	fmt.Fprintf(b, "## %s\n\n", title)
	for i, value := range values {
		if i > 0 {
			fmt.Fprint(b, ", ")
		}
		fmt.Fprintf(b, "`%s`", value)
	}
	fmt.Fprintln(b)
	fmt.Fprintln(b)
}

func renderFixtureRows(b *strings.Builder, rows []fixtureRow) {
	fmt.Fprintln(b, "## Product Fixture Rows")
	fmt.Fprintln(b)
	for _, row := range rows {
		fmt.Fprintf(b, "- `%s`: `%s`\n", row.Name, row.TmpColor)
	}
	fmt.Fprintln(b)
}
