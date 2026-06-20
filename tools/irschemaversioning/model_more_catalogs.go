package main

func fakePlaceholderValues() []string {
	return []string{"unknown", "none", "pending", "n/a", "na", "tbd", "-"}
}

func violationCodes() []string {
	return []string{"MISSING_FIELD", "FORBIDDEN_FIELD", "FAKE_PLACEHOLDER", "UNKNOWN_SCOPE", "INVALID_FSMVERSION"}
}

func nativeConfigClasses() []string {
	return []string{"forbidden", "pre-execute", "required", "phase-dependent"}
}

func scopeRules() []scopeRule {
	return []scopeRule{
		{Scope: "system", Required: 0, Forbidden: 12},
		{Scope: "runtime", Required: 1, Forbidden: 4},
		{Scope: "task", Required: 1, Forbidden: 10, Conditional: 1},
		{Scope: "run", Required: 10, Conditional: 2},
	}
}
