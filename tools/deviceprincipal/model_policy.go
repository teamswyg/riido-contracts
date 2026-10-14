package main

import "fmt"

var apiFixturePaths = []string{
	"apicontract/fixtures/control-plane-ai-agent-client.dsl.riido.json",
	"apicontract/fixtures/control-plane-ai-agent-client.ir.riido.json",
}

func loadPolicyRules(root, policyID string) ([]string, error) {
	var expected []string
	for _, path := range apiFixturePaths {
		doc, err := readLooseJSONFile[apiContractDoc](resolve(root, path))
		if err != nil {
			return nil, err
		}
		policy, ok := findPolicy(doc.Policies, policyID)
		if !ok {
			return nil, fmt.Errorf("%s missing policy %s", path, policyID)
		}
		if expected == nil {
			expected = policy.Rules
			continue
		}
		if !sameStrings(expected, policy.Rules) {
			return nil, fmt.Errorf("%s policy rules drift from dsl fixture", path)
		}
	}
	return expected, nil
}

func findPolicy(policies []apiPolicy, policyID string) (apiPolicy, bool) {
	for _, policy := range policies {
		if policy.PolicyID == policyID {
			return policy, true
		}
	}
	return apiPolicy{}, false
}
