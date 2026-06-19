package main

type page struct {
	NodeID     string `json:"node_id"`
	Name       string `json:"name"`
	ChildCount int    `json:"child_count"`
}

type node struct {
	NodeID string `json:"node_id"`
	Name   string `json:"name"`
}

type pageInventory struct {
	PageID string `json:"page_id"`
	Nodes  []node `json:"nodes"`
}

type coverageEntry struct {
	PageID                   string        `json:"page_id,omitempty"`
	NodeID                   string        `json:"node_id"`
	Name                     string        `json:"name"`
	CoverageStatus           string        `json:"coverage_status"`
	EvidenceKind             string        `json:"evidence_kind"`
	SSOTDocs                 []string      `json:"ssot_docs,omitempty"`
	OwnerRepos               []string      `json:"owner_repos,omitempty"`
	GeneratedPaths           []string      `json:"generated_paths,omitempty"`
	CoveredFacts             []string      `json:"covered_facts,omitempty"`
	Reason                   string        `json:"reason,omitempty"`
	AbsorbedByTopLevelNodeID string        `json:"absorbed_by_top_level_node_id,omitempty"`
	DirectionLoop            directionLoop `json:"direction_loop,omitempty"`
}

type directionLoop struct {
	TopDown  string `json:"top_down,omitempty"`
	BottomUp string `json:"bottom_up,omitempty"`
}
