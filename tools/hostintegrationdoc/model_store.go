package main

import "github.com/teamswyg/riido-contracts/hostintegration"

func storeManagedExclusive() bool {
	allowed := map[hostintegration.DistributionChannel]bool{}
	for _, channel := range hostintegration.StoreManagedDistributionChannels() {
		allowed[channel] = true
	}
	for _, channel := range hostintegration.DistributionChannels() {
		if channel.StoreManaged() != allowed[channel] {
			return false
		}
	}
	return true
}
