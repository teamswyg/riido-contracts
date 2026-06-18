package main

import (
	"bytes"
	"fmt"
)

func writeNativeConfigRequirement(b *bytes.Buffer, enum enumSpec) {
	if enum.Package != "ir" || enum.Type != "EventType" {
		return
	}
	groups := nativeConfigGroups(enum)
	fmt.Fprintf(b, "func (code %s) NativeConfigRequirement() NativeConfigRequirement {\n", enum.CodeType)
	fmt.Fprintln(b, "\tswitch code {")
	for _, item := range nativeConfigOrder() {
		values := groups[item.Key]
		if len(values) == 0 {
			continue
		}
		writeCaseList(b, "\t", enumCodeRefs(enum, values))
		fmt.Fprintf(b, "\t\treturn %s\n", item.Go)
	}
	fmt.Fprintln(b, "\tdefault:")
	fmt.Fprintln(b, "\t\treturn NativeConfigRequired")
	fmt.Fprintln(b, "\t}")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintf(b, "func (value %s) NativeConfigRequirement() NativeConfigRequirement {\n", enum.Type)
	fmt.Fprintln(b, "\treturn value.Code().NativeConfigRequirement()")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
}
