package main

import (
	"bytes"
	"fmt"
)

func writeCaseList(b *bytes.Buffer, indent string, refs []string) {
	fmt.Fprintf(b, "%scase ", indent)
	for index, ref := range refs {
		if index > 0 {
			fmt.Fprint(b, ", ")
		}
		fmt.Fprint(b, ref)
	}
	fmt.Fprintln(b, ":")
}
