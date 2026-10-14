package apicontract

type OpenAPISpec struct {
	OpenAPI            string                 `json:"openapi"`
	Info               OpenAPIInfo            `json:"info"`
	Tags               []OpenAPITag           `json:"tags,omitempty"`
	RiidoClientModules []ClientModule         `json:"x-riido-client-modules,omitempty"`
	Paths              map[string]OpenAPIPath `json:"paths"`
	Components         OpenAPIComponents      `json:"components"`
}

type OpenAPIInfo struct {
	Title   string `json:"title"`
	Version string `json:"version"`
}

type OpenAPITag struct {
	Name string `json:"name"`
}

type OpenAPIPath map[string]OpenAPIOperation
