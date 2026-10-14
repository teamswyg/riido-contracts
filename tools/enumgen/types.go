package main

type document struct {
	Enums       []enumSpec
	Transitions []transitionSpec
}

type enumSpec struct {
	Package     string
	Type        string
	CodeType    string
	StringType  string
	ConstPrefix string
	AllFunc     string
	CodeAllFunc string
	Values      []enumValue
}

type enumValue struct {
	Const string
	Value string
	Attrs map[string]string
}

type transitionSpec struct {
	Package   string
	Name      string
	FromEnum  string
	ToEnum    string
	EventEnum string
	AllFunc   string
	Validate  string
	AllowSame bool
	Entries   []transitionEntry
}

type transitionEntry struct {
	From  string
	To    string
	Event string
}
