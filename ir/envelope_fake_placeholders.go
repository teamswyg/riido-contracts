package ir

func validateFakePlaceholders(e CanonicalEvent) []EnvelopeViolation {
	var v []EnvelopeViolation
	for _, f := range identityFields(e) {
		if isFakePlaceholder(f.val) {
			v = append(v, EnvelopeViolation{
				Code:   "FAKE_PLACEHOLDER",
				Field:  f.name,
				Detail: f.val,
			})
		}
	}
	return v
}

func identityFields(e CanonicalEvent) []envelopeField {
	return []envelopeField{
		{"RuntimeID", e.RuntimeID},
		{"CapabilityFingerprint", e.CapabilityFingerprint},
		{"NativeConfigVersion", e.NativeConfigVersion},
		{"ProviderVersion", e.ProviderVersion},
		{"AdapterID", e.AdapterID},
		{"AdapterVersion", e.AdapterVersion},
		{"ProtocolVersion", e.ProtocolVersion},
		{"TaskID", e.TaskID},
		{"RunID", e.RunID},
	}
}
