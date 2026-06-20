package main

import "fmt"

func verifyProjection(model model) error {
	if !model.DSLIRMatch {
		return fmt.Errorf("DSL and IR assignment operation tuples differ")
	}
	if !model.OpenAPIMatch {
		return fmt.Errorf("OpenAPI assignment projection differs")
	}
	return nil
}
