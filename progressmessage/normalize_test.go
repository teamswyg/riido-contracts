package progressmessage

import "testing"

func TestNormalizeLabelForCode(t *testing.T) {
	tests := []struct {
		name  string
		code  int
		label string
		want  string
	}{
		{name: "collecting trims lookup suffix", code: 1101, label: "GitHub 조회 중", want: "GitHub"},
		{name: "completed trims repeated suffixes", code: 1104, label: "테스트 실행 완료 완료", want: "테스트"},
		{name: "unknown only trims spaces", code: 9999, label: "  그대로  ", want: "그대로"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NormalizeLabelForCode(tt.code, tt.label); got != tt.want {
				t.Fatalf("NormalizeLabelForCode(%d, %q) = %q, want %q", tt.code, tt.label, got, tt.want)
			}
		})
	}
}

func TestNormalizeArgsForCodeCopiesOnlyWhenChanged(t *testing.T) {
	args := map[string]string{"label": "GitHub 조회 중", "count": "2"}
	got := NormalizeArgsForCode(1101, args)
	if got["label"] != "GitHub" || got["count"] != "2" {
		t.Fatalf("normalized args = %+v", got)
	}
	if args["label"] != "GitHub 조회 중" {
		t.Fatalf("source args mutated: %+v", args)
	}
}
