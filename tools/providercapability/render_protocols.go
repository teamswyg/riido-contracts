package main

import (
	"fmt"
	"strings"
)

func renderProtocols(b *strings.Builder, protocols []protocolRow) {
	b.WriteString("## Protocol Critical Args\n\n")
	b.WriteString("| ProtocolKind | Args |\n| --- | --- |\n")
	for _, row := range protocols {
		args := "none"
		if len(row.Args) > 0 {
			args = "`" + strings.Join(row.Args, "`, `") + "`"
		}
		fmt.Fprintf(b, "| `%s` | %s |\n", row.Kind, args)
	}
	b.WriteString("\n")
}
