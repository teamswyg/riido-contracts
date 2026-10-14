package main

import "fmt"

func verifyCounts(model model) error {
	m := model.Manifest
	if count := boolCount(model.Operation.OperationID != ""); count != m.ExpectedOperationCount {
		return fmt.Errorf("operation count = %d, want %d", count, m.ExpectedOperationCount)
	}
	if len(model.Schemas) != m.ExpectedSchemaCount {
		return fmt.Errorf("schema count = %d, want %d", len(model.Schemas), m.ExpectedSchemaCount)
	}
	if len(model.Policy.Rules) != m.ExpectedPolicyRuleCount {
		return fmt.Errorf("policy rule count = %d, want %d", len(model.Policy.Rules), m.ExpectedPolicyRuleCount)
	}
	if model.ScenarioCount != m.ExpectedScenarioCount {
		return fmt.Errorf("scenario count = %d, want %d", model.ScenarioCount, m.ExpectedScenarioCount)
	}
	return nil
}

func boolCount(ok bool) int {
	if ok {
		return 1
	}
	return 0
}
