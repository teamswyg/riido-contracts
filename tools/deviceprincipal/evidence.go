package main

type evidence struct {
	SchemaVersion         string       `json:"schema_version"`
	ID                    string       `json:"id"`
	Status                string       `json:"status"`
	GeneratedDoc          string       `json:"generated_doc"`
	Package               string       `json:"package"`
	Workflow              string       `json:"workflow"`
	EvidenceArtifact      string       `json:"evidence_artifact"`
	PrincipalCount        int          `json:"principal_count"`
	DaemonHeaderCount     int          `json:"daemon_header_count"`
	ClientHeaderCount     int          `json:"client_header_count"`
	SnapshotInterval      int          `json:"runtime_snapshot_interval_seconds"`
	RuntimeStaleAfter     int          `json:"runtime_stale_after_seconds"`
	OwnershipEdgeCount    int          `json:"ownership_edge_count"`
	BindingSourceCount    int          `json:"binding_source_count"`
	ExcludedFallbackCount int          `json:"excluded_fallback_count"`
	SecretSinkCount       int          `json:"secret_non_exposure_sink_count"`
	DependencyPhraseCount int          `json:"dependency_phrase_count"`
	BindingFieldCount     int          `json:"binding_field_count"`
	PolicyRuleCount       int          `json:"policy_rule_count"`
	Loop                  evidenceLoop `json:"loop"`
}
