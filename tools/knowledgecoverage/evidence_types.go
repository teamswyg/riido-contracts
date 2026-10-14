package main

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
	ManifestDirectLoopCount    int                   `json:"manifest_direct_loop_count"`
	ManifestDelegatedLoopCount int                   `json:"manifest_delegated_loop_count"`
	ManifestMissingLoopCount   int                   `json:"manifest_missing_loop_count"`
	ManifestMissingLoopByGroup []manifestGroupCount  `json:"manifest_missing_loop_by_group"`
	ManifestMissingLoopSamples []manifestGroupSample `json:"manifest_missing_loop_samples"`
	ManifestLoopBudget         manifestLoopBudget    `json:"manifest_loop_budget"`
	ProblemSummaries           []string              `json:"problem_summaries"`
	EvidenceArtifact           string                `json:"evidence_artifact"`
	Workflow                   string                `json:"workflow"`
	WorkflowTriggerPathCount   int                   `json:"workflow_trigger_path_count"`
	WorkflowTriggerPaths       []string              `json:"workflow_trigger_paths"`
	Loop                       evidenceLoop          `json:"loop"`
}
