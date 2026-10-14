package ir

func validateRunScopeNativeConfig(e CanonicalEvent) []EnvelopeViolation {
	switch NativeConfigRequirementOf(e.Type) {
	case NativeConfigForbidden:
		return []EnvelopeViolation{{
			Code:   "FORBIDDEN_FIELD",
			Field:  "Scope",
			Detail: "EventType " + string(e.Type) + " is not allowed in RunScope",
		}}
	case NativeConfigRequired:
		if e.NativeConfigVersion == "" {
			return []EnvelopeViolation{{
				Code:   "MISSING_FIELD",
				Field:  "NativeConfigVersion",
				Detail: "required for execution-bound RunScope event " + string(e.Type),
			}}
		}
	case NativeConfigOptionalPreExecute, NativeConfigPhaseDependent:
		return nil
	}
	return nil
}
