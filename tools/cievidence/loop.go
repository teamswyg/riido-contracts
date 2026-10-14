package main

func ciLoop() evidenceLoop {
	return evidenceLoop{
		Observation:   "Contracts broad CI workflows already verify stdlib-only dependencies, enumgen, fsmgen, lint, and tests, but they did not publish a machine-readable evidence artifact.",
		Hypothesis:    "A compact CI evidence artifact can make broad workflow coverage observable without duplicating domain-specific generated readers.",
		Execute:       "Inspect the target workflow for the required command surface and publish the command presence result as JSON from public CI.",
		Evaluate:      "The evidence tool fails when a required broad CI command disappears or when an unknown workflow id is requested.",
		Retrospective: "Broad CI remains a thin orchestration layer while evidence consumers can verify which contract gates it actually ran.",
	}
}
