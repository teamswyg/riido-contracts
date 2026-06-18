package ir

// NativeConfigRequirement classifies an EventType by how it treats
// NativeConfigVersion (NCV) in RunScope contexts.
//
// Source of truth: ir-event-log.md §3.0.1 / §3.0.2,
// ir-schema-versioning.md §1.5.3.1 / §1.5.3.2.
type NativeConfigRequirement int

const (
	// NativeConfigForbidden — this EventType is never RunScope-typical;
	// NCV must not be present regardless of context.
	NativeConfigForbidden NativeConfigRequirement = iota

	// NativeConfigOptionalPreExecute — RunScope event that occurs before
	// native config materialization; NCV may legitimately be absent.
	NativeConfigOptionalPreExecute

	// NativeConfigRequired — RunScope event that occurs after native
	// config has been injected; NCV must be present.
	NativeConfigRequired

	// NativeConfigPhaseDependent — RunScope event that may occur in
	// either pre-execute OR execution-bound phase.
	NativeConfigPhaseDependent
)

// NativeConfigRequirementOf returns the static classification for an EventType.
func NativeConfigRequirementOf(t EventType) NativeConfigRequirement {
	return t.Code().NativeConfigRequirement()
}
