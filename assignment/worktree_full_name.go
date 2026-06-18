package assignment

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
