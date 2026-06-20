package main

import (
	"reflect"

	capabilitypkg "github.com/teamswyg/riido-contracts/provider/capability"
)

func buildModel(m manifest) model {
	return model{
		Manifest:                 m,
		ProviderCapabilityFields: reflect.TypeOf(capabilitypkg.ProviderCapability{}).NumField(),
		FingerprintInputFields:   reflect.TypeOf(capabilitypkg.CapabilityFingerprintInput{}).NumField(),
		Protocols:                protocolRows(),
		EventStreamFormats:       stringValues(capabilitypkg.AllEventStreamFormats()),
		ProtocolMaturities:       stringValues(capabilitypkg.AllProtocolMaturities()),
		CompatibilityStatuses:    stringValues(capabilitypkg.AllCompatibilityStatuses()),
		CriticalArgSetCount:      criticalArgSetCount(),
	}
}

func stringValues[T ~string](values []T) []string {
	out := make([]string, len(values))
	for i, value := range values {
		out[i] = string(value)
	}
	return out
}
