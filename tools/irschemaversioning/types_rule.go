package main

type scopeRule struct {
	Scope       string `json:"scope"`
	Required    int    `json:"required"`
	Forbidden   int    `json:"forbidden"`
	Conditional int    `json:"conditional"`
}
