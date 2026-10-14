package main

import (
	"fmt"
	"sort"
)

type visitState int

const (
	visitUnseen visitState = iota
	visitActive
	visitDone
)

func verifyAcyclic(graph map[string][]string) error {
	state := map[string]visitState{}
	var visit func(string) error
	visit = func(node string) error {
		switch state[node] {
		case visitUnseen:
		case visitActive:
			return fmt.Errorf("repo dependency cycle detected at %s", node)
		case visitDone:
			return nil
		}
		state[node] = visitActive
		next := append([]string(nil), graph[node]...)
		sort.Strings(next)
		for _, child := range next {
			if err := visit(child); err != nil {
				return err
			}
		}
		state[node] = visitDone
		return nil
	}
	return visitGraphNodes(graph, state, visit)
}

func visitGraphNodes(graph map[string][]string, state map[string]visitState, visit func(string) error) error {
	nodes := make([]string, 0, len(graph))
	for node := range graph {
		nodes = append(nodes, node)
	}
	sort.Strings(nodes)
	for _, node := range nodes {
		if state[node] == visitUnseen {
			if err := visit(node); err != nil {
				return err
			}
		}
	}
	return nil
}
