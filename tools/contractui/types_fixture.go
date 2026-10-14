package main

type irDocument struct {
	ContractID string      `json:"contract_id"`
	Context    string      `json:"context"`
	Service    service     `json:"service"`
	Operations []operation `json:"operations"`
}

type dslDocument struct {
	Operations []operation `json:"operations"`
}

type openAPIDocument struct {
	Components openAPIComponents `json:"components"`
}

type openAPIComponents struct {
	Schemas map[string]any `json:"schemas"`
}
