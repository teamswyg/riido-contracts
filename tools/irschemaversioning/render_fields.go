package main

import (
	"fmt"
	"strings"
)

func renderFields(b *strings.Builder, model model) {
	b.WriteString("## Validator Surface\n\n")
	fmt.Fprintf(b, "- RunScope required identity/surface fields: `%d`\n", model.RunRequiredFieldCount)
	fmt.Fprintf(b, "- NativeConfigVersion classes: `%d`; RunContext fields: `%s`\n",
		model.NativeConfigClassCount, strings.Join(model.RunContextFieldNames, "`, `"))
	fmt.Fprintf(b, "- CanonicalEvent field names: `%s`\n\n",
		strings.Join(model.CanonicalEventFieldNames, "`, `"))
}
