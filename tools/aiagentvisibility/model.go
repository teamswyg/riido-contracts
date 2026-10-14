package main

type model struct {
	Manifest      manifest
	Operations    []operation
	Schemas       []schema
	Policy        policySpec
	Enum          enumSpec
	ScenarioCount int
	DSLIRMatch    bool
	OpenAPIMatch  bool
}
