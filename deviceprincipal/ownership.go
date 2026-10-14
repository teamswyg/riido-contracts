package deviceprincipal

type OwnershipEdge struct {
	From string
	To   string
}

func OwnershipEdges() []OwnershipEdge {
	return []OwnershipEdge{
		{From: string(UserPrincipal), To: "Device"},
		{From: "Device", To: "Runtime"},
		{From: "Workspace", To: "Agent"},
		{From: "Agent", To: "selected Runtime binding"},
	}
}

func BindingSources() []string {
	return []string{
		string(DevicePrincipal),
		"latest daemon runtime snapshot",
		"workspace-scoped agents with saved runtime_id",
	}
}
