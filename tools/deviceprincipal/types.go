package main

type manifest struct {
	SchemaVersion                    string       `json:"schema_version"`
	ID                               string       `json:"id"`
	Title                            string       `json:"title"`
	RiidoTask                        string       `json:"riido_task"`
	Summary                          string       `json:"summary"`
	GeneratedDoc                     string       `json:"generated_doc"`
	Workflow                         string       `json:"workflow"`
	EvidenceArtifact                 string       `json:"evidence_artifact"`
	Package                          string       `json:"package"`
	APIPolicyID                      string       `json:"api_policy_id"`
	DependencyFactID                 string       `json:"dependency_fact_id"`
	ExpectedPrincipalCount           int          `json:"expected_principal_count"`
	ExpectedDaemonHeaderCount        int          `json:"expected_daemon_header_count"`
	ExpectedClientHeaderCount        int          `json:"expected_client_header_count"`
	ExpectedSnapshotIntervalSeconds  int          `json:"expected_runtime_snapshot_interval_seconds"`
	ExpectedRuntimeStaleAfterSeconds int          `json:"expected_runtime_stale_after_seconds"`
	ExpectedOwnershipEdgeCount       int          `json:"expected_ownership_edge_count"`
	ExpectedBindingSourceCount       int          `json:"expected_binding_source_count"`
	ExpectedExcludedFallbackCount    int          `json:"expected_excluded_fallback_count"`
	ExpectedSecretNonExposureSinks   int          `json:"expected_secret_non_exposure_sink_count"`
	ExpectedDependencyPhraseCount    int          `json:"expected_dependency_phrase_count"`
	ExpectedBindingFieldCount        int          `json:"expected_binding_field_count"`
	PolicyRulePrefixes               []string     `json:"policy_rule_prefixes"`
	Invariants                       []string     `json:"invariants"`
	Loop                             evidenceLoop `json:"loop"`
}
