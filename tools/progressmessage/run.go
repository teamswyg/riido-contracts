package main

import (
	"errors"
	"flag"
	"strings"
)

func run(args []string) error {
	command := "verify"
	if len(args) > 0 && !strings.HasPrefix(args[0], "-") {
		command = args[0]
		args = args[1:]
	}
	switch command {
	case "generate":
		return generate()
	case "verify":
		opt, err := parseOptions(args)
		if err != nil {
			return err
		}
		return verify(opt)
	default:
		return errors.New("usage: go run ./tools/progressmessage [verify|generate] [flags]")
	}
}

func parseOptions(args []string) (options, error) {
	opt := options{root: ".", manifest: docManifestPath}
	fs := flag.NewFlagSet("progressmessage", flag.ContinueOnError)
	fs.StringVar(&opt.root, "root", opt.root, "repository root")
	fs.StringVar(&opt.manifest, "manifest", opt.manifest, "doc manifest path")
	fs.BoolVar(&opt.writeDoc, "write-doc", false, "write generated reader")
	fs.BoolVar(&opt.checkDoc, "check-doc", false, "verify generated reader")
	fs.StringVar(&opt.evidenceOut, "evidence-out", "", "write evidence JSON")
	return opt, fs.Parse(args)
}
