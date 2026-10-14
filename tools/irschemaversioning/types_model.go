package main

type model struct {
	Manifest                 manifest
	CanonicalEventFields     int
	EventScopeCount          int
	CommonRequiredCount      int
	ActorIDConditional       int
	RunRequiredFieldCount    int
	FakePlaceholderFields    int
	FakePlaceholderValues    int
	ViolationCodeCount       int
	NativeConfigClassCount   int
	RunContextFields         int
	ValidateEntrypoints      int
	ScopeRules               []scopeRule
	CanonicalEventFieldNames []string
	RunContextFieldNames     []string
}
