package main

type manifest struct {
	SchemaVersion                   string                  `json:"schema_version"`
	ID                              string                  `json:"id"`
	RiidoTask                       string                  `json:"riido_task"`
	Workflow                        string                  `json:"workflow"`
	EvidenceArtifact                string                  `json:"evidence_artifact"`
	Loop                            evidenceLoop            `json:"loop"`
	StabilizedBy                    []string                `json:"stabilized_by"`
	HumanDoc                        string                  `json:"human_doc"`
	RelatedManifests                []string                `json:"related_manifests"`
	Figma                           figmaSource             `json:"figma"`
	InspectionMethod                inspectionMethod        `json:"inspection_method"`
	SupportingToolLimitations       []toolLimitation        `json:"supporting_tool_limitations"`
	CoveragePolicy                  coveragePolicy          `json:"coverage_policy"`
	APIAnnotationContentPolicy      annotationContentPolicy `json:"api_generated_annotation_content_policy"`
	ToolLimitationFiles             []string                `json:"tool_limitation_files"`
	ExpectedTopLevelNodeFiles       []string                `json:"expected_top_level_node_files"`
	PageInventoryFiles              []string                `json:"page_inventory_files"`
	CoverageEntryFiles              []string                `json:"coverage_entry_files"`
	NonUICoverageEntryFiles         []string                `json:"non_ui_coverage_entry_files"`
	APIAnnotationInventoryFiles     []string                `json:"api_generated_annotation_inventory_files"`
	APIAnnotationFiles              []string                `json:"api_generated_annotation_files"`
	VerifiedEvidenceNodeFiles       []string                `json:"verified_evidence_node_files"`
	ExpectedPages                   []page                  `json:"expected_pages"`
	ExpectedTopLevelNodes           []node                  `json:"expected_top_level_nodes"`
	NonUITopLevelInventory          []pageInventory         `json:"non_ui_top_level_inventory"`
	NonUITopLevelNodes              []coverageEntry         `json:"non_ui_top_level_nodes"`
	Entries                         []coverageEntry         `json:"entries"`
	APIGeneratedAnnotationInventory []annotationInventory   `json:"api_generated_annotation_inventory"`
	APIGeneratedAnnotations         []annotation            `json:"api_generated_annotations"`
	VerifiedEvidenceNodes           []node                  `json:"verified_evidence_nodes"`
}

type coverageEntryDocument struct {
	SchemaVersion string        `json:"schema_version"`
	Entry         coverageEntry `json:"entry"`
}

type pageInventoryDocument struct {
	SchemaVersion string        `json:"schema_version"`
	Inventory     pageInventory `json:"inventory"`
}

type toolLimitationDocument struct {
	SchemaVersion string         `json:"schema_version"`
	Limitation    toolLimitation `json:"limitation"`
}

type annotationInventoryDocument struct {
	SchemaVersion string              `json:"schema_version"`
	Inventory     annotationInventory `json:"inventory"`
}

type annotationDocument struct {
	SchemaVersion string     `json:"schema_version"`
	Annotation    annotation `json:"annotation"`
}

type figmaSource struct {
	FileKey          string `json:"file_key"`
	FileName         string `json:"file_name"`
	PageID           string `json:"page_id"`
	PageName         string `json:"page_name"`
	InspectedAt      string `json:"inspected_at"`
	InspectionSource string `json:"inspection_source"`
}
