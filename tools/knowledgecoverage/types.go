package main

type manifest struct {
	SchemaVersion    string       `json:"schema_version"`
	ID               string       `json:"id"`
	Title            string       `json:"title"`
	GeneratedDoc     string       `json:"generated_doc"`
	Workflow         string       `json:"workflow"`
	EvidenceArtifact string       `json:"evidence_artifact"`
	ScanRoots        []string     `json:"scan_roots"`
	ScanFiles        []string     `json:"scan_files"`
	GeneratedMarkers []string     `json:"generated_markers"`
	Loop             evidenceLoop `json:"loop"`
}

type evidenceLoop struct {
	Observation   string `json:"observation"`
	Hypothesis    string `json:"hypothesis"`
	Execute       string `json:"execute"`
	Evaluate      string `json:"evaluate"`
	Retrospective string `json:"retrospective"`
}

type docRecord struct {
	Path                string `json:"path"`
	Lines               int    `json:"lines"`
	Classification      string `json:"classification"`
	HasGeneratedMarker  bool   `json:"has_generated_marker"`
	HasExecutableMarker bool   `json:"has_executable_marker"`
	HasAdjacentManifest bool   `json:"has_adjacent_manifest"`
}

type scanReport struct {
	Docs                     []docRecord
	ScannedCount             int
	GeneratedCount           int
	ExecutableCount          int
	GeneratedAdjacentCount   int
	ExecutableAdjacentCount  int
	AdjacentCount            int
	ManualCount              int
	ManualSamples            []docRecord
	ManifestInventory        int
	ManifestInventoryByGroup []manifestGroupCount
}

type manifestGroupCount struct {
	Group string `json:"group"`
	Count int    `json:"count"`
}
