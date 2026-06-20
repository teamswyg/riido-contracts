package main

func addPolicies(model model, dsl contractFixture) model {
	for _, policyID := range model.Manifest.RequiredPolicies {
		model.Policies = append(model.Policies, findPolicy(dsl.Policies, policyID))
	}
	return model
}

func findPolicy(policies []policySpec, id string) policySpec {
	for _, policy := range policies {
		if policy.PolicyID == id {
			return policy
		}
	}
	return policySpec{}
}
