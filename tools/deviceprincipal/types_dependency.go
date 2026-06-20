package main

type dependencyMap struct {
	Facts []dependencyFact `json:"facts"`
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
