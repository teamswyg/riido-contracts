package capability

import (
	"slices"
	"testing"
)

func TestProtocolCriticalArgsCatalog(t *testing.T) {
	tests := []struct {
		name string
		kind ProtocolKind
		want []string
	}{
		{
			name: "claude stream json",
			kind: ProtocolClaudeStreamJSON,
			want: []string{"-p", "--print", "--output-format", "--input-format", "--permission-mode", "--mcp-config", "--strict-mcp-config", "--verbose"},
		},
		{
			name: "codex exec jsonl",
			kind: ProtocolCodexExecJSONL,
			want: []string{"--json"},
		},
		{
			name: "codex app server",
			kind: ProtocolCodexAppServer,
			want: []string{"--listen"},
		},
		{
			name: "openclaw agent json",
			kind: ProtocolOpenClawAgentJSON,
			want: []string{"--local", "--json", "--session-id", "--message", "--model", "--system-prompt"},
		},
		{
			name: "cursor agent stream json",
			kind: ProtocolCursorAgentStreamJSON,
			want: []string{"-p", "--output-format", "--yolo"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProtocolCriticalArgs(tt.kind); !slices.Equal(got, tt.want) {
				t.Fatalf("ProtocolCriticalArgs(%q) = %v, want %v", tt.kind, got, tt.want)
			}
		})
	}
}

func TestProtocolCriticalArgsReturnsCopy(t *testing.T) {
	got := ProtocolCriticalArgs(ProtocolClaudeStreamJSON)
	got[0] = "--mutated"
	if again := ProtocolCriticalArgs(ProtocolClaudeStreamJSON); again[0] != "-p" {
		t.Fatalf("ProtocolCriticalArgs returned mutable backing store: %v", again)
	}
}
