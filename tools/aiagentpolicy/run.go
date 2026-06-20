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
	model, err := build(root, opt)
	if err != nil {
		return err
	}
	if opt.writeDoc {
		if err := writeFile(resolve(root, model.Manifest.GeneratedDoc), []byte(model.Document)); err != nil {
			return err
		}
	}
	if opt.checkDoc {
		if err := verifyDoc(resolve(root, model.Manifest.GeneratedDoc), model.Document); err != nil {
			return err
		}
	}
	if opt.evidenceOut != "" {
		if err := writeEvidence(resolve(root, opt.evidenceOut), model); err != nil {
			return err
		}
	}
	fmt.Printf("aiagentpolicy: verified sections=%d assertions=%d\n",
		len(model.Manifest.Sections), model.PolicyAssertionCount)
	return nil
}
