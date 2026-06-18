package apicontract

type OpenAPIOperation struct {
	OperationID    string                     `json:"operationId"`
	Summary        string                     `json:"summary,omitempty"`
	Tags           []string                   `json:"tags,omitempty"`
	Parameters     []OpenAPIParameter         `json:"parameters,omitempty"`
	RequestBody    *OpenAPIRequestBody        `json:"requestBody,omitempty"`
	Responses      map[string]OpenAPIResponse `json:"responses"`
	Security       []map[string][]string      `json:"security,omitempty"`
	RiidoScopes    []string                   `json:"x-riido-auth-scopes,omitempty"`
	RiidoRBAC      string                     `json:"x-riido-rbac-policy,omitempty"`
	RiidoKind      string                     `json:"x-riido-operation-kind,omitempty"`
	RiidoClient    *ClientMeta                `json:"x-riido-client,omitempty"`
	RiidoScenarios []string                   `json:"x-riido-scenarios,omitempty"`
}

type OpenAPIParameter struct {
	Name     string         `json:"name"`
	In       string         `json:"in"`
	Required bool           `json:"required"`
	Schema   map[string]any `json:"schema"`
}
