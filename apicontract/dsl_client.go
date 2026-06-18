package apicontract

type ClientModule struct {
	Module      string            `json:"module"`
	Description string            `json:"description,omitempty"`
	Namespaces  []ClientNamespace `json:"namespaces,omitempty"`
}

type ClientNamespace struct {
	Path        []string `json:"path"`
	Description string   `json:"description,omitempty"`
}

type ClientMeta struct {
	Module        string   `json:"module"`
	FacadePath    []string `json:"facade_path"`
	GeneratedPath string   `json:"generated_path,omitempty"`
	CacheTag      string   `json:"cache_tag,omitempty"`
	Invalidates   []string `json:"invalidates,omitempty"`
}
