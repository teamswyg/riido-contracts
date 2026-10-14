package main

type evidence struct {
	SchemaVersion            string       `json:"schema_version"`
	ID                       string       `json:"id"`
	Status                   string       `json:"status"`
	GeneratedDoc             string       `json:"generated_doc"`
	Package                  string       `json:"package"`
	Workflow                 string       `json:"workflow"`
	EvidenceArtifact         string       `json:"evidence_artifact"`
	DistributionChannelCount int          `json:"distribution_channel_count"`
	StoreManagedChannelCount int          `json:"store_managed_channel_count"`
	ProviderStatusCount      int          `json:"provider_routing_status_count"`
	NonOwnedSurfaceCount     int          `json:"non_owned_surface_count"`
	DistributionValid        bool         `json:"distribution_valid"`
	ProviderRoutingValid     bool         `json:"provider_routing_valid"`
	StoreManagedExclusive    bool         `json:"store_managed_exclusive"`
	Loop                     evidenceLoop `json:"loop"`
}
