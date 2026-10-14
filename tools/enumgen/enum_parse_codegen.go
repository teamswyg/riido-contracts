package main

import (
	"bytes"
	"fmt"
)

func writeEnumParseMap(b *bytes.Buffer, enum enumSpec) {
	fmt.Fprintf(
		b,
		"var %s = map[%s]%s{\n",
		enumPrivateName(enum, "CodeByString"),
		enum.StringType,
		enum.CodeType,
	)
	for _, value := range enum.Values {
		fmt.Fprintf(b, "\t%s: %s,\n", enum.stringConst(value.Const), enum.codeConst(value.Const))
	}
	fmt.Fprintln(b, "}")
}

func writeEnumParseMethods(b *bytes.Buffer, enum enumSpec) {
	fmt.Fprintf(b, "func (value %s) Code() %s {\n", enum.Type, enum.CodeType)
	fmt.Fprintf(b, "\treturn Parse%sCode(string(value))\n", enum.Type)
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintf(b, "func Parse%sCode(value string) %s {\n", enum.Type, enum.CodeType)
	fmt.Fprintf(b, "\tcode, ok := %s[%s(value)]\n", enumPrivateName(enum, "CodeByString"), enum.StringType)
	fmt.Fprintln(b, "\tif !ok {")
	fmt.Fprintf(b, "\t\treturn %sUnknown\n", enum.CodeType)
	fmt.Fprintln(b, "\t}")
	fmt.Fprintln(b, "\treturn code")
	fmt.Fprintln(b, "}")
}

func writeEnumStringMap(b *bytes.Buffer, enum enumSpec) {
	fmt.Fprintf(
		b,
		"var %s = map[%s]%s{\n",
		enumPrivateName(enum, "StringByCode"),
		enum.CodeType,
		enum.StringType,
	)
	for _, value := range enum.Values {
		fmt.Fprintf(b, "\t%s: %s,\n", enum.codeConst(value.Const), enum.stringConst(value.Const))
	}
	fmt.Fprintln(b, "}")
}
