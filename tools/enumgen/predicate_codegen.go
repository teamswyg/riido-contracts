package main

import (
	"bytes"
	"fmt"
)

func writePredicate(b *bytes.Buffer, enum enumSpec, attr, method string) {
	values := enum.valuesWithAttr(attr, "true")
	if len(values) == 0 {
		return
	}
	fmt.Fprintf(b, "func (code %s) %s() bool {\n", enum.CodeType, method)
	fmt.Fprintln(b, "\tswitch code {")
	writeCaseList(b, "\t", enumCodeRefs(enum, values))
	fmt.Fprintln(b, "\t\treturn true")
	fmt.Fprintln(b, "\tdefault:")
	fmt.Fprintln(b, "\t\treturn false")
	fmt.Fprintln(b, "\t}")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintf(b, "func (value %s) %s() bool {\n", enum.Type, method)
	fmt.Fprintf(b, "\treturn value.Code().%s()\n", method)
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
}
