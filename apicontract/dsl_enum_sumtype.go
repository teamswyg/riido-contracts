package apicontract

type Enum struct {
	Name        string      `json:"name"`
	Type        string      `json:"type"`
	Description string      `json:"description,omitempty"`
	Values      []EnumValue `json:"values"`
}

type EnumValue struct {
	Value       string `json:"value"`
	Description string `json:"description,omitempty"`
}

type SumType struct {
	Name          string           `json:"name"`
	Discriminator string           `json:"discriminator"`
	Description   string           `json:"description,omitempty"`
	Variants      []SumTypeVariant `json:"variants"`
}

type SumTypeVariant struct {
	Kind        string `json:"kind"`
	Schema      string `json:"schema"`
	Description string `json:"description,omitempty"`
}
