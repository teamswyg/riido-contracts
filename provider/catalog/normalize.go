package catalog

import "strings"

func Normalize(value string) Kind {
	normalized := strings.TrimSpace(strings.ToLower(value))
	switch normalized {
	case "claude-code":
		return KindClaudeCode
	default:
		return Kind(normalized)
	}
}

func String(value string) string {
	return string(Normalize(value))
}
