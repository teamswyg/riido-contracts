package ir

func validateTaskScope(e CanonicalEvent) []EnvelopeViolation {
	v := requireNonEmpty("TaskID", e.TaskID)
	v = append(v, forbid(e, []string{
		"RunID", "RuntimeID", "CapabilityFingerprint",
		"ProviderKind", "ProtocolKind", "ProviderVersion",
		"AdapterID", "AdapterVersion", "ProtocolVersion",
		"NativeConfigVersion",
	})...)
	if e.Type.IsTransition() && e.FSMVersion <= 0 {
		v = append(v, EnvelopeViolation{
			Code:   "INVALID_FSMVERSION",
			Field:  "FSMVersion",
			Detail: "TaskScope transition events require FSMVersion >= 1",
		})
	}
	return v
}
