package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func verifyPolicyShape(m manifest) error {
	if !reflect.DeepEqual(sectionTitles(m), m.RequiredSectionOrder) {
		return errors.New("section order does not match required_section_order")
	}
	if !reflect.DeepEqual(policyAssertions(m), m.RequiredPolicyAssertions) {
		return errors.New("policy assertions do not match required_policy_assertions")
	}
	if len(m.RequiredTerms) == 0 || len(m.RequiredGeneratedReaders) == 0 {
		return errors.New("required terms and generated readers are required")
	}
	for _, section := range m.Sections {
		if blank(section.Title) || len(section.Body) == 0 {
			return fmt.Errorf("section %q is incomplete", section.Title)
		}
		if sectionHasTopLevelHeading(section) {
			return fmt.Errorf("section %q embeds a top-level heading", section.Title)
		}
	}
	return verifyTerms(m)
}

func sectionTitles(m manifest) []string {
	titles := make([]string, 0, len(m.Sections))
	for _, section := range m.Sections {
		titles = append(titles, section.Title)
	}
	return titles
}

func sectionHasTopLevelHeading(section policySection) bool {
	for _, line := range section.Body {
		if strings.HasPrefix(strings.TrimSpace(line), "## ") {
			return true
		}
	}
	return false
}
