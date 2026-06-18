package apicontract

type figmaCoverageManifest struct {
	SchemaVersion                       string                                 `json:"schema_version"`
	ID                                  string                                 `json:"id"`
	RiidoTask                           string                                 `json:"riido_task"`
	StabilizedBy                        []string                               `json:"stabilized_by"`
	HumanDoc                            string                                 `json:"human_doc"`
	RelatedManifests                    []string                               `json:"related_manifests"`
	Figma                               figmaCoverageSource                    `json:"figma"`
	InspectionMethod                    figmaCoverageInspectionMethod          `json:"inspection_method"`
	SupportingToolLimitations           []figmaSupportingToolLimitation        `json:"supporting_tool_limitations"`
	CoveragePolicy                      figmaCoveragePolicy                    `json:"coverage_policy"`
	APIGeneratedAnnotationContentPolicy figmaAPIGeneratedAnnotationContentRule `json:"api_generated_annotation_content_policy"`
	ExpectedPages                       []figmaCoveragePage                    `json:"expected_pages"`
	ExpectedTopLevelNodes               []figmaCoverageNode                    `json:"expected_top_level_nodes"`
	NonUITopLevelInventory              []figmaNonUITopLevelInventory          `json:"non_ui_top_level_inventory"`
	VerifiedEvidenceNodes               []figmaCoverageNode                    `json:"verified_evidence_nodes"`
	NonUITopLevelNodes                  []figmaCoverageEntry                   `json:"non_ui_top_level_nodes"`
	APIGeneratedAnnotations             []figmaAPIGeneratedAnnotation          `json:"api_generated_annotations"`
	APIGeneratedAnnotationInventory     []figmaAPIGeneratedAnnotationGroup     `json:"api_generated_annotation_inventory"`
	Entries                             []figmaCoverageEntry                   `json:"entries"`
}
