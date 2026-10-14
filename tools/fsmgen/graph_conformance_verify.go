package main

func verifyGraphConformance(spec fsmMetadata) error {
	graph, err := fsmGraphFromSpec(spec)
	if err != nil {
		return err
	}
	if err := verifyGraphBoundaries(spec, graph); err != nil {
		return err
	}
	if err := verifyGraphReachability(spec, graph); err != nil {
		return err
	}
	return verifyGraphOutgoing(spec, graph)
}

type fsmGraph struct {
	Vertices map[string]bool
	Outgoing map[string][]string
	Seen     map[string]bool
}

func newFSMGraph() fsmGraph {
	return fsmGraph{
		Vertices: map[string]bool{},
		Outgoing: map[string][]string{},
		Seen:     map[string]bool{},
	}
}
