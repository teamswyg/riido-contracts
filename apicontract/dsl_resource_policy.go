package apicontract

type Resource struct {
	Name       string           `json:"name"`
	OwnerField string           `json:"owner_field,omitempty"`
	Visibility []VisibilityRule `json:"visibility,omitempty"`
}

type VisibilityRule struct {
	Name  string   `json:"name"`
	Read  []string `json:"read,omitempty"`
	Write []string `json:"write,omitempty"`
}

type Policy struct {
	PolicyID    string   `json:"policy_id"`
	Kind        string   `json:"kind"`
	Description string   `json:"description"`
	Rules       []string `json:"rules,omitempty"`
}
