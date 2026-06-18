package ir

import "testing"

func TestNativeConfigRequirementOf(t *testing.T) {
	for _, c := range nativeConfigRequirementCases() {
		t.Run(string(c.evt), func(t *testing.T) {
			if got := NativeConfigRequirementOf(c.evt); got != c.want {
				t.Errorf("NativeConfigRequirementOf(%s) = %v, want %v", c.evt, got, c.want)
			}
		})
	}
}

type nativeConfigRequirementCase struct {
	evt  EventType
	want NativeConfigRequirement
}

func nativeConfigRequirementCases() []nativeConfigRequirementCase {
	return append(
		nativeConfigOptionalCases(),
		append(nativeConfigRequiredCases(), nativeConfigForbiddenCases()...)...,
	)
}
