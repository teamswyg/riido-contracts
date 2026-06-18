package catalog

const (
	DefaultClaudeModelID   = "claude-default"
	DefaultCodexModelID    = "codex-default"
	DefaultCursorModelID   = "cursor-auto"
	DefaultGeminiModelID   = "gemini-default"
	DefaultOpenClawModelID = "openclaw-default"
	DefaultRuntimeModelID  = "runtime-default"
)

func DefaultModelID(value string) string {
	switch Normalize(value) {
	case KindClaude, KindClaudeCode:
		return DefaultClaudeModelID
	case KindCodex:
		return DefaultCodexModelID
	case KindCursor:
		return DefaultCursorModelID
	case KindGemini:
		return DefaultGeminiModelID
	case KindOpenClaw:
		return DefaultOpenClawModelID
	default:
		return DefaultRuntimeModelID
	}
}
