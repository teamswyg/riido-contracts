package main

type bundle struct {
	SchemaVersion string     `json:"schema_version"`
	Source        string     `json:"source"`
	Contracts     []contract `json:"contracts"`
}

type contract struct {
	ContractID     string         `json:"contract_id"`
	Context        string         `json:"context"`
	Service        service        `json:"service"`
	SourceFiles    sourceFiles    `json:"source_files"`
	OperationCount int            `json:"operation_count"`
	Operations     []operation    `json:"operations"`
	Schemas        map[string]any `json:"schemas"`
}

type sourceFiles struct {
	IR      string `json:"ir"`
	OpenAPI string `json:"openapi"`
}

type service struct {
	Name          string `json:"name"`
	SchemaVersion string `json:"schema_version"`
}
