package main

import (
	"bytes"
	"fmt"
)

func writeEnumCodeMethods(b *bytes.Buffer, enum enumSpec) {
	fmt.Fprintf(b, "func (code %s) IsKnown() bool {\n", enum.CodeType)
	fmt.Fprintf(b, "\t_, ok := %s[code]\n", enumPrivateName(enum, "StringByCode"))
	fmt.Fprintln(b, "\treturn ok")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintf(b, "func (code %s) StringValue() %s {\n", enum.CodeType, enum.StringType)
	fmt.Fprintf(b, "\treturn %s[code]\n", enumPrivateName(enum, "StringByCode"))
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintf(b, "func (code %s) String() string {\n", enum.CodeType)
	fmt.Fprintln(b, "\treturn string(code.StringValue())")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintf(b, "func (code %s) %s() %s {\n", enum.CodeType, enum.Type, enum.Type)
	fmt.Fprintf(b, "\treturn %s(code.StringValue())\n", enum.Type)
	fmt.Fprintln(b, "}")
}

func writeEnumDomainMethods(b *bytes.Buffer, enum enumSpec) {
	fmt.Fprintf(b, "func (value %s) Valid() bool {\n", enum.Type)
	fmt.Fprintln(b, "\treturn value.Code().IsKnown()")
	fmt.Fprintln(b, "}")
}
