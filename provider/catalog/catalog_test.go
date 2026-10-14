package catalog

import "testing"

func TestNormalizeProviderKind(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  Kind
	}{
		{name: "claude code alias", input: " Claude-Code ", want: KindClaudeCode},
		{name: "codex", input: "CODEX", want: KindCodex},
		{name: "unknown stable lowercase", input: " Custom_Runtime ", want: Kind("custom_runtime")},
		{name: "unknown preserves separator", input: "Future-Provider", want: Kind("future-provider")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Normalize(tt.input); got != tt.want {
				t.Fatalf("Normalize(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestModelOverrideSuppressesDefaults(t *testing.T) {
	if got := ModelOverride(string(KindCodex), DefaultCodexModelID); got != "" {
		t.Fatalf("default codex model override = %q, want empty", got)
	}
	if got := ModelOverride("other", DefaultRuntimeModelID); got != "" {
		t.Fatalf("default runtime model override = %q, want empty", got)
	}
	if got := ModelOverride(string(KindCodex), "gpt-custom"); got != "gpt-custom" {
		t.Fatalf("custom codex model override = %q", got)
	}
}

func TestClaudeFamilyIncludesClaudeCodeSurface(t *testing.T) {
	if !IsClaudeFamily("claude") || !IsClaudeFamily("claude_code") {
		t.Fatal("claude family must include claude and claude_code")
	}
}
