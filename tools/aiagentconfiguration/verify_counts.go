package main

func verifyCounts(model model) error {
	m := model.Manifest
	checks := map[string][2]int{
		"operations":        {len(model.Operations), m.ExpectedOperationCount},
		"schemas":           {len(model.Schemas), m.ExpectedSchemaCount},
		"policies":          {len(model.Policies), m.ExpectedPolicyCount},
		"enums":             {len(model.Enums), m.ExpectedEnumCount},
		"scenarios":         {model.ScenarioCount, m.ExpectedScenarioCount},
		"bounded fields":    {len(m.BoundedFields), 4},
		"invariants":        {len(m.Invariants), 4},
		"required op count": {len(m.RequiredOperations), m.ExpectedOperationCount},
	}
	for name, values := range checks {
		if err := countMismatch(name, values[0], values[1]); err != nil {
			return err
		}
	}
	return nil
}
