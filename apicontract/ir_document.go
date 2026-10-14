package apicontract

type IRDocument struct {
	SchemaVersion       string         `json:"schema_version"`
	ContractID          string         `json:"contract_id"`
	SourceSchemaVersion string         `json:"source_schema_version"`
	Context             string         `json:"context"`
	Service             Service        `json:"service"`
	ClientModules       []ClientModule `json:"client_modules,omitempty"`
	Resources           []Resource     `json:"resources,omitempty"`
	Policies            []Policy       `json:"policies,omitempty"`
	Enums               []Enum         `json:"enums,omitempty"`
	SumTypes            []SumType      `json:"sum_types,omitempty"`
	Components          []IRComponent  `json:"components"`
	Operations          []IROperation  `json:"operations"`
	Scenarios           []IRScenario   `json:"scenarios,omitempty"`
}

type IRComponent struct {
	Name   string `json:"name"`
	Schema Schema `json:"schema"`
}
