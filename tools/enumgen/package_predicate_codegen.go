package main

import (
	"bytes"
	"fmt"
)

func writePackagePredicate(b *bytes.Buffer, enum enumSpec, attr, method string) {
	if enum.Package != "assignment" || len(enum.valuesWithAttr(attr, "true")) == 0 {
		return
	}
	fmt.Fprintf(b, "func %s(value %s) bool {\n", method, enum.Type)
	fmt.Fprintf(b, "\treturn value.Code().%s()\n", method)
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
}
