package main

import (
	"fmt"
	"path/filepath"

	"github.com/teamswyg/riido-contracts/apicontract"
)

func storeGeneratedFixture(
	generated map[string][]byte,
	dslPath string,
	ir apicontract.IRDocument,
	openAPI apicontract.OpenAPISpec,
) error {
	irJSON, err := apicontract.MarshalCanonical(ir)
	if err != nil {
		return fmt.Errorf("marshal IR: %w", err)
	}
	openAPIJSON, err := apicontract.MarshalCanonical(openAPI)
	if err != nil {
		return fmt.Errorf("marshal OpenAPI: %w", err)
	}
	stem := filepath.Clean(dslPath[:len(dslPath)-len(".dsl.riido.json")])
	generated[stem+".ir.riido.json"] = irJSON
	generated[stem+".openapi.json"] = openAPIJSON
	return nil
}
