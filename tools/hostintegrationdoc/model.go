package main

type model struct {
	Manifest              manifest
	DistributionChannels  []string
	StoreManagedChannels  []string
	ProviderStatuses      []string
	NonOwnedSurfaces      []string
	DistributionValid     bool
	ProviderRoutingValid  bool
	StoreManagedExclusive bool
}
