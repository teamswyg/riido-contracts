package main

import "fmt"

func verifyCounts(model model) error {
	m := model.Manifest
	if len(model.Operations) != m.ExpectedOperationCount {
		return fmt.Errorf("operation count = %d, want %d", len(model.Operations), m.ExpectedOperationCount)
	}
	if len(model.Schemas) != m.ExpectedSchemaCount {
		return fmt.Errorf("schema count = %d, want %d", len(model.Schemas), m.ExpectedSchemaCount)
	}
	if len(model.Policies) != m.ExpectedPolicyCount {
		return fmt.Errorf("policy count = %d, want %d", len(model.Policies), m.ExpectedPolicyCount)
	}
	if model.ScenarioCount != m.ExpectedScenarioCount {
		return fmt.Errorf("scenario count = %d, want %d", model.ScenarioCount, m.ExpectedScenarioCount)
	}
	return nil
}
