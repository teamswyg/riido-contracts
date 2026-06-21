package main

type manifest struct {
	SchemaVersion        string        `json:"schema_version"`
	ID                   string        `json:"id"`
	RiidoTask            string        `json:"riido_task"`
	Workflow             string        `json:"workflow"`
	EvidenceArtifact     string        `json:"evidence_artifact"`
	Loop                 evidenceLoop  `json:"loop"`
	GeneratedDoc         string        `json:"generated_doc"`
	Summary              string        `json:"summary"`
	PromotionConditions  []string      `json:"promotion_conditions"`
	SingleRuntimeRule    string        `json:"single_runtime_rule"`
	SchemaVersionAxes    []versionAxis `json:"schema_version_axes"`
	ModuleTagRule        string        `json:"module_tag_rule"`
	RuntimeTagModel      []runtimeTag  `json:"runtime_tag_model"`
	RuntimeTagRule       string        `json:"runtime_tag_rule"`
	BreakingChangeRules  []string      `json:"breaking_change_rules"`
	AdditiveChangeRule   string        `json:"additive_change_rule"`
	DownstreamImportRule string        `json:"downstream_import_rule"`
}

type versionAxis struct {
	Axis string `json:"axis"`
	Rule string `json:"rule"`
}

type runtimeTag struct {
	Pattern string `json:"pattern"`
	Meaning string `json:"meaning"`
}
