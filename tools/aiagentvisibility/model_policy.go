package main

func findPolicy(policies []policySpec, id string) policySpec {
	for _, policy := range policies {
		if policy.PolicyID == id {
			return policy
		}
	}
	return policySpec{}
}

func hasRule(policy policySpec, rule string) bool {
	for _, got := range policy.Rules {
		if got == rule {
			return true
		}
	}
	return false
}
