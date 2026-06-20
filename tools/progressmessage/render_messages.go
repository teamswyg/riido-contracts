package main

import (
	"fmt"
	"strings"

	"github.com/teamswyg/riido-contracts/progressmessage"
)

func renderMessages(b *strings.Builder, messages []progressmessage.MessageDefinition) {
	b.WriteString("## Message Catalog\n\n")
	b.WriteString("| Code | Key | Usage | Category | Args | Korean copy |\n")
	b.WriteString("| --- | --- | --- | --- | --- | --- |\n")
	for _, message := range messages {
		fmt.Fprintf(b, "| `%d` | `%s` | `%s` | `%s` | %s | %s |\n",
			message.Code,
			message.Key,
			message.Usage,
			message.Category,
			cell(renderArgs(message.Args)),
			cell(message.Locales[progressmessage.DefaultLocale]),
		)
	}
	b.WriteString("\n")
}

func renderArgs(args []progressmessage.MessageArg) string {
	if len(args) == 0 {
		return "-"
	}
	parts := make([]string, 0, len(args))
	for _, arg := range args {
		parts = append(parts, arg.Name+":"+arg.Type)
	}
	return strings.Join(parts, ", ")
}
