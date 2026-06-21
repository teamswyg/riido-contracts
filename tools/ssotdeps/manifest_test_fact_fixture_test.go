package main

func minimalFact() fact {
	return fact{
		ID:             "agent-concept",
		Fact:           "Agent means a task-assignable abstraction of a configured runtime",
		HumanDocPhrase: "Agent means a task-assignable abstraction of a configured runtime",
		SourceRefs: []sourceRef{
			{
				Repo:           localRepo,
				Path:           "docs/20-domain/ai-agent-policy.md",
				RequiredPhrase: "Agent",
			},
		},
		Owner: ownerRef{
			Repo: localRepo,
			Path: "docs/20-domain/ai-agent-policy.md",
		},
		Downstreams: []downstream{
			{
				Repo:       "riido-control-plane",
				LocalScope: "test projection",
			},
		},
	}
}
