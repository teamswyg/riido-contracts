package main

import "encoding/json"

func writeEvidence(path string, model model) error {
	m := model.Manifest
	body, err := json.MarshalIndent(evidence{
		SchemaVersion:          "riido-ai-agent-task-assignment-evidence.v1",
		ID:                     m.ID,
		Status:                 "verified",
		GeneratedDoc:           m.GeneratedDoc,
		Workflow:               m.Workflow,
		EvidenceArtifact:       m.EvidenceArtifact,
		OperationCount:         len(model.Operations),
		SchemaCount:            len(model.Schemas),
		PolicyCount:            len(model.Policies),
		ScenarioCount:          model.ScenarioCount,
		DSLIRMatch:             model.DSLIRMatch,
		OpenAPIMatch:           model.OpenAPIMatch,
		ForbiddenFieldsAbsent:  model.ForbiddenFieldsAbsent,
		NoDiffPathFieldsAbsent: model.NoDiffPathsAbsent,
		Loop:                   m.Loop,
	}, "", "  ")
	if err != nil {
		return err
	}
	return writeFile(path, append(body, '\n'))
}
