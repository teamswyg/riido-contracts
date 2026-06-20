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
	if len(model.Policy.Rules) != m.ExpectedPolicyRuleCount {
		return fmt.Errorf("policy rule count = %d, want %d", len(model.Policy.Rules), m.ExpectedPolicyRuleCount)
	}
	if len(model.Enum.Values) != m.ExpectedEnumValueCount {
		return fmt.Errorf("enum value count = %d, want %d", len(model.Enum.Values), m.ExpectedEnumValueCount)
	}
	if model.ScenarioCount != m.ExpectedScenarioCount {
		return fmt.Errorf("scenario count = %d, want %d", model.ScenarioCount, m.ExpectedScenarioCount)
	}
	return nil
}
