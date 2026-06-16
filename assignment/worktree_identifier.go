package assignment

import (
	"net/url"
	"strings"
)

const publicGitHubRepositoryURLPrefix = "https://github.com/"

// NormalizePublicGitHubRepositoryFullName returns the canonical owner/repo
// identifier accepted by AssignmentWorktree, or an empty string when the input
// is not a public GitHub repository identifier.
func NormalizePublicGitHubRepositoryFullName(raw string) string {
	owner, repo, ok := normalizePublicGitHubRepositoryParts(raw)
	if !ok {
		return ""
	}
	return owner + "/" + repo
}

// IsPublicGitHubRepositoryFullName reports whether raw is accepted as a public
// GitHub owner/repo identifier for AssignmentWorktree.
func IsPublicGitHubRepositoryFullName(raw string) bool {
	return NormalizePublicGitHubRepositoryFullName(raw) != ""
}

// NormalizePublicGitHubRepositoryURL returns a canonical public GitHub
// repository URL accepted by AssignmentWorktree, or an empty string when the URL
// includes unsupported or credential-bearing components.
func NormalizePublicGitHubRepositoryURL(raw string) string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return ""
	}
	parsed, err := url.Parse(raw)
	if err != nil {
		return ""
	}
	if parsed.Scheme != "https" || !strings.EqualFold(parsed.Host, "github.com") || parsed.User != nil {
		return ""
	}
	if parsed.RawQuery != "" || parsed.ForceQuery || parsed.Fragment != "" || parsed.RawFragment != "" {
		return ""
	}
	fullName := NormalizePublicGitHubRepositoryFullName(parsed.Path)
	if fullName == "" {
		return ""
	}
	return publicGitHubRepositoryURLPrefix + fullName
}

// IsPublicGitHubRepositoryURL reports whether raw is accepted as a public GitHub
// repository URL for AssignmentWorktree.
func IsPublicGitHubRepositoryURL(raw string) bool {
	return NormalizePublicGitHubRepositoryURL(raw) != ""
}

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
		switch {
		case ch >= 'a' && ch <= 'z':
		case ch >= 'A' && ch <= 'Z':
		case ch >= '0' && ch <= '9':
		case ch == '-' || ch == '_' || ch == '.':
		default:
			return false
		}
	}
	return true
}
