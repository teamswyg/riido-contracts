package main

type manifest struct {
	SchemaVersion    string           `json:"schema_version"`
	ID               string           `json:"id"`
	RiidoTask        string           `json:"riido_task"`
	HumanDoc         string           `json:"human_doc"`
	Facts            []fact           `json:"facts"`
	RepoDependencies []repoDependency `json:"repo_dependencies"`
}

type fact struct {
	ID             string       `json:"id"`
	Fact           string       `json:"fact"`
	HumanDocPhrase string       `json:"human_doc_phrase"`
	SourceRefs     []sourceRef  `json:"source_refs"`
	Owner          ownerRef     `json:"owner"`
	Downstreams    []downstream `json:"downstreams"`
}

type sourceRef struct {
	Repo           string `json:"repo"`
	Path           string `json:"path"`
	RequiredPhrase string `json:"required_phrase"`
}

type ownerRef struct {
	Repo string `json:"repo"`
	Path string `json:"path"`
}

type downstream struct {
	Repo       string `json:"repo"`
	LocalScope string `json:"local_scope"`
}

type repoDependency struct {
	ID         string   `json:"id"`
	FromRepo   string   `json:"from_repo"`
	ToRepo     string   `json:"to_repo"`
	FactIDs    []string `json:"fact_ids"`
	LocalScope string   `json:"local_scope"`
}
