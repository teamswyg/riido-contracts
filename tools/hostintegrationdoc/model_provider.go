package main

import "github.com/teamswyg/riido-contracts/hostintegration"

func providerStatusStrings() []string {
	out := make([]string, 0, len(hostintegration.ProviderRoutingStatuses()))
	for _, status := range hostintegration.ProviderRoutingStatuses() {
		out = append(out, string(status))
	}
	return out
}

func providerHealthStatusStrings() []string {
	out := make([]string, 0, len(hostintegration.ProviderHealthStatuses()))
	for _, status := range hostintegration.ProviderHealthStatuses() {
		out = append(out, string(status))
	}
	return out
}

func providerDiagnosticCodeStrings() []string {
	out := make([]string, 0, len(hostintegration.ProviderDiagnosticCodes()))
	for _, code := range hostintegration.ProviderDiagnosticCodes() {
		out = append(out, string(code))
	}
	return out
}
