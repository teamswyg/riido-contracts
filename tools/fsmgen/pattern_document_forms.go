package main

import (
	"errors"
	"fmt"
)

func applyPatternDocumentForm(doc *patternDocument, form node) error {
	switch atom(form.list[0]) {
	case "sum-type":
		sumType, err := parsePatternSumType(form)
		if err != nil {
			return err
		}
		doc.SumType = sumType
	case "conformance-profile":
		profile, err := parseConformanceProfile(form)
		if err != nil {
			return err
		}
		if _, ok := doc.Profiles[profile.Name]; ok {
			return fmt.Errorf("duplicate conformance profile %s", profile.Name)
		}
		doc.Profiles[profile.Name] = profile
	default:
		return fmt.Errorf("unknown fsm-pattern-gen form %q", atom(form.list[0]))
	}
	return nil
}

func validatePatternDocument(doc patternDocument) (patternDocument, error) {
	if doc.SumType.Type == "" || len(doc.SumType.Values) == 0 {
		return patternDocument{}, errors.New("pattern sum-type is required")
	}
	if len(doc.Profiles) == 0 {
		return patternDocument{}, errors.New("at least one conformance-profile is required")
	}
	return doc, nil
}
