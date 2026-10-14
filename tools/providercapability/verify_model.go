package main

import "fmt"

func verifyModel(model model) error {
	m := model.Manifest
	checks := map[string][2]int{
		"provider capability fields": {model.ProviderCapabilityFields, m.ExpectedProviderCapabilityFields},
		"fingerprint input fields":   {model.FingerprintInputFields, m.ExpectedFingerprintInputFields},
		"protocol count":             {len(model.Protocols), m.ExpectedProtocolCount},
		"event stream formats":       {len(model.EventStreamFormats), m.ExpectedEventStreamFormatCount},
		"protocol maturities":        {len(model.ProtocolMaturities), m.ExpectedProtocolMaturityCount},
		"compatibility statuses":     {len(model.CompatibilityStatuses), m.ExpectedCompatibilityStatusCount},
		"critical arg sets":          {model.CriticalArgSetCount, m.ExpectedProtocolCriticalArgSets},
	}
	for name, values := range checks {
		if values[0] != values[1] {
			return fmt.Errorf("%s = %d, want %d", name, values[0], values[1])
		}
	}
	return nil
}
