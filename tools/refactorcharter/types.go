package main

type charter struct {
	SchemaVersion     string     `json:"schema_version"`
	ID                string     `json:"id"`
	RiidoTask         string     `json:"riido_task"`
	Mode              string     `json:"mode"`
	LineBudget        lineBudget `json:"line_budget"`
	SemanticUnits     []string   `json:"semantic_units"`
	RequiredArtifacts []string   `json:"required_artifacts"`
	Scan              scanConfig `json:"scan"`
}

type lineBudget struct {
	TargetMaxLines      int `json:"target_max_lines"`
	RecommendedMinLines int `json:"recommended_min_lines"`
	RecommendedMaxLines int `json:"recommended_max_lines"`
}

type scanConfig struct {
	Roots                  []string `json:"roots"`
	IncludeExtensions      []string `json:"include_extensions"`
	GeneratedPathFragments []string `json:"generated_path_fragments"`
	GeneratedMarkers       []string `json:"generated_markers"`
}

type finding struct {
	Path  string `json:"path"`
	Lines int    `json:"lines"`
}

type scanReport struct {
	FilesScanned int
	Findings     []finding
}
