package ir

// ValidateEnvelope checks that a CanonicalEvent satisfies the scope-aware
// envelope rules from ir-schema-versioning.md §1.5.
func ValidateEnvelope(e CanonicalEvent) []EnvelopeViolation {
	var v []EnvelopeViolation
	if !e.Scope.IsValid() {
		return append(v, EnvelopeViolation{
			Code:   "UNKNOWN_SCOPE",
			Field:  "Scope",
			Detail: string(e.Scope),
		})
	}
	v = append(v, validateCommonEnvelope(e)...)
	v = append(v, validateFakePlaceholders(e)...)
	v = append(v, validateScopedEnvelope(e)...)
	return v
}

func validateScopedEnvelope(e CanonicalEvent) []EnvelopeViolation {
	switch e.Scope {
	case EventScopeSystem:
		return validateSystemScope(e)
	case EventScopeRuntime:
		return validateRuntimeScope(e)
	case EventScopeTask:
		return validateTaskScope(e)
	case EventScopeRun:
		return validateRunScope(e)
	}
	return nil
}
