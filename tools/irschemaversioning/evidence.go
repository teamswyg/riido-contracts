package main

type evidence struct {
	SchemaVersion          string       `json:"schema_version"`
	ID                     string       `json:"id"`
	Status                 string       `json:"status"`
	GeneratedDoc           string       `json:"generated_doc"`
	Package                string       `json:"package"`
	Workflow               string       `json:"workflow"`
	EvidenceArtifact       string       `json:"evidence_artifact"`
	CanonicalEventFields   int          `json:"canonical_event_fields"`
	EventScopeCount        int          `json:"event_scope_count"`
	CommonRequiredCount    int          `json:"common_static_required_count"`
	RunRequiredFieldCount  int          `json:"run_required_field_count"`
	FakePlaceholderFields  int          `json:"fake_placeholder_fields"`
	FakePlaceholderValues  int          `json:"fake_placeholder_values"`
	ViolationCodeCount     int          `json:"violation_code_count"`
	NativeConfigClassCount int          `json:"native_config_class_count"`
	RunContextFields       int          `json:"run_context_fields"`
	ValidateEntrypoints    int          `json:"validate_entrypoints"`
	ScopeRules             []scopeRule  `json:"scope_rules"`
	Loop                   evidenceLoop `json:"loop"`
}
