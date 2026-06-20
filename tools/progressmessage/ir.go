package main

import (
	"os"

	"github.com/teamswyg/riido-contracts/progressmessage"
)

func buildIR(root string) (progressmessage.DSLDocument, progressmessage.IRDocument, error) {
	dsl, err := progressmessage.LoadDSL(os.DirFS(root), dslPath)
	if err != nil {
		return progressmessage.DSLDocument{}, progressmessage.IRDocument{}, err
	}
	ir, err := progressmessage.GenerateIR(dsl)
	if err != nil {
		return progressmessage.DSLDocument{}, progressmessage.IRDocument{}, err
	}
	return dsl, ir, nil
}
