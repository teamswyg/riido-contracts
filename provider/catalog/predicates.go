package catalog

func IsCodex(value string) bool {
	return Normalize(value) == KindCodex
}

func IsClaudeFamily(value string) bool {
	switch Normalize(value) {
	case KindClaude, KindClaudeCode:
		return true
	default:
		return false
	}
}
