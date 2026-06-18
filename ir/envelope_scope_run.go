package ir

func validateRunScope(e CanonicalEvent) []EnvelopeViolation {
	var v []EnvelopeViolation
	for _, f := range runScopeRequiredFields(e) {
		if f.val == "" {
			v = append(v, EnvelopeViolation{Code: "MISSING_FIELD", Field: f.name, Detail: "required for RunScope"})
		}
	}
	v = append(v, validateRunScopeNativeConfig(e)...)
	if e.Type.IsTransition() && e.FSMVersion <= 0 {
		v = append(v, EnvelopeViolation{
			Code:   "INVALID_FSMVERSION",
			Field:  "FSMVersion",
			Detail: "RunScope transition events require FSMVersion >= 1",
		})
	}
	return v
}

func runScopeRequiredFields(e CanonicalEvent) []envelopeField {
	return []envelopeField{
		{"TaskID", e.TaskID},
		{"RunID", e.RunID},
		{"RuntimeID", e.RuntimeID},
		{"CapabilityFingerprint", e.CapabilityFingerprint},
		{"ProviderKind", e.ProviderKind},
		{"ProtocolKind", e.ProtocolKind},
		{"ProviderVersion", e.ProviderVersion},
		{"AdapterID", e.AdapterID},
		{"AdapterVersion", e.AdapterVersion},
		{"ProtocolVersion", e.ProtocolVersion},
	}
}
