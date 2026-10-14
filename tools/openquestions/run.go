package main

import (
	"flag"
	"fmt"
)

type options struct {
	root        string
	manifest    string
	writeDoc    bool
	checkDoc    bool
	evidenceOut string
}

func run(args []string) error {
	var opt options
	fs := flag.NewFlagSet("openquestions", flag.ContinueOnError)
	fs.StringVar(&opt.root, "root", ".", "repository root")
	fs.StringVar(&opt.manifest, "manifest", defaultManifest, "manifest path")
	fs.BoolVar(&opt.writeDoc, "write-doc", false, "write generated reader")
	fs.BoolVar(&opt.checkDoc, "check-doc", false, "verify generated reader")
	fs.StringVar(&opt.evidenceOut, "evidence-out", "", "write evidence JSON")
	if err := fs.Parse(args); err != nil {
		return err
	}
	root, err := resolveRoot(opt.root)
	if err != nil {
		return err
	}
	m, err := loadManifest(resolve(root, opt.manifest))
	if err != nil {
		return err
	}
	if err := verifyManifest(root, m); err != nil {
		return err
	}
	doc := renderDoc(m)
	if opt.writeDoc {
		if err := writeFile(resolve(root, m.GeneratedDoc), []byte(doc)); err != nil {
			return err
		}
	}
	if opt.checkDoc {
		if err := verifyDoc(resolve(root, m.GeneratedDoc), doc); err != nil {
			return err
		}
	}
	if opt.evidenceOut != "" {
		if err := writeEvidence(resolve(root, opt.evidenceOut), m); err != nil {
			return err
		}
	}
	fmt.Printf("openquestions: %s open=%d resolved=%d\n", status(m), countStatus(m, "open"), countStatus(m, "resolved"))
	return nil
}
