package main

type model struct {
	Manifest               manifest
	Operations             []operation
	OnboardingOperations   []operation
	DirectCreateOperations []operation
	FixtureSchema          schema
	ListSchema             schema
	CreateRequestSchema    schema
	FixtureFields          []string
	CreateRequestFields    []string
	ScenarioCount          int
	DSLIRMatch             bool
	OpenAPIMatch           bool
	NoDiffPathsClean       bool
}
