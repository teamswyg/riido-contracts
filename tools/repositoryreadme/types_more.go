package main

type fsmDoc struct {
	Intro    []string     `json:"intro"`
	Sections []fsmSection `json:"sections"`
}

type fsmSection struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"-"`
}

type evidenceLoop struct {
	Observation   string `json:"observation"`
	Hypothesis    string `json:"hypothesis"`
	Execute       string `json:"execute"`
	Evaluate      string `json:"evaluate"`
	Retrospective string `json:"retrospective"`
}
