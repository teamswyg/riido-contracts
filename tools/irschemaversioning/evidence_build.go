package main

import (
	"encoding/json"
	"os"
)

func newEvidence(model model) evidence {
	m := model.Manifest
	return evidence{
		SchemaVersion:          "riido-contracts-ir-schema-versioning-evidence.v1",
		ID:                     m.ID,
		Status:                 "verified",
		GeneratedDoc:           m.GeneratedDoc,
		Package:                m.Package,
		Workflow:               m.Workflow,
		EvidenceArtifact:       m.EvidenceArtifact,
		CanonicalEventFields:   model.CanonicalEventFields,
		EventScopeCount:        model.EventScopeCount,
		CommonRequiredCount:    model.CommonRequiredCount,
		RunRequiredFieldCount:  model.RunRequiredFieldCount,
		FakePlaceholderFields:  model.FakePlaceholderFields,
		FakePlaceholderValues:  model.FakePlaceholderValues,
		ViolationCodeCount:     model.ViolationCodeCount,
		NativeConfigClassCount: model.NativeConfigClassCount,
		RunContextFields:       model.RunContextFields,
		ValidateEntrypoints:    model.ValidateEntrypoints,
		ScopeRules:             model.ScopeRules,
		Loop:                   m.Loop,
	}
}

func writeEvidence(path string, model model) error {
	body, err := json.MarshalIndent(newEvidence(model), "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(body, '\n'), 0o644)
}
