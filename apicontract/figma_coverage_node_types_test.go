package apicontract

type figmaCoverageNode struct {
	NodeID string `json:"node_id"`
	Name   string `json:"name"`
}

type figmaCoveragePage struct {
	NodeID     string `json:"node_id"`
	Name       string `json:"name"`
	ChildCount int    `json:"child_count"`
}

type figmaNonUITopLevelInventory struct {
	PageID string              `json:"page_id"`
	Nodes  []figmaCoverageNode `json:"nodes"`
}

type figmaCoverageEntry struct {
	NodeID                   string                 `json:"node_id"`
	PageID                   string                 `json:"page_id,omitempty"`
	Name                     string                 `json:"name"`
	CoverageStatus           string                 `json:"coverage_status"`
	EvidenceKind             string                 `json:"evidence_kind"`
	AbsorbedByTopLevelNodeID string                 `json:"absorbed_by_top_level_node_id,omitempty"`
	SSOTDocs                 []string               `json:"ssot_docs,omitempty"`
	OwnerRepos               []string               `json:"owner_repos,omitempty"`
	GeneratedPaths           []string               `json:"generated_paths,omitempty"`
	CoveredFacts             []string               `json:"covered_facts,omitempty"`
	DirectionLoop            figmaCoverageDirection `json:"direction_loop"`
	Reason                   string                 `json:"reason,omitempty"`
}

type figmaCoverageDirection struct {
	TopDown  string `json:"top_down,omitempty"`
	BottomUp string `json:"bottom_up,omitempty"`
}
