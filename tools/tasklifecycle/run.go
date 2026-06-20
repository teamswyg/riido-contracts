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
	model, doc, err := build(root, opt)
	if err != nil {
		return err
	}
	if opt.writeDoc {
		if err := writeFile(resolve(root, model.Manifest.GeneratedDoc), []byte(doc)); err != nil {
			return err
		}
	}
	if opt.checkDoc {
		if err := verifyDoc(resolve(root, model.Manifest.GeneratedDoc), doc); err != nil {
			return err
		}
	}
	if opt.evidenceOut != "" {
		if err := writeEvidence(resolve(root, opt.evidenceOut), model); err != nil {
			return err
		}
	}
	fmt.Printf("tasklifecycle: verified states=%d transitions=%d\n", len(model.States), model.TransitionCount)
	return nil
}
