package assignment

import "testing"

func TestNormalizePublicGitHubRepositoryFullName(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want string
	}{
		{name: "plain", raw: "teamswyg/riido-daemon", want: "teamswyg/riido-daemon"},
		{name: "trims whitespace", raw: " teamswyg/riido-daemon ", want: "teamswyg/riido-daemon"},
		{name: "trims edge slashes", raw: "/teamswyg/riido-daemon/", want: "teamswyg/riido-daemon"},
		{name: "preserves allowed case and punctuation", raw: "Team.Swyg/riido_daemon-1", want: "Team.Swyg/riido_daemon-1"},
		{name: "rejects missing repo", raw: "teamswyg", want: ""},
		{name: "rejects extra path", raw: "teamswyg/riido-daemon/extra", want: ""},
		{name: "rejects query", raw: "teamswyg/riido-daemon?token=secret", want: ""},
		{name: "rejects encoded query", raw: "teamswyg/riido-daemon%3Ftoken=secret", want: ""},
		{name: "rejects owner dot", raw: "./riido-daemon", want: ""},
		{name: "rejects repo dotdot", raw: "teamswyg/..", want: ""},
		{name: "rejects unicode", raw: "teamswyg/뤼이도", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NormalizePublicGitHubRepositoryFullName(tt.raw)
			if got != tt.want {
				t.Fatalf("NormalizePublicGitHubRepositoryFullName(%q) = %q, want %q", tt.raw, got, tt.want)
			}
			if (got != "") != IsPublicGitHubRepositoryFullName(tt.raw) {
				t.Fatalf("IsPublicGitHubRepositoryFullName(%q) disagrees with normalized value %q", tt.raw, got)
			}
		})
	}
}

func TestNormalizePublicGitHubRepositoryURL(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want string
	}{
		{name: "plain", raw: "https://github.com/teamswyg/riido-daemon", want: "https://github.com/teamswyg/riido-daemon"},
		{name: "trims whitespace", raw: " https://github.com/teamswyg/riido-daemon ", want: "https://github.com/teamswyg/riido-daemon"},
		{name: "canonicalizes trailing slash", raw: "https://github.com/teamswyg/riido-daemon/", want: "https://github.com/teamswyg/riido-daemon"},
		{name: "preserves allowed case and punctuation", raw: "https://github.com/Team.Swyg/riido_daemon-1", want: "https://github.com/Team.Swyg/riido_daemon-1"},
		{name: "rejects userinfo", raw: "https://token:secret@github.com/teamswyg/riido-daemon", want: ""},
		{name: "rejects query", raw: "https://github.com/teamswyg/riido-daemon?token=secret", want: ""},
		{name: "rejects force query", raw: "https://github.com/teamswyg/riido-daemon?", want: ""},
		{name: "rejects fragment", raw: "https://github.com/teamswyg/riido-daemon#token=secret", want: ""},
		{name: "rejects non github", raw: "https://example.com/teamswyg/riido-daemon", want: ""},
		{name: "rejects http", raw: "http://github.com/teamswyg/riido-daemon", want: ""},
		{name: "rejects port", raw: "https://github.com:443/teamswyg/riido-daemon", want: ""},
		{name: "rejects extra path", raw: "https://github.com/teamswyg/riido-daemon/tree/main", want: ""},
		{name: "rejects encoded query in path", raw: "https://github.com/teamswyg/riido-daemon%3Ftoken=secret", want: ""},
		{name: "rejects missing repo", raw: "https://github.com/teamswyg", want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NormalizePublicGitHubRepositoryURL(tt.raw)
			if got != tt.want {
				t.Fatalf("NormalizePublicGitHubRepositoryURL(%q) = %q, want %q", tt.raw, got, tt.want)
			}
			if (got != "") != IsPublicGitHubRepositoryURL(tt.raw) {
				t.Fatalf("IsPublicGitHubRepositoryURL(%q) disagrees with normalized value %q", tt.raw, got)
			}
		})
	}
}
