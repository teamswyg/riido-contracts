package main

import "fmt"

func verifyProjection(model model) error {
	if !model.DSLIRMatch {
		return fmt.Errorf("DSL and IR visibility operation tuples differ")
	}
	if !model.OpenAPIMatch {
		return fmt.Errorf("OpenAPI visibility projection differs")
	}
	return nil
}
