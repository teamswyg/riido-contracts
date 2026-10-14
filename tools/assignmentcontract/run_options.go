package main

import "flag"

func parseOptions(args []string) (options, error) {
	opt := options{root: ".", manifest: defaultManifest}
	fs := flag.NewFlagSet("assignmentcontract", flag.ContinueOnError)
	fs.StringVar(&opt.root, "root", opt.root, "repository root")
	fs.StringVar(&opt.manifest, "manifest", opt.manifest, "doc manifest path")
	fs.BoolVar(&opt.writeDoc, "write-doc", false, "write generated reader")
	fs.BoolVar(&opt.checkDoc, "check-doc", false, "verify generated reader")
	fs.StringVar(&opt.evidenceOut, "evidence-out", "", "write evidence JSON")
	return opt, fs.Parse(args)
}
