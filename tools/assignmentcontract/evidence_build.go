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
		AssignmentStateFiles: len(c.AssignmentStateFiles),
		PollActionCount:      len(c.PollActions),
		PollActionFiles:      len(c.PollActionFiles),
		TaskEventCount:       len(c.TaskEvents),
		TaskEventFiles:       len(c.TaskEventFiles),
		PayloadFieldCount:    len(c.AssignmentPayloadFields),
		PayloadFieldFiles:    len(c.PayloadFieldFiles),
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
