package main

import (
	"flag"
	"io"
)

func run(args []string, out io.Writer) error {
	command, flagArgs := parseFSMCommand(args)
	if err := parseFSMFlags(flagArgs); err != nil {
		return err
	}
	plan, err := buildFSMRunPlan()
	if err != nil {
		return err
	}
	return runFSMCommand(command, plan, out)
}

func parseFSMFlags(args []string) error {
	fs := flag.NewFlagSet("fsmgen", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	return fs.Parse(args)
}

func runFSMCommand(command fsmCommand, plan fsmRunPlan, out io.Writer) error {
	switch command {
	case fsmCommandConformance:
		return runFSMConformance(plan, out)
	case fsmCommandGenerate:
		return runFSMGenerate(plan, out)
	case fsmCommandVerify:
		return runFSMVerify(plan, out)
	default:
		return unknownFSMCommandError()
	}
}
