package main

import (
	"errors"
	"strings"
)

type enumCommand int

const (
	enumCommandVerify enumCommand = iota
	enumCommandGenerate
	enumCommandUnknown
)

func parseEnumCommand(args []string) (enumCommand, []string) {
	if len(args) == 0 || strings.HasPrefix(args[0], "-") {
		return enumCommandVerify, args
	}
	switch args[0] {
	case "verify":
		return enumCommandVerify, args[1:]
	case "generate":
		return enumCommandGenerate, args[1:]
	default:
		return enumCommandUnknown, args[1:]
	}
}

func unknownEnumCommandError() error {
	return errors.New("usage: go run ./tools/enumgen [verify|generate]")
}
