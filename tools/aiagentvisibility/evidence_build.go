package main

import "encoding/json"

func writeEvidence(path string, model model) error {
	m := model.Manifest
	body, err := json.MarshalIndent(evidence{
		SchemaVersion:    "riido-ai-agent-visibility-evidence.v1",
		ID:               m.ID,
		Status:           "verified",
		GeneratedDoc:     m.GeneratedDoc,
		Workflow:         m.Workflow,
		EvidenceArtifact: m.EvidenceArtifact,
		OperationCount:   len(model.Operations),
		SchemaCount:      len(model.Schemas),
		PolicyRuleCount:  len(model.Policy.Rules),
		EnumValueCount:   len(model.Enum.Values),
		ScenarioCount:    model.ScenarioCount,
		DSLIRMatch:       model.DSLIRMatch,
		OpenAPIMatch:     model.OpenAPIMatch,
		Loop:             m.Loop,
	}, "", "  ")
	if err != nil {
		return err
	}
	return writeFile(path, append(body, '\n'))
}
