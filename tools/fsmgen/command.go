package main

import (
	"errors"
	"strings"
)

type fsmCommand int

const (
	fsmCommandVerify fsmCommand = iota
	fsmCommandGenerate
	fsmCommandConformance
	fsmCommandUnknown
)

func parseFSMCommand(args []string) (fsmCommand, []string) {
	if len(args) == 0 || strings.HasPrefix(args[0], "-") {
		return fsmCommandVerify, args
	}
	switch args[0] {
	case "verify":
		return fsmCommandVerify, args[1:]
	case "generate":
		return fsmCommandGenerate, args[1:]
	case "conformance":
		return fsmCommandConformance, args[1:]
	default:
		return fsmCommandUnknown, args[1:]
	}
}

func unknownFSMCommandError() error {
	return errors.New("usage: go run ./tools/fsmgen [verify|generate|conformance]")
}
