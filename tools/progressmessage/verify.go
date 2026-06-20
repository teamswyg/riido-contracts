package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/teamswyg/riido-contracts/progressmessage"
)

func verify(opt options) error {
	root, err := resolveRoot(opt.root)
	if err != nil {
		return err
	}
	dsl, ir, err := verifyIR(root)
	if err != nil {
		return err
	}
	m, err := loadDocManifest(resolve(root, opt.manifest))
	if err != nil {
		return err
	}
	if err := verifyManifest(root, m); err != nil {
		return err
	}
	doc := renderDoc(m, dsl, ir)
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
		if err := writeEvidence(resolve(root, opt.evidenceOut), m, ir); err != nil {
			return err
		}
	}
	fmt.Printf("progressmessage: %s messages=%d\n", status(ir), len(ir.Messages))
	return nil
}

func verifyIR(root string) (progressmessage.DSLDocument, progressmessage.IRDocument, error) {
	dsl, ir, want, err := buildIR(root)
	if err != nil {
		return progressmessage.DSLDocument{}, progressmessage.IRDocument{}, err
	}
	got, err := os.ReadFile(resolve(root, irPath))
	if err != nil {
		return progressmessage.DSLDocument{}, progressmessage.IRDocument{}, fmt.Errorf("read %s: %w", irPath, err)
	}
	if !bytes.Equal(got, want) {
		return progressmessage.DSLDocument{}, progressmessage.IRDocument{}, fmt.Errorf("%s drifted; run go run ./tools/progressmessage generate", irPath)
	}
	return dsl, ir, nil
}
