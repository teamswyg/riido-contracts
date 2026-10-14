package main

type schema struct {
	Name       string     `json:"name"`
	Required   []string   `json:"required"`
	Properties []property `json:"properties"`
}

type property struct {
	Name                    string `json:"name"`
	AdditionalPropertiesRef string `json:"additional_properties_ref"`
}

type policySpec struct {
	PolicyID string   `json:"policy_id"`
	Kind     string   `json:"kind"`
	Rules    []string `json:"rules"`
}
