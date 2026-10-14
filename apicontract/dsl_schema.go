package apicontract

type Schema struct {
	Name        string     `json:"name"`
	Description string     `json:"description,omitempty"`
	Type        string     `json:"type"`
	Required    []string   `json:"required,omitempty"`
	Properties  []Property `json:"properties,omitempty"`
}

type Property struct {
	Name                    string    `json:"name,omitempty"`
	Type                    string    `json:"type,omitempty"`
	Description             string    `json:"description,omitempty"`
	Format                  string    `json:"format,omitempty"`
	MaxLength               *int      `json:"max_length,omitempty"`
	Enum                    []string  `json:"enum,omitempty"`
	Ref                     string    `json:"ref,omitempty"`
	Items                   *Property `json:"items,omitempty"`
	AdditionalProperties    bool      `json:"additional_properties,omitempty"`
	AdditionalPropertiesRef string    `json:"additional_properties_ref,omitempty"`
}
