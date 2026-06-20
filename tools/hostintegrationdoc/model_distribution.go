package main

import "github.com/teamswyg/riido-contracts/hostintegration"

func distributionStrings() []string {
	out := make([]string, 0, len(hostintegration.DistributionChannels()))
	for _, channel := range hostintegration.DistributionChannels() {
		out = append(out, string(channel))
	}
	return out
}

func storeManagedStrings() []string {
	out := make([]string, 0, len(hostintegration.StoreManagedDistributionChannels()))
	for _, channel := range hostintegration.StoreManagedDistributionChannels() {
		out = append(out, string(channel))
	}
	return out
}
