package ir

type envelopeField struct {
	name string
	val  string
}

func requireNonEmpty(field, val string) []EnvelopeViolation {
	if val == "" {
		return []EnvelopeViolation{{Code: "MISSING_FIELD", Field: field}}
	}
	return nil
}

func forbid(e CanonicalEvent, fields []string) []EnvelopeViolation {
	var v []EnvelopeViolation
	for _, name := range fields {
		if val := fieldValue(e, name); val != "" {
			v = append(v, EnvelopeViolation{
				Code:   "FORBIDDEN_FIELD",
				Field:  name,
				Detail: "must be absent for this scope",
			})
		}
	}
	return v
}
