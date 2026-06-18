package assignment

type AssignmentWorktree struct {
	RepositoryFullName string `json:"repository_full_name,omitempty"`
	RepositoryURL      string `json:"repository_url,omitempty"`
	BranchName         string `json:"branch_name,omitempty"`
	IsPrivate          bool   `json:"is_private,omitempty"`
	Source             string `json:"source,omitempty"`
}
