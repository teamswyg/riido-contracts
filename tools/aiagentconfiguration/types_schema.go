package main

type schema struct {
	Name       string     `json:"name"`
	Type       string     `json:"type"`
	Required   []string   `json:"required"`
	Properties []property `json:"properties"`
}

type property struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Ref       string `json:"ref"`
	MaxLength int    `json:"max_length"`
}

type policySpec struct {
	PolicyID string   `json:"policy_id"`
	Kind     string   `json:"kind"`
	Rules    []string `json:"rules"`
}

type enumSpec struct {
	Name   string      `json:"name"`
	Values []enumValue `json:"values"`
}

type enumValue struct {
	Value string `json:"value"`
}
