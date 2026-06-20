package main

type manifest struct {
	SchemaVersion                 string       `json:"schema_version"`
	ID                            string       `json:"id"`
	Title                         string       `json:"title"`
	RiidoTask                     string       `json:"riido_task"`
	Summary                       string       `json:"summary"`
	GeneratedDoc                  string       `json:"generated_doc"`
	Workflow                      string       `json:"workflow"`
	EvidenceArtifact              string       `json:"evidence_artifact"`
	Package                       string       `json:"package"`
	ExpectedDistributionChannels  int          `json:"expected_distribution_channel_count"`
	ExpectedStoreManagedChannels  int          `json:"expected_store_managed_channel_count"`
	ExpectedProviderRoutingStatus int          `json:"expected_provider_routing_status_count"`
	ExpectedNonOwnedSurfaces      int          `json:"expected_non_owned_surface_count"`
	Invariants                    []string     `json:"invariants"`
	Loop                          evidenceLoop `json:"loop"`
}
