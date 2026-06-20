package main

import (
	"fmt"
)

func run(args []string) error {
	opt, err := parseOptions(args)
	if err != nil {
		return err
	}
	root, err := resolveRoot(opt.root)
	if err != nil {
		return err
	}
	m, c, doc, err := build(root, opt)
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
		if err := writeEvidence(resolve(root, opt.evidenceOut), m, c); err != nil {
			return err
		}
	}
	fmt.Printf("assignmentcontract: verified states=%d poll_actions=%d\n", len(c.AssignmentStates), len(c.PollActions))
	return nil
}
