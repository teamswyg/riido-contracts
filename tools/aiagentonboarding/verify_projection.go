package main

import "fmt"

func verifyProjectionFlags(model model) error {
	if !model.DSLIRMatch {
		return fmt.Errorf("DSL and IR onboarding operation tuples differ")
	}
	if !model.OpenAPIMatch {
		return fmt.Errorf("OpenAPI onboarding projection differs")
	}
	if !model.NoDiffPathsClean {
		return fmt.Errorf("no-diff provider/waitlist path boundary drifted")
	}
	return nil
}
