package main

import (
	"fmt"
	"strings"

	"github.com/teamswyg/riido-contracts/progressmessage"
)

func irMessageFile(message progressmessage.MessageDefinition) string {
	key := strings.NewReplacer(".", "-", "_", "-").Replace(message.Key)
	return fmt.Sprintf("%s/%04d-%s.ir.riido.json", irMessageDir, message.Code, key)
}

func irMessageFileRefs(messages []progressmessage.MessageDefinition) []string {
	refs := make([]string, 0, len(messages))
	for _, message := range messages {
		refs = append(refs, strings.TrimPrefix(irMessageFile(message), "progressmessage/"))
	}
	return refs
}
