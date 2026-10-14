package main

import "fmt"

func verifyGraphBoundaries(spec fsmMetadata, graph fsmGraph) error {
	for _, state := range spec.StartPoints {
		if !graph.Vertices[state] {
			return fmt.Errorf("transitions %s start-point %s is not part of graph", spec.TransitionName, state)
		}
	}
	for _, state := range spec.EndPoints {
		if !graph.Vertices[state] {
			return fmt.Errorf("transitions %s end-point %s is not part of graph", spec.TransitionName, state)
		}
		if len(graph.Outgoing[state]) > 0 {
			return fmt.Errorf("transitions %s end-point %s must not have outgoing transitions", spec.TransitionName, state)
		}
	}
	return nil
}

func verifyGraphReachability(spec fsmMetadata, graph fsmGraph) error {
	reachable := reachableVertices(spec.StartPoints, graph.Outgoing)
	for state := range graph.Vertices {
		if !reachable[state] {
			return fmt.Errorf("transitions %s state %s is unreachable from start-points", spec.TransitionName, state)
		}
	}
	return nil
}

func verifyGraphOutgoing(spec fsmMetadata, graph fsmGraph) error {
	endSet := stringSet(spec.EndPoints)
	for state := range graph.Vertices {
		if endSet[state] {
			continue
		}
		if len(graph.Outgoing[state]) == 0 {
			return fmt.Errorf("transitions %s non-end state %s has no outgoing transitions", spec.TransitionName, state)
		}
	}
	return nil
}
