package main

import "github.com/teamswyg/riido-contracts/hostintegration"

func allDistributionChannelsValid() bool {
	for _, channel := range hostintegration.DistributionChannels() {
		if !channel.Valid() {
			return false
		}
	}
	return !hostintegration.DistributionChannel("unknown").Valid()
}

func allProviderStatusesValid() bool {
	for _, status := range hostintegration.ProviderRoutingStatuses() {
		if !status.Valid() {
			return false
		}
	}
	return !hostintegration.ProviderRoutingStatus("unknown").Valid()
}
