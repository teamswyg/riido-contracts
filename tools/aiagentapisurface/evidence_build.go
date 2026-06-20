package main

import "encoding/json"

func writeEvidence(path string, model model) error {
	m := model.Manifest
	body, err := json.MarshalIndent(evidence{
		SchemaVersion:     "riido-ai-agent-api-surface-evidence.v1",
		ID:                m.ID,
		Status:            "verified",
		GeneratedDoc:      m.GeneratedDoc,
		Workflow:          m.Workflow,
		EvidenceArtifact:  m.EvidenceArtifact,
		OperationCount:    len(model.Operations),
		V1OperationCount:  model.V1Count,
		V2OperationCount:  model.V2Count,
		V2OnlyOperations:  model.V2Only,
		OpenAPIPathCount:  model.OpenAPIPathCount,
		OpenAPIOpCount:    model.OpenAPIOpCount,
		StreamVariants:    model.StreamVariants,
		DSLIRMatch:        model.DSLIRMatch,
		IROpenAPIMatch:    model.IROpenAPIMatch,
		V2CoversV1:        model.V2CoversV1,
		StreamVariantPass: model.StreamVariantPass,
		Loop:              m.Loop,
	}, "", "  ")
	if err != nil {
		return err
	}
	return writeFile(path, append(body, '\n'))
}
