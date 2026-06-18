package ir

// ValidateEnvelopeWithRunContext is the dynamic-context companion to
// ValidateEnvelope. It performs the same static checks first, then resolves
// PhaseDependent NativeConfigVersion requirements using the provided run
// context.
func ValidateEnvelopeWithRunContext(e CanonicalEvent, ctx RunContext) []EnvelopeViolation {
	v := ValidateEnvelope(e)
	if e.Scope != EventScopeRun {
		return v
	}
	if NativeConfigRequirementOf(e.Type) != NativeConfigPhaseDependent {
		return v
	}
	if ctx.NativeConfigEstablished && e.NativeConfigVersion == "" {
		v = append(v, EnvelopeViolation{
			Code:   "MISSING_FIELD",
			Field:  "NativeConfigVersion",
			Detail: "PhaseDependent event " + string(e.Type) + " requires NCV after run has crossed NativeConfigInjected",
		})
	}
	return v
}
