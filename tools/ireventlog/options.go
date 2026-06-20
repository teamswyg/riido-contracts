package main

import "flag"

type options struct {
	root        string
	manifest    string
	writeDoc    bool
	checkDoc    bool
	evidenceOut string
}

func parseOptions(args []string) (options, error) {
	var opt options
	fs := flag.NewFlagSet("ireventlog", flag.ContinueOnError)
	fs.StringVar(&opt.root, "root", ".", "repository root")
	fs.StringVar(&opt.manifest, "manifest", defaultManifest, "manifest path")
	fs.BoolVar(&opt.writeDoc, "write-doc", false, "write generated markdown")
	fs.BoolVar(&opt.checkDoc, "check-doc", false, "verify generated markdown")
	fs.StringVar(&opt.evidenceOut, "evidence-out", "", "write evidence JSON")
	return opt, fs.Parse(args)
}
