package assignment

import "strings"

func normalizePublicGitHubRepositoryParts(raw string) (string, string, bool) {
	parts := strings.Split(strings.Trim(strings.TrimSpace(raw), "/"), "/")
	if len(parts) != 2 {
		return "", "", false
	}
	if !validPublicGitHubRepositoryPart(parts[0]) || !validPublicGitHubRepositoryPart(parts[1]) {
		return "", "", false
	}
	return parts[0], parts[1], true
}

func validPublicGitHubRepositoryPart(part string) bool {
	if part == "" || part == "." || part == ".." {
		return false
	}
	for _, ch := range part {
		if !validPublicGitHubRepositoryPartRune(ch) {
			return false
		}
	}
	return true
}

func validPublicGitHubRepositoryPartRune(ch rune) bool {
	switch {
	case ch >= 'a' && ch <= 'z':
	case ch >= 'A' && ch <= 'Z':
	case ch >= '0' && ch <= '9':
	case ch == '-' || ch == '_' || ch == '.':
	default:
		return false
	}
	return true
}
