package main

import "encoding/json"

type evidence struct {
	SchemaVersion            string       `json:"schema_version"`
	ID                       string       `json:"id"`
	Status                   string       `json:"status"`
	GeneratedDoc             string       `json:"generated_doc"`
	Package                  string       `json:"package"`
	Workflow                 string       `json:"workflow"`
	Artifact                 string       `json:"evidence_artifact"`
	ProviderCapabilityFields int          `json:"provider_capability_fields"`
	FingerprintInputFields   int          `json:"fingerprint_input_fields"`
	ProtocolCount            int          `json:"protocol_count"`
	ProtocolCriticalArgSets  int          `json:"protocol_critical_arg_sets"`
	EventStreamFormatCount   int          `json:"event_stream_format_count"`
	ProtocolMaturityCount    int          `json:"protocol_maturity_count"`
	CompatibilityStatusCount int          `json:"compatibility_status_count"`
	InvariantCount           int          `json:"invariant_count"`
	Loop                     evidenceLoop `json:"loop"`
}

func writeEvidence(path string, model model) error {
	body, err := json.MarshalIndent(newEvidence(model), "", "  ")
	if err != nil {
		return err
	}
	return writeFile(path, append(body, '\n'))
}
