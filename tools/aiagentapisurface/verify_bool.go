package main

import "errors"

func verifyBooleans(model model) error {
	if !model.DSLIRMatch {
		return errors.New("DSL and IR operation tuples differ")
	}
	if !model.IROpenAPIMatch {
		return errors.New("IR and OpenAPI operation tuples differ")
	}
	if !model.V2CoversV1 {
		return errors.New("v2 surface does not cover all v1 operations")
	}
	if !model.StreamVariantPass {
		return errors.New("ClientStreamEvent variants do not cover required variants")
	}
	return nil
}
