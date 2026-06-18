package apicontract

type OpenAPIRequestBody struct {
	Required bool                    `json:"required,omitempty"`
	Content  map[string]OpenAPIMedia `json:"content"`
}

type OpenAPIResponse struct {
	Description string                  `json:"description"`
	Content     map[string]OpenAPIMedia `json:"content,omitempty"`
}

type OpenAPIMedia struct {
	Schema map[string]any `json:"schema"`
}

type OpenAPIComponents struct {
	Schemas         map[string]map[string]any        `json:"schemas"`
	SecuritySchemes map[string]OpenAPISecurityScheme `json:"securitySchemes,omitempty"`
}

type OpenAPISecurityScheme struct {
	Type         string `json:"type"`
	In           string `json:"in,omitempty"`
	Name         string `json:"name,omitempty"`
	Scheme       string `json:"scheme,omitempty"`
	BearerFormat string `json:"bearerFormat,omitempty"`
}
