package deviceprincipal

func ExcludedFallbacks() []string {
	return []string{
		"team_id",
		"teamId",
		"OpenAPI task context URL",
		"Open API key",
		"X-Workspace-Api-Key",
	}
}
