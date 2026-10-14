package main

type model struct {
	Manifest                 manifest
	ProviderCapabilityFields int
	FingerprintInputFields   int
	Protocols                []protocolRow
	EventStreamFormats       []string
	ProtocolMaturities       []string
	CompatibilityStatuses    []string
	CriticalArgSetCount      int
}

type protocolRow struct {
	Kind string
	Args []string
}
