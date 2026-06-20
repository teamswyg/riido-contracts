package main

import "github.com/teamswyg/riido-contracts/hostintegration"

func providerStatusStrings() []string {
	out := make([]string, 0, len(hostintegration.ProviderRoutingStatuses()))
	for _, status := range hostintegration.ProviderRoutingStatuses() {
		out = append(out, string(status))
	}
	return out
}
