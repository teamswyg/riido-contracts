package main

import "github.com/teamswyg/riido-contracts/hostintegration"

func buildModel(m manifest) model {
	return model{
		Manifest:              m,
		DistributionChannels:  distributionStrings(),
		StoreManagedChannels:  storeManagedStrings(),
		ProviderStatuses:      providerStatusStrings(),
		NonOwnedSurfaces:      hostintegration.NonOwnedSurfaces(),
		DistributionValid:     allDistributionChannelsValid(),
		ProviderRoutingValid:  allProviderStatusesValid(),
		StoreManagedExclusive: storeManagedExclusive(),
	}
}
