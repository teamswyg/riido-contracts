package ir

func validateSystemScope(e CanonicalEvent) []EnvelopeViolation {
	v := forbid(e, []string{
		"TaskID", "RunID", "RuntimeID", "CapabilityFingerprint",
		"ProviderKind", "ProtocolKind", "ProviderVersion",
		"AdapterID", "AdapterVersion", "ProtocolVersion",
		"NativeConfigVersion",
	})
	if e.FSMVersion != 0 {
		v = append(v, EnvelopeViolation{
			Code:   "FORBIDDEN_FIELD",
			Field:  "FSMVersion",
			Detail: "FSMVersion must be 0 for SystemScope",
		})
	}
	return v
}
