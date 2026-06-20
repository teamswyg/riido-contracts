package main

import (
	"fmt"
	"strings"
)

func renderTransitions(b *strings.Builder, transitions []string) {
	b.WriteString("## Transition Events\n\n")
	fmt.Fprintf(b, "`%s`\n\n", strings.Join(transitions, "`, `"))
}
