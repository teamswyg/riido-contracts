package main

func newEvidence(m manifest, summaries []fixtureSummary) evidence {
	return evidence{
		SchemaVersion:          evidenceSchema,
		ID:                     m.ID,
		Status:                 "verified",
		FixtureCount:           len(summaries),
		OperationCount:         totalOperations(summaries),
		GeneratedPathCount:     totalGeneratedPaths(summaries),
		RequiredGeneratedPaths: append([]string(nil), m.RequiredGeneratedPaths...),
		EvidenceArtifact:       m.EvidenceArtifact,
		Workflow:               m.Workflow,
		Fixtures:               summaries,
		Loop:                   m.Loop,
	}
}

func totalGeneratedPaths(summaries []fixtureSummary) int {
	total := 0
	for _, summary := range summaries {
		total += summary.GeneratedPathCount
	}
	return total
}
