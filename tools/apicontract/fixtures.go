package main

import (
	"fmt"
	"path/filepath"

	"github.com/teamswyg/riido-contracts/apicontract"
)

func generatedFixtures() (map[string][]byte, error) {
	dslPaths, err := filepath.Glob(dslGlob)
	if err != nil {
		return nil, err
	}
	if len(dslPaths) == 0 {
		return nil, fmt.Errorf("no DSL fixtures match %s", dslGlob)
	}
	generated := map[string][]byte{}
	for _, dslPath := range dslPaths {
		if err := addGeneratedFixture(generated, dslPath); err != nil {
			return nil, err
		}
	}
	return generated, nil
}

func addGeneratedFixture(generated map[string][]byte, dslPath string) error {
	dsl, err := loadDSL(dslPath)
	if err != nil {
		return err
	}
	ir, err := apicontract.GenerateIR(dsl)
	if err != nil {
		return err
	}
	openAPI, err := apicontract.GenerateOpenAPI(ir)
	if err != nil {
		return err
	}
	return storeGeneratedFixture(generated, dslPath, ir, openAPI)
}
