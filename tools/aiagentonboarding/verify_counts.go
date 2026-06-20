package main

func verifyCounts(model model) error {
	m := model.Manifest
	checks := map[string][2]int{
		"operations":             {len(model.Operations), m.ExpectedOperationCount},
		"onboarding operations":  {len(model.OnboardingOperations), m.ExpectedOnboardingOpCount},
		"direct create ops":      {len(model.DirectCreateOperations), m.ExpectedDirectCreateOpCount},
		"fixture fields":         {len(model.FixtureFields), m.ExpectedFixtureFieldCount},
		"create request fields":  {len(model.CreateRequestFields), m.ExpectedCreateRequestFieldCount},
		"scenario count":         {model.ScenarioCount, m.ExpectedScenarioCount},
		"fixture row count":      {len(m.FixtureRows), 4},
		"required op count":      {len(m.RequiredOperations), m.ExpectedOperationCount},
		"no-diff fragment count": {len(m.NoDiffPathFragments), 3},
	}
	for name, values := range checks {
		if err := countMismatch(name, values[0], values[1]); err != nil {
			return err
		}
	}
	return nil
}
