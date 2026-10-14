package main

import (
	"bytes"
	"fmt"
)

func writeEnumAllCodeValues(b *bytes.Buffer, enum enumSpec) {
	fmt.Fprintf(b, "var %s = []%s{\n", enumPrivateName(enum, "AllCodes"), enum.CodeType)
	for _, value := range enum.Values {
		fmt.Fprintf(b, "\t%s,\n", enum.codeConst(value.Const))
	}
	fmt.Fprintln(b, "}")
}

func writeEnumAllMethods(b *bytes.Buffer, enum enumSpec) {
	fmt.Fprintf(b, "func %s() []%s {\n", enum.CodeAllFunc, enum.CodeType)
	fmt.Fprintf(
		b,
		"\treturn append([]%s(nil), %s...)\n",
		enum.CodeType,
		enumPrivateName(enum, "AllCodes"),
	)
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintf(b, "func %s() []%s {\n", enum.AllFunc, enum.Type)
	fmt.Fprintf(b, "\tcodes := %s()\n", enum.CodeAllFunc)
	fmt.Fprintf(b, "\tout := make([]%s, len(codes))\n", enum.Type)
	fmt.Fprintln(b, "\tfor index, code := range codes {")
	fmt.Fprintf(b, "\t\tout[index] = code.%s()\n", enum.Type)
	fmt.Fprintln(b, "\t}")
	fmt.Fprintln(b, "\treturn out")
	fmt.Fprintln(b, "}")
}
