package main

import "fmt"

func verifyProjection(model model) error {
	if !model.DSLIRMatch {
		return fmt.Errorf("DSL and IR configuration operation tuples differ")
	}
	if !model.OpenAPIMatch {
		return fmt.Errorf("OpenAPI configuration projection differs")
	}
	if !model.StreamPass {
		return fmt.Errorf("stream variant %s missing", model.Manifest.RequiredStreamVariant)
	}
	return nil
}
