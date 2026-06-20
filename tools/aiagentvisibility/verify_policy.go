package main

import "fmt"

func verifyPolicy(model model) error {
	if model.Policy.PolicyID != model.Manifest.PolicyID {
		return fmt.Errorf("policy %s missing", model.Manifest.PolicyID)
	}
	if model.Policy.Kind != "rbac" {
		return fmt.Errorf("policy %s kind = %s, want rbac", model.Policy.PolicyID, model.Policy.Kind)
	}
	for _, rule := range model.Manifest.RequiredPolicyRules {
		if !hasRule(model.Policy, rule) {
			return fmt.Errorf("policy %s rule %q missing", model.Policy.PolicyID, rule)
		}
	}
	return nil
}
