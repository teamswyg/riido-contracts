package main

import "fmt"

func run(args []string) error {
	opt, err := parseOptions(args)
	if err != nil {
		return err
	}
	root, err := resolveRoot(opt.root)
	if err != nil {
		return err
	}
	m, summaries, doc, err := build(root, opt)
	if err != nil {
		return err
	}
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
		if err := writeEvidence(resolve(root, opt.evidenceOut), m, summaries); err != nil {
			return err
		}
	}
	fmt.Printf("apicontractdoc: verified fixtures=%d operations=%d\n", len(summaries), totalOperations(summaries))
	return nil
}
