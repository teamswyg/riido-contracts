package main

type schemaExpectation struct {
	Schema string   `json:"schema"`
	Fields []string `json:"fields"`
}

type enumExpectation struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type evidenceLoop struct {
	Observation   string `json:"observation"`
	Hypothesis    string `json:"hypothesis"`
	Execute       string `json:"execute"`
	Evaluate      string `json:"evaluate"`
	Retrospective string `json:"retrospective"`
}
