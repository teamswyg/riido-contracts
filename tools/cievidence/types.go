package main

type evidence struct {
	SchemaVersion string          `json:"schema_version"`
	ID            string          `json:"id"`
	Status        string          `json:"status"`
	Workflow      string          `json:"workflow"`
	Commands      []commandRecord `json:"commands"`
	Loop          evidenceLoop    `json:"loop"`
}

type commandRecord struct {
	Command string `json:"command"`
	Found   bool   `json:"found"`
}

type evidenceLoop struct {
	Observation   string `json:"observation"`
	Hypothesis    string `json:"hypothesis"`
	Execute       string `json:"execute"`
	Evaluate      string `json:"evaluate"`
	Retrospective string `json:"retrospective"`
}
