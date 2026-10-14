package main

import (
	"fmt"
	"strings"
)

func cell(value string) string {
	value = strings.ReplaceAll(value, "\n", " ")
	return strings.ReplaceAll(value, "|", "\\|")
}

func writeBullets(b *strings.Builder, values []string) {
	for _, value := range values {
		b.WriteString("- " + value + "\n")
	}
	b.WriteString("\n")
}

func writeNumbered(b *strings.Builder, values []string) {
	for i, value := range values {
		fmt.Fprintf(b, "%d. %s\n", i+1, value)
	}
	b.WriteString("\n")
}
