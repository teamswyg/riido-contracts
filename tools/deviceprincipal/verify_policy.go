package main

import (
	"fmt"
	"strings"
)

func verifyPolicyPrefixes(model model) error {
	for _, prefix := range model.Manifest.PolicyRulePrefixes {
		if !hasRulePrefix(model.PolicyRules, prefix) {
			return fmt.Errorf("policy %s missing rule prefix %q",
				model.Manifest.APIPolicyID, prefix)
		}
	}
	return nil
}

func hasRulePrefix(rules []string, prefix string) bool {
	for _, rule := range rules {
		if strings.HasPrefix(rule, prefix) {
			return true
		}
	}
	return false
}
