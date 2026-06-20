package main

import "strings"

func policyAssertions(m manifest) []string {
	section, ok := findSection(m, "Policy Assertions")
	if !ok {
		return nil
	}
	assertions := make([]string, 0, len(section.Subsections))
	for _, line := range section.Body {
		if strings.HasPrefix(line, "### ") {
			assertions = append(assertions, strings.TrimPrefix(line, "### "))
		}
	}
	return assertions
}
