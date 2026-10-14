package ir

func validateCommonEnvelope(e CanonicalEvent) []EnvelopeViolation {
	var v []EnvelopeViolation
	v = append(v, requireNonEmpty("EventID", e.EventID)...)
	if e.OccurredAt.IsZero() {
		v = append(v, EnvelopeViolation{Code: "MISSING_FIELD", Field: "OccurredAt"})
	}
	if e.EventSchemaVersion <= 0 {
		v = append(v, EnvelopeViolation{Code: "MISSING_FIELD", Field: "EventSchemaVersion", Detail: "must be >= 1"})
	}
	v = append(v, requireNonEmpty("Type", string(e.Type))...)
	v = append(v, requireNonEmpty("ActorKind", string(e.ActorKind))...)
	if e.ActorKind != ActorSystem && e.ActorID == "" {
		v = append(v, EnvelopeViolation{Code: "MISSING_FIELD", Field: "ActorID", Detail: "required unless ActorKind=system"})
	}
	v = append(v, requireNonEmpty("RiidoDaemonVersion", e.RiidoDaemonVersion)...)
	v = append(v, requireNonEmpty("PolicyBundleVersion", e.PolicyBundleVersion)...)
	return v
}
