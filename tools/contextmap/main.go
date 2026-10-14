package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "contextmap:", err)
		os.Exit(1)
	}
}

func run(args []string, out io.Writer) error {
	command := "verify"
	if len(args) > 0 && !strings.HasPrefix(args[0], "-") {
		command = args[0]
		args = args[1:]
	}
	switch command {
	case "render":
		return runRender(args, out)
	case "verify":
		return runVerify(args, out)
	case "write-doc":
		return runWriteDoc(args)
	default:
		return fmt.Errorf("usage: go run ./tools/contextmap [render|verify|write-doc]")
	}
}
