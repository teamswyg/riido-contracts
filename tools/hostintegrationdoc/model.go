package main

type model struct {
	Manifest                manifest
	DistributionChannels    []string
	StoreManagedChannels    []string
	ProviderStatuses        []string
	ProviderHealthStatuses  []string
	ProviderDiagnosticCodes []string
	NonOwnedSurfaces        []string
	DistributionValid       bool
	ProviderRoutingValid    bool
	ProviderHealthValid     bool
	StoreManagedExclusive   bool
}
