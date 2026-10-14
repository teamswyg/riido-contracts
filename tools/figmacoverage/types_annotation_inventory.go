package main

type annotationInventory struct {
	UIArea                 string             `json:"ui_area"`
	CategoryID             string             `json:"category_id"`
	CategoryLabel          string             `json:"category_label"`
	FigmaGeneratedPath     string             `json:"figma_generated_path"`
	CanonicalGeneratedPath string             `json:"canonical_generated_path"`
	OperationKind          string             `json:"operation_kind"`
	Background             string             `json:"background"`
	AnnotationCount        int                `json:"annotation_count"`
	Sources                []annotationSource `json:"sources"`
}

type annotationSource struct {
	PageID              string   `json:"page_id"`
	TopLevelNodeID      string   `json:"top_level_node_id"`
	CoverageEntryNodeID string   `json:"coverage_entry_node_id"`
	NodeIDs             []string `json:"node_ids"`
}

type annotation struct {
	NodeID                 string `json:"node_id"`
	TopLevelNodeID         string `json:"top_level_node_id"`
	CoverageEntryNodeID    string `json:"coverage_entry_node_id"`
	CategoryID             string `json:"category_id"`
	CategoryLabel          string `json:"category_label"`
	FigmaLabel             string `json:"figma_label"`
	FigmaGeneratedPath     string `json:"figma_generated_path"`
	CanonicalGeneratedPath string `json:"canonical_generated_path"`
	ResolutionStatus       string `json:"resolution_status"`
	Resolution             string `json:"resolution"`
}
