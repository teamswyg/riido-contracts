package main

import (
	"errors"
	"fmt"
)

func validateFSMMetadata(spec fsmMetadata) (fsmMetadata, error) {
	if spec.Package == "" || spec.TransitionName == "" {
		return fsmMetadata{}, errors.New("transitions block missing package or name")
	}
	if spec.FromEnum == "" || spec.ToEnum == "" {
		return fsmMetadata{}, fmt.Errorf("transitions %s missing from-enum or to-enum", spec.TransitionName)
	}
	if missingFSMMetadataContractFields(spec) {
		return fsmMetadata{}, fmt.Errorf("transitions %s missing fsm-name, fsm-type-union, pattern-source, conformance-profile, or readme-section", spec.TransitionName)
	}
	if len(spec.Patterns) == 0 || len(spec.StartPoints) == 0 || len(spec.EndPoints) == 0 {
		return fsmMetadata{}, fmt.Errorf("transitions %s missing patterns, start-points, or end-points", spec.TransitionName)
	}
	if len(spec.Entries) == 0 {
		return fsmMetadata{}, fmt.Errorf("transitions %s has no transition entries", spec.TransitionName)
	}
	return spec, nil
}

func missingFSMMetadataContractFields(spec fsmMetadata) bool {
	return spec.FSMName == "" ||
		spec.TypeUnion == "" ||
		spec.PatternSource == "" ||
		spec.ConformanceProfile == "" ||
		spec.ReadmeSection == ""
}
