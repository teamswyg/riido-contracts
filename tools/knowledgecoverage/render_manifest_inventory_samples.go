package main

import (
	"fmt"
	"strings"
)

func renderManifestInventorySamples(b *strings.Builder, samples []manifestGroupSample) {
	b.WriteString("## Manifest Inventory Samples\n\n")
	b.WriteString("| Group | Sample paths |\n| --- | --- |\n")
	for _, sample := range samples {
		fmt.Fprintf(b, "| `%s` | %s |\n", sample.Group, renderSamplePaths(sample.Paths))
	}
	b.WriteString("\n")
}

func renderSamplePaths(paths []string) string {
	if len(paths) == 0 {
		return "None"
	}
	quoted := make([]string, 0, len(paths))
	for _, path := range paths {
		quoted = append(quoted, "`"+path+"`")
	}
	return strings.Join(quoted, "<br>")
}
