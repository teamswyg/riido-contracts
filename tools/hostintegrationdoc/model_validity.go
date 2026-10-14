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

func allProviderHealthVocabularyValid() bool {
	for _, status := range hostintegration.ProviderHealthStatuses() {
		if !status.Valid() {
			return false
		}
	}
	for _, code := range hostintegration.ProviderDiagnosticCodes() {
		if !code.Valid() {
			return false
		}
	}
	return !hostintegration.ProviderHealthStatus("other").Valid() &&
		!hostintegration.ProviderDiagnosticCode("raw-error").Valid()
}
