package main

type model struct {
	Manifest      manifest
	Operation     operation
	Schemas       []schema
	Policy        policySpec
	ScenarioCount int
	DSLIRMatch    bool
	OpenAPIMatch  bool
	MapShapePass  bool
}
