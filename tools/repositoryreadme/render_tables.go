package main

import (
	"fmt"
	"strings"
)

func renderDocLinks(b *strings.Builder, links []docLink) {
	b.WriteString("## 어떤 문서를 보면 되나\n\n")
	b.WriteString("| 알고 싶은 것 | 문서 |\n| --- | --- |\n")
	for _, link := range links {
		fmt.Fprintf(b, "| %s | [`%s`](%s) |\n", link.Topic, link.Path, link.Path)
	}
	b.WriteByte('\n')
}

func renderPackages(b *strings.Builder, packages []packageRef) {
	b.WriteString("## 현재 package\n\n")
	for _, pkg := range packages {
		fmt.Fprintf(b, "- `%s`: %s\n", pkg.Name, pkg.Description)
	}
	b.WriteByte('\n')
}
