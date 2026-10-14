package main

type manifest struct {
	SchemaVersion        string               `json:"schema_version"`
	ID                   string               `json:"id"`
	RiidoTask            string               `json:"riido_task"`
	GeneratedDoc         string               `json:"generated_doc"`
	Goal                 string               `json:"goal"`
	PromotionRule        promotionRule        `json:"promotion_rule"`
	CandidateContracts   []candidateContract  `json:"candidate_contracts"`
	RepositoryBoundaries repositoryBoundaries `json:"repository_boundaries"`
	Versioning           versioning           `json:"versioning"`
	MigrationOrder       []string             `json:"migration_order"`
	MigrationSlices      []migrationSlice     `json:"migration_slices"`
	ValidationGates      validationGates      `json:"validation_gates"`
	MigrationWorkMap     []workMapEntry       `json:"migration_work_map"`
	Workflow             string               `json:"workflow"`
	EvidenceArtifact     string               `json:"evidence_artifact"`
	Loop                 evidenceLoop         `json:"loop"`
}

type promotionRule struct {
	Conditions []string `json:"conditions"`
	Fallback   string   `json:"fallback"`
}

type repositoryBoundaries struct {
	MayContain     []string `json:"may_contain"`
	MustNotContain []string `json:"must_not_contain"`
}

type validationGates struct {
	RequiredCommands     []string `json:"required_commands"`
	ArchitectureCommands []string `json:"architecture_commands"`
	FixtureChecks        []string `json:"fixture_checks"`
}
