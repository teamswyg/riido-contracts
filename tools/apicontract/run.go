package main

import "errors"

func run(args []string) error {
	command := "verify"
	if len(args) > 0 {
		command = args[0]
	}
	switch command {
	case "generate":
		return generate()
	case "verify":
		return verify()
	default:
		return errors.New("usage: go run ./tools/apicontract [verify|generate]")
	}
}
