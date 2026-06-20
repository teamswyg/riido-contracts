package main

func newEvidence(model model) evidence {
	m := model.Manifest
	return evidence{
		SchemaVersion: evidenceSchema,
		ID:            m.ID,
		Status:        "verified",
		GeneratedDoc:  m.GeneratedDoc,
		SourceDSL:     m.SourceDSL,
		Workflow:      m.Workflow,
		Artifact:      m.EvidenceArtifact,
		FSMSchema:     model.FSMSchema,
		States:        len(model.States),
		StartStates:   len(model.StartStates),
		Terminals:     len(model.TerminalStates),
		Transitions:   model.TransitionCount,
		Invariants:    len(m.Invariants),
		Loop:          m.Loop,
	}
}
