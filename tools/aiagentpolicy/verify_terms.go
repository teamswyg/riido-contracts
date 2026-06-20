package main

import (
	"fmt"
	"strings"
)

func verifyTerms(m manifest) error {
	termSection, ok := findSection(m, "Ubiquitous Language")
	if !ok {
		return fmt.Errorf("missing Ubiquitous Language section")
	}
	body := strings.Join(termSection.Body, "\n")
	for _, term := range m.RequiredTerms {
		if !strings.Contains(body, "| "+term+" |") {
			return fmt.Errorf("required term %q is missing from table", term)
		}
	}
	return nil
}

func findSection(m manifest, title string) (policySection, bool) {
	for _, section := range m.Sections {
		if section.Title == title {
			return section, true
		}
	}
	return policySection{}, false
}
