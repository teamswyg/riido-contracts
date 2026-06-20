package main

import "encoding/json"

type evidence struct {
	SchemaVersion     string       `json:"schema_version"`
	ID                string       `json:"id"`
	Status            string       `json:"status"`
	ScannedCount      int          `json:"scanned_count"`
	GeneratedCount    int          `json:"generated_count"`
	ExecutableCount   int          `json:"executable_count"`
	AdjacentCount     int          `json:"adjacent_manifest_count"`
	ManualCount       int          `json:"manual_count"`
	ManualSamples     []docRecord  `json:"manual_samples"`
	ManifestInventory int          `json:"manifest_inventory_count"`
	EvidenceArtifact  string       `json:"evidence_artifact"`
	Workflow          string       `json:"workflow"`
	Loop              evidenceLoop `json:"loop"`
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
		ExecutableCount: report.ExecutableCount, AdjacentCount: report.AdjacentCount,
		ManualCount: report.ManualCount, ManualSamples: report.ManualSamples,
		ManifestInventory: report.ManifestInventory, EvidenceArtifact: m.EvidenceArtifact,
		Workflow: m.Workflow, Loop: m.Loop,
	}
}

func status(report scanReport) string {
	if report.ManualCount > 0 {
		return "advisory_findings"
	}
	return "verified"
}
