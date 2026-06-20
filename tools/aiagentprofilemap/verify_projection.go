package main

import "fmt"

func verifyProjection(model model) error {
	if !model.DSLIRMatch {
		return fmt.Errorf("DSL and IR assigned-profile-map operation tuples differ")
	}
	if !model.OpenAPIMatch {
		return fmt.Errorf("OpenAPI assigned-profile-map projection differs")
	}
	return nil
}
