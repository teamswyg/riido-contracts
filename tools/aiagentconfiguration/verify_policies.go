package main

import "fmt"

func verifyPolicies(model model) error {
	for _, id := range model.Manifest.RequiredPolicies {
		policy := findPolicy(model.Policies, id)
		if policy.PolicyID == "" {
			return fmt.Errorf("policy %s missing", id)
		}
		if len(policy.Rules) == 0 {
			return fmt.Errorf("policy %s has no rules", id)
		}
	}
	return nil
}
