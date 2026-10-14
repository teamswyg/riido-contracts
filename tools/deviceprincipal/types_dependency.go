package main

type dependencyMap struct {
	FactFiles []string         `json:"fact_files"`
	Facts     []dependencyFact `json:"facts"`
}

type dependencyFactDocument struct {
	Fact dependencyFact `json:"fact"`
}

type dependencyFact struct {
	ID        string      `json:"id"`
	SourceRef []sourceRef `json:"source_refs"`
}

type sourceRef struct {
	Repo           string `json:"repo"`
	Path           string `json:"path"`
	RequiredPhrase string `json:"required_phrase"`
}
