package main

import (
	"path/filepath"
	"sort"
	"strings"
)

func buildBundle() (bundle, error) {
	root, err := repoRoot()
	if err != nil {
		return bundle{}, err
	}
	paths, err := filepath.Glob(filepath.Join(root, fixtureGlob))
	if err != nil {
		return bundle{}, err
	}
	sort.Strings(paths)
	contracts := make([]contract, 0, len(paths))
	for _, path := range paths {
		irPath, err := filepath.Rel(root, path)
		if err != nil {
			return bundle{}, err
		}
		item, err := buildContract(root, irPath)
		if err != nil {
			return bundle{}, err
		}
		contracts = append(contracts, item)
	}
	return bundle{
		SchemaVersion: bundleSchemaVersion,
		Source:        "apicontract DSL -> IR -> OpenAPI fixtures",
		Contracts:     contracts,
	}, nil
}

func openAPIPathForIR(irPath string) string {
	return strings.TrimSuffix(irPath, ".ir.riido.json") + ".openapi.json"
}
