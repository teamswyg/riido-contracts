package main

func newEvidence(m manifest, c contract) evidence {
	return evidence{
		SchemaVersion:        evidenceSchema,
		ID:                   m.ID,
		Status:               "verified",
		Contract:             m.Contract,
		ContractSchema:       c.SchemaVersion,
		ServiceSchema:        c.ServiceSchemaVersion,
		AssignmentStateCount: len(c.AssignmentStates),
		PollActionCount:      len(c.PollActions),
		TaskEventCount:       len(c.TaskEvents),
		PayloadFieldCount:    len(c.AssignmentPayloadFields),
		EvidenceArtifact:     m.EvidenceArtifact,
		Workflow:             m.Workflow,
		Loop:                 m.Loop,
	}
}
