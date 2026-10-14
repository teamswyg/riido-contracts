package main

import "encoding/json"

func writeEvidence(path string, model model) error {
	m := model.Manifest
	body, err := json.MarshalIndent(evidence{
		SchemaVersion:    "riido-ai-agent-assigned-profile-map-evidence.v1",
		ID:               m.ID,
		Status:           "verified",
		GeneratedDoc:     m.GeneratedDoc,
		Workflow:         m.Workflow,
		EvidenceArtifact: m.EvidenceArtifact,
		OperationCount:   1,
		SchemaCount:      len(model.Schemas),
		PolicyRuleCount:  len(model.Policy.Rules),
		ScenarioCount:    model.ScenarioCount,
		DSLIRMatch:       model.DSLIRMatch,
		OpenAPIMatch:     model.OpenAPIMatch,
		MapShapePass:     model.MapShapePass,
		Loop:             m.Loop,
	}, "", "  ")
	if err != nil {
		return err
	}
	return writeFile(path, append(body, '\n'))
}
