package main

import (
	"bytes"
	"fmt"
)

func writeEnumTypes(b *bytes.Buffer, enum enumSpec) {
	fmt.Fprintf(b, "type %s uint16\n\n", enum.CodeType)
	fmt.Fprintf(b, "type %s string\n", enum.StringType)
}

func writeEnumCodeConsts(b *bytes.Buffer, enum enumSpec) {
	fmt.Fprintln(b, "const (")
	fmt.Fprintf(b, "\t%sUnknown %s = iota\n", enum.CodeType, enum.CodeType)
	for _, value := range enum.Values {
		fmt.Fprintf(b, "\t%s\n", enum.codeConst(value.Const))
	}
	fmt.Fprintln(b, ")")
}

func writeEnumStringConsts(b *bytes.Buffer, enum enumSpec) {
	fmt.Fprintln(b, "const (")
	for _, value := range enum.Values {
		fmt.Fprintf(
			b,
			"\t%s %s = %q\n",
			enum.stringConst(value.Const),
			enum.StringType,
			value.Value,
		)
	}
	fmt.Fprintln(b, ")")
}

func writeEnumDomainConsts(b *bytes.Buffer, enum enumSpec) {
	fmt.Fprintln(b, "const (")
	for _, value := range enum.Values {
		fmt.Fprintf(
			b,
			"\t%s %s = %s(%s)\n",
			value.Const,
			enum.Type,
			enum.Type,
			enum.stringConst(value.Const),
		)
	}
	fmt.Fprintln(b, ")")
}
