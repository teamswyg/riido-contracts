package main

import "encoding/json"

func writeEvidence(path string, m manifest, report scanReport) error {
	body, err := json.MarshalIndent(newEvidence(m, report), "", "  ")
	if err != nil {
		return err
	}
	return writeFile(path, append(body, '\n'))
}

func newEvidence(m manifest, report scanReport) evidence {
	return evidence{
		SchemaVersion: evidenceVersion, ID: m.ID, Status: status(report),
		ScannedCount: report.ScannedCount, GeneratedCount: report.GeneratedCount,
		ExecutableCount: report.ExecutableCount, GeneratedAdjacentCount: report.GeneratedAdjacentCount,
		ExecutableAdjacentCount: report.ExecutableAdjacentCount, AdjacentCount: report.AdjacentCount,
		ManualCount: report.ManualCount, ManualSamples: manualSamples(report),
		ManifestInventory: report.ManifestInventory, ManifestInventoryByGroup: report.ManifestInventoryByGroup,
		ManifestInventorySamples:   report.ManifestInventorySamples,
		ManifestLoopCount:          report.ManifestLoops.Complete,
		ManifestDirectLoopCount:    report.ManifestLoops.Direct,
		ManifestDelegatedLoopCount: report.ManifestLoops.Delegated,
		ManifestMissingLoopCount:   report.ManifestLoops.Missing,
		ManifestMissingLoopByGroup: report.ManifestLoops.MissingGroups,
		ManifestMissingLoopSamples: report.ManifestLoops.MissingSamples,
		ManifestLoopBudget:         m.ManifestLoopBudget,
		ProblemSummaries:           report.Problems,
		EvidenceArtifact:           m.EvidenceArtifact,
		Workflow:                   m.Workflow,
		WorkflowTriggerPathCount:   len(m.WorkflowTriggerPaths),
		WorkflowTriggerPaths:       m.WorkflowTriggerPaths,
		Loop:                       m.Loop,
	}
}

func manualSamples(report scanReport) []docRecord {
	if report.ManualSamples == nil {
		return []docRecord{}
	}
	return report.ManualSamples
}

func status(report scanReport) string {
	if report.ManualCount > 0 || len(report.Problems) > 0 {
		return "failed"
	}
	return "verified"
}
