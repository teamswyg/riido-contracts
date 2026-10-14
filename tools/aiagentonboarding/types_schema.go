package main

type schema struct {
	Name       string     `json:"name"`
	Type       string     `json:"type"`
	Required   []string   `json:"required"`
	Properties []property `json:"properties"`
}

type property struct {
	Name  string       `json:"name"`
	Type  string       `json:"type"`
	Ref   string       `json:"ref"`
	Items propertyItem `json:"items"`
}

type propertyItem struct {
	Ref string `json:"ref"`
}
