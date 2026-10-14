package main

type manifest struct {
	SchemaVersion                  string       `json:"schema_version"`
	ID                             string       `json:"id"`
	Title                          string       `json:"title"`
	RiidoTask                      string       `json:"riido_task"`
	Summary                        string       `json:"summary"`
	GeneratedDoc                   string       `json:"generated_doc"`
	Workflow                       string       `json:"workflow"`
	EvidenceArtifact               string       `json:"evidence_artifact"`
	Package                        string       `json:"package"`
	ExpectedCanonicalEventFields   int          `json:"expected_canonical_event_fields"`
	ExpectedEventScopeCount        int          `json:"expected_event_scope_count"`
	ExpectedCommonRequiredCount    int          `json:"expected_common_static_required_count"`
	ExpectedActorIDConditional     int          `json:"expected_actor_id_conditional_count"`
	ExpectedRunRequiredFieldCount  int          `json:"expected_run_required_field_count"`
	ExpectedFakePlaceholderFields  int          `json:"expected_fake_placeholder_field_count"`
	ExpectedFakePlaceholderValues  int          `json:"expected_fake_placeholder_value_count"`
	ExpectedViolationCodeCount     int          `json:"expected_violation_code_count"`
	ExpectedNativeConfigClassCount int          `json:"expected_native_config_class_count"`
	ExpectedRunContextFields       int          `json:"expected_run_context_fields"`
	ExpectedValidateEntrypoints    int          `json:"expected_validate_entrypoints"`
	ScopeRules                     []scopeRule  `json:"scope_rules"`
	Invariants                     []string     `json:"invariants"`
	Loop                           evidenceLoop `json:"loop"`
}
