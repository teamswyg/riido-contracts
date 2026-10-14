package main

import (
	"flag"
	"io"
)

func run(args []string, out io.Writer) error {
	command, flagArgs := parseEnumCommand(args)
	if err := parseEnumFlags(flagArgs); err != nil {
		return err
	}
	plan, err := buildEnumRunPlan()
	if err != nil {
		return err
	}
	return runEnumCommand(command, plan, out)
}

func parseEnumFlags(args []string) error {
	fs := flag.NewFlagSet("enumgen", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	return fs.Parse(args)
}

func runEnumCommand(command enumCommand, plan enumRunPlan, out io.Writer) error {
	switch command {
	case enumCommandGenerate:
		return runEnumGenerate(plan, out)
	case enumCommandVerify:
		return runEnumVerify(plan, out)
	default:
		return unknownEnumCommandError()
	}
}
