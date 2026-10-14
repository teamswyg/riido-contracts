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
	ExpectedProviderCapabilityFields int          `json:"expected_provider_capability_fields"`
	ExpectedFingerprintInputFields   int          `json:"expected_fingerprint_input_fields"`
	ExpectedProtocolCount            int          `json:"expected_protocol_count"`
	ExpectedEventStreamFormatCount   int          `json:"expected_event_stream_format_count"`
	ExpectedProtocolMaturityCount    int          `json:"expected_protocol_maturity_count"`
	ExpectedCompatibilityStatusCount int          `json:"expected_compatibility_status_count"`
	ExpectedProtocolCriticalArgSets  int          `json:"expected_protocol_critical_arg_sets"`
	Invariants                       []string     `json:"invariants"`
	Loop                             evidenceLoop `json:"loop"`
}
