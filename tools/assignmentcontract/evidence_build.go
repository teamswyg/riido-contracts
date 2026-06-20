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
		PayloadFields:        payloadEvidenceFrom(c.AssignmentPayloadFields),
		EvidenceArtifact:     m.EvidenceArtifact,
		Workflow:             m.Workflow,
		Loop:                 m.Loop,
	}
}

func payloadEvidenceFrom(fields []payloadField) []payloadEvidence {
	out := make([]payloadEvidence, 0, len(fields))
	for _, field := range fields {
		out = append(out, payloadEvidence(field))
	}
	return out
}
