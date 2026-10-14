package ir

// EnvelopeViolation describes a single rule violation found by ValidateEnvelope.
type EnvelopeViolation struct {
	Code   string // "MISSING_FIELD" | "FORBIDDEN_FIELD" | "FAKE_PLACEHOLDER" | "UNKNOWN_SCOPE" | "INVALID_FSMVERSION"
	Field  string
	Detail string
}
