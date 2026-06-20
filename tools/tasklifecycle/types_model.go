package main

type model struct {
	Manifest        manifest
	FSMSchema       int
	States          []stateRow
	StartStates     []string
	TerminalStates  []string
	TransitionCount int
	Transitions     []transitionGroup
}

type stateRow struct {
	Name string
	Kind string
}

type transitionGroup struct {
	From  string
	Edges []transitionEdge
}

type transitionEdge struct {
	Trigger string
	To      []string
}
