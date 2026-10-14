package fsmmeta

import "testing"

func TestFSMPatternCodeParsing(t *testing.T) {
	if got := ParseFSMPatternCode("flat"); got != FSMPatternCodeFlat {
		t.Fatalf("ParseFSMPatternCode(flat) = %v, want %v", got, FSMPatternCodeFlat)
	}
	if got := ParseFSMPatternCode("hierarchical"); got != FSMPatternCodeHierarchical {
		t.Fatalf("ParseFSMPatternCode(hierarchical) = %v, want %v", got, FSMPatternCodeHierarchical)
	}
	if got := ParseFSMPatternCode("unknown"); got != FSMPatternCodeUnknown {
		t.Fatalf("ParseFSMPatternCode(unknown) = %v, want %v", got, FSMPatternCodeUnknown)
	}
}

func TestFSMPatternRoundTrip(t *testing.T) {
	codes := AllFSMPatternCodes()
	values := AllFSMPatterns()
	if len(codes) != len(values) {
		t.Fatalf("pattern code count = %d, value count = %d", len(codes), len(values))
	}
	for index, code := range codes {
		if !code.IsKnown() {
			t.Fatalf("code %v is not known", code)
		}
		value := values[index]
		if !value.Valid() {
			t.Fatalf("value %q is not valid", value)
		}
		if value.Code() != code {
			t.Fatalf("value %q code = %v, want %v", value, value.Code(), code)
		}
		if code.FSMPattern() != value {
			t.Fatalf("code %v value = %q, want %q", code, code.FSMPattern(), value)
		}
	}
}
