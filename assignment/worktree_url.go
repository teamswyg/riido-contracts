package assignment

import (
	"net/url"
	"strings"
)

const publicGitHubRepositoryURLPrefix = "https://github.com/"

// NormalizePublicGitHubRepositoryURL returns a canonical public GitHub
// repository URL accepted by AssignmentWorktree, or an empty string when the URL
// includes unsupported or credential-bearing components.
func NormalizePublicGitHubRepositoryURL(raw string) string {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return ""
	}
	parsed, err := url.Parse(raw)
	if err != nil || !validPublicGitHubRepositoryURL(parsed) {
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

func validPublicGitHubRepositoryURL(parsed *url.URL) bool {
	if parsed.Scheme != "https" || !strings.EqualFold(parsed.Host, "github.com") {
		return false
	}
	if parsed.User != nil {
		return false
	}
	return parsed.RawQuery == "" && !parsed.ForceQuery &&
		parsed.Fragment == "" && parsed.RawFragment == ""
}
