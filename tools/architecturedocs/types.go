package main

type manifest struct {
	SchemaVersion     string           `json:"schema_version"`
	ID                string           `json:"id"`
	RiidoTask         string           `json:"riido_task"`
	Summary           string           `json:"summary"`
	GeneratedDocs     generatedDocs    `json:"generated_docs"`
	RequiredDocs      []string         `json:"required_docs"`
	ContextMapDoc     string           `json:"context_map_doc"`
	PackageRules      packageRules     `json:"package_rules"`
	Packages          []packageEntry   `json:"packages"`
	ContractShape     contractShape    `json:"contract_shape"`
	PublicGates       []publicGate     `json:"public_gates"`
	DownstreamGates   []downstreamGate `json:"downstream_gates"`
	LocalCommands     []string         `json:"local_commands"`
	StaleScanPaths    []string         `json:"stale_scan_paths"`
	StaleRuntimeWords []string         `json:"stale_runtime_words"`
	Loop              evidenceLoop     `json:"loop"`
}

type generatedDocs struct {
	ModuleDecomposition string `json:"module_decomposition"`
	IntegrationMatrix   string `json:"integration_matrix"`
}

type packageRules struct {
	Dependency  string `json:"dependency"`
	Forbidden   string `json:"forbidden"`
	TagBoundary string `json:"tag_boundary"`
}

type packageEntry struct {
	Path       string `json:"path"`
	Display    string `json:"display"`
	Role       string `json:"role"`
	MustNotOwn string `json:"must_not_own"`
}

type contractShape struct {
	Allowed   []string `json:"allowed"`
	Forbidden []string `json:"forbidden"`
}

type publicGate struct {
	Surface            string `json:"surface"`
	Verification       string `json:"verification"`
	ExternalDependency string `json:"external_dependency"`
}

type downstreamGate struct {
	Consumer     string `json:"consumer"`
	ExpectedGate string `json:"expected_gate"`
}

type evidenceLoop struct {
	Observation   string `json:"observation"`
	Hypothesis    string `json:"hypothesis"`
	Execute       string `json:"execute"`
	Evaluate      string `json:"evaluate"`
	Retrospective string `json:"retrospective"`
}
