package main

import (
	"path/filepath"
	"sort"
)

func buildContract(root, irPath string) (contract, error) {
	var ir irDocument
	if err := readJSON(filepath.Join(root, irPath), &ir); err != nil {
		return contract{}, err
	}
	var dsl dslDocument
	if err := readJSON(filepath.Join(root, dslPathForIR(irPath)), &dsl); err != nil {
		return contract{}, err
	}
	openAPIPath := openAPIPathForIR(irPath)
	var openAPI openAPIDocument
	if err := readJSON(filepath.Join(root, openAPIPath), &openAPI); err != nil {
		return contract{}, err
	}
	ops := append([]operation(nil), ir.Operations...)
	attachScenarios(ops, dsl)
	for i := range ops {
		ops[i].PathParams = pathParams(ops[i].Path)
	}
	sort.Slice(ops, func(i, j int) bool {
		if ops[i].Path == ops[j].Path {
			return ops[i].Method < ops[j].Method
		}
		return ops[i].Path < ops[j].Path
	})
	return contract{
		ContractID:     ir.ContractID,
		Context:        ir.Context,
		Service:        ir.Service,
		SourceFiles:    sourceFiles{IR: irPath, OpenAPI: openAPIPath},
		OperationCount: len(ops),
		Operations:     ops,
		Schemas:        openAPI.Components.Schemas,
	}, nil
}
