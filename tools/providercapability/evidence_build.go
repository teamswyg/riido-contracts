package main

func newEvidence(model model) evidence {
	m := model.Manifest
	return evidence{
		SchemaVersion:            evidenceSchema,
		ID:                       m.ID,
		Status:                   "verified",
		GeneratedDoc:             m.GeneratedDoc,
		Package:                  m.Package,
		Workflow:                 m.Workflow,
		Artifact:                 m.EvidenceArtifact,
		ProviderCapabilityFields: model.ProviderCapabilityFields,
		FingerprintInputFields:   model.FingerprintInputFields,
		ProtocolCount:            len(model.Protocols),
		ProtocolCriticalArgSets:  model.CriticalArgSetCount,
		EventStreamFormatCount:   len(model.EventStreamFormats),
		ProtocolMaturityCount:    len(model.ProtocolMaturities),
		CompatibilityStatusCount: len(model.CompatibilityStatuses),
		InvariantCount:           len(m.Invariants),
		Loop:                     m.Loop,
	}
}
