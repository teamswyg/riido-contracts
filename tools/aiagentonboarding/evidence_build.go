package main

import "encoding/json"

func writeEvidence(path string, model model) error {
	m := model.Manifest
	body, err := json.MarshalIndent(evidence{
		SchemaVersion:           "riido-ai-agent-onboarding-evidence.v1",
		ID:                      m.ID,
		Status:                  "verified",
		GeneratedDoc:            m.GeneratedDoc,
		Workflow:                m.Workflow,
		EvidenceArtifact:        m.EvidenceArtifact,
		OperationCount:          len(model.Operations),
		OnboardingOpCount:       len(model.OnboardingOperations),
		DirectCreateOpCount:     len(model.DirectCreateOperations),
		FixtureFieldCount:       len(model.FixtureFields),
		CreateRequestFieldCount: len(model.CreateRequestFields),
		ScenarioCount:           model.ScenarioCount,
		FixtureRows:             m.FixtureRows,
		DSLIRMatch:              model.DSLIRMatch,
		OpenAPIMatch:            model.OpenAPIMatch,
		NoDiffPathsClean:        model.NoDiffPathsClean,
		Loop:                    m.Loop,
	}, "", "  ")
	if err != nil {
		return err
	}
	return writeFile(path, append(body, '\n'))
}
