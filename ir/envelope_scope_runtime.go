package ir

func validateRuntimeScope(e CanonicalEvent) []EnvelopeViolation {
	v := requireNonEmpty("RuntimeID", e.RuntimeID)
	v = append(v, forbid(e, []string{
		"TaskID", "RunID", "NativeConfigVersion",
	})...)
	if e.FSMVersion != 0 {
		v = append(v, EnvelopeViolation{
			Code:   "FORBIDDEN_FIELD",
			Field:  "FSMVersion",
			Detail: "FSMVersion must be 0 for RuntimeScope",
		})
	}
	return v
}
