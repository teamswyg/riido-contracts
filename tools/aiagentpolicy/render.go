package main

import (
	"fmt"
	"strings"
)

func renderDoc(m manifest) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# %s\n\n%s\n\n", m.Title, generatedNotice)
	fmt.Fprintf(&b, "> Riido task: %s\n\n", m.RiidoTask)
	writeLines(&b, m.Intro)
	for _, section := range m.Sections {
		fmt.Fprintf(&b, "## %s\n", section.Title)
		writeLines(&b, section.Body)
	}
	return b.String()
}

func writeLines(b *strings.Builder, lines []string) {
	for _, line := range lines {
		fmt.Fprintf(b, "%s\n", line)
	}
}
