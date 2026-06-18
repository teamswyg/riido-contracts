package apicontract

type DSLDocument struct {
	SchemaVersion string         `json:"schema_version"`
	ContractID    string         `json:"contract_id"`
	Context       string         `json:"context"`
	Service       Service        `json:"service"`
	ClientModules []ClientModule `json:"client_modules,omitempty"`
	Resources     []Resource     `json:"resources,omitempty"`
	Policies      []Policy       `json:"policies,omitempty"`
	Enums         []Enum         `json:"enums,omitempty"`
	SumTypes      []SumType      `json:"sum_types,omitempty"`
	Schemas       []Schema       `json:"schemas"`
	Operations    []DSLOperation `json:"operations"`
}

type Service struct {
	Name          string `json:"name"`
	SchemaVersion string `json:"schema_version"`
}
