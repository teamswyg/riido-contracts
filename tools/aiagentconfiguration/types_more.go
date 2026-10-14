package main

type sumType struct {
	Name     string    `json:"name"`
	Variants []variant `json:"variants"`
}

type variant struct {
	Kind   string `json:"kind"`
	Schema string `json:"schema"`
}

type openAPIDoc struct {
	Paths map[string]map[string]openAPIOperation `json:"paths"`
}

type openAPIOperation struct {
	OperationID string `json:"operationId"`
}

type model struct {
	Manifest      manifest
	Operations    []operation
	Schemas       []schema
	Policies      []policySpec
	Enums         []enumSpec
	ScenarioCount int
	DSLIRMatch    bool
	OpenAPIMatch  bool
	StreamPass    bool
}
