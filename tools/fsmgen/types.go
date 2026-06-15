package main

type generatedArtifact struct {
	Path string
	Body []byte
}

type patternDocument struct {
	SumType  patternSumType
	Profiles map[string]conformanceProfile
}

type patternSumType struct {
	Package     string
	Type        string
	CodeType    string
	StringType  string
	ConstPrefix string
	OutputPath  string
	Values      []patternValue
}

type patternValue struct {
	Const string
	Value string
}

type conformanceProfile struct {
	Name     string
	Allowed  []string
	Rejected []string
}

type readmeSection struct {
	ID      string
	Content string
}

type fsmMetadata struct {
	Package            string
	TransitionName     string
	FromEnum           string
	ToEnum             string
	EventEnum          string
	AllowSame          bool
	FSMName            string
	TypeUnion          string
	PatternSource      string
	ConformanceProfile string
	Patterns           []string
	StartPoints        []string
	EndPoints          []string
	ReadmeSection      string
	Entries            []fsmTransitionEntry
}

type fsmTransitionEntry struct {
	From  string
	To    string
	Event string
}
