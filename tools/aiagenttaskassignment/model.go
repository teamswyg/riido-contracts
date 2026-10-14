package main

type model struct {
	Manifest              manifest
	Operations            []operation
	Schemas               []schema
	Policies              []policySpec
	ScenarioCount         int
	DSLIRMatch            bool
	OpenAPIMatch          bool
	ForbiddenFieldsAbsent bool
	NoDiffPathsAbsent     bool
}
