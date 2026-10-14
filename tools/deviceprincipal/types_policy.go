package main

type apiContractDoc struct {
	Policies []apiPolicy `json:"policies"`
}

type apiPolicy struct {
	PolicyID    string   `json:"policy_id"`
	Kind        string   `json:"kind"`
	Description string   `json:"description"`
	Rules       []string `json:"rules"`
}
