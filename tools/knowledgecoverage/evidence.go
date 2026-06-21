package main

import "encoding/json"

type evidence struct {
	SchemaVersion              string                `json:"schema_version"`
	ID                         string                `json:"id"`
	Status                     string                `json:"status"`
	ScannedCount               int                   `json:"scanned_count"`
	GeneratedCount             int                   `json:"generated_count"`
	ExecutableCount            int                   `json:"executable_count"`
	GeneratedAdjacentCount     int                   `json:"generated_adjacent_manifest_count"`
	ExecutableAdjacentCount    int                   `json:"executable_adjacent_manifest_count"`
	AdjacentCount              int                   `json:"adjacent_manifest_count"`
	ManualCount                int                   `json:"manual_count"`
	ManualSamples              []docRecord           `json:"manual_samples"`
	ManifestInventory          int                   `json:"manifest_inventory_count"`
	ManifestInventoryByGroup   []manifestGroupCount  `json:"manifest_inventory_by_group"`
	ManifestInventorySamples   []manifestGroupSample `json:"manifest_inventory_samples"`
	ManifestLoopCount          int                   `json:"manifest_loop_count"`
	ManifestMissingLoopCount   int                   `json:"manifest_missing_loop_count"`
	ManifestMissingLoopByGroup []manifestGroupCount  `json:"manifest_missing_loop_by_group"`
	ManifestMissingLoopSamples []manifestGroupSample `json:"manifest_missing_loop_samples"`
	EvidenceArtifact           string                `json:"evidence_artifact"`
	Workflow                   string                `json:"workflow"`
	Loop                       evidenceLoop          `json:"loop"`
}

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
		ManifestMissingLoopCount:   report.ManifestLoops.Missing,
		ManifestMissingLoopByGroup: report.ManifestLoops.MissingGroups,
		ManifestMissingLoopSamples: report.ManifestLoops.MissingSamples,
		EvidenceArtifact:           m.EvidenceArtifact,
		Workflow:                   m.Workflow, Loop: m.Loop,
	}
}

func manualSamples(report scanReport) []docRecord {
	if report.ManualSamples == nil {
		return []docRecord{}
	}
	return report.ManualSamples
}

func status(report scanReport) string {
	if report.ManualCount > 0 {
		return "failed"
	}
	return "verified"
}
