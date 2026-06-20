package main

import "flag"

func parseOptions(args []string) (options, error) {
	fs := flag.NewFlagSet("providercapability", flag.ContinueOnError)
	opt := options{}
	fs.StringVar(&opt.root, "repo", ".", "repository root")
	fs.StringVar(&opt.manifest, "manifest", defaultManifest, "provider capability manifest")
	fs.StringVar(&opt.evidenceOut, "evidence-out", "", "write evidence JSON")
	fs.BoolVar(&opt.writeDoc, "write-doc", false, "write generated markdown")
	fs.BoolVar(&opt.checkDoc, "check-doc", false, "verify generated markdown")
	if err := fs.Parse(args); err != nil {
		return options{}, err
	}
	return opt, nil
}
