package main

import (
	"fmt"
	"strings"
)

func renderVocabulary(b *strings.Builder, model model) {
	b.WriteString("## Vocabulary\n\n")
	fmt.Fprintf(b, "- Event stream formats: `%s`\n", strings.Join(model.EventStreamFormats, "`, `"))
	fmt.Fprintf(b, "- Protocol maturities: `%s`\n", strings.Join(model.ProtocolMaturities, "`, `"))
	fmt.Fprintf(b, "- Compatibility statuses: `%s`\n\n", strings.Join(model.CompatibilityStatuses, "`, `"))
}
