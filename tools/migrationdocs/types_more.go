package main

type candidateContract struct {
	Candidate string `json:"candidate"`
	Source    string `json:"source"`
	Decision  string `json:"decision"`
}

type versioning struct {
	Intro string        `json:"intro"`
	Axes  []versionAxis `json:"axes"`
}

type versionAxis struct {
	Axis             string `json:"axis"`
	OwnerBeforeSplit string `json:"owner_before_split"`
	ContractHandling string `json:"contract_handling"`
}

type migrationSlice struct {
	Title   string   `json:"title"`
	Intro   []string `json:"intro"`
	Does    []string `json:"does"`
	DoesNot string   `json:"does_not"`
}

type workMapEntry struct {
	Area             string `json:"area"`
	RiidoTask        string `json:"riido_task"`
	TargetRepository string `json:"target_repository"`
}

type evidenceLoop struct {
	Observation   string `json:"observation"`
	Hypothesis    string `json:"hypothesis"`
	Execute       string `json:"execute"`
	Evaluate      string `json:"evaluate"`
	Retrospective string `json:"retrospective"`
}
