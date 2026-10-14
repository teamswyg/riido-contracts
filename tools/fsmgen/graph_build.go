package main

import "fmt"

func fsmGraphFromSpec(spec fsmMetadata) (fsmGraph, error) {
	graph := newFSMGraph()
	for _, entry := range spec.Entries {
		if err := validateGraphEntry(spec, entry); err != nil {
			return fsmGraph{}, err
		}
		key := graphEntryKey(entry)
		if graph.Seen[key] {
			return fsmGraph{}, fmt.Errorf("transitions %s has duplicate transition %s -> %s", spec.TransitionName, entry.From, entry.To)
		}
		graph.Seen[key] = true
		graph.Vertices[entry.From] = true
		graph.Vertices[entry.To] = true
		graph.Outgoing[entry.From] = append(graph.Outgoing[entry.From], entry.To)
	}
	return graph, nil
}

func validateGraphEntry(spec fsmMetadata, entry fsmTransitionEntry) error {
	if entry.From == "" || entry.To == "" {
		return fmt.Errorf("transitions %s has blank from/to", spec.TransitionName)
	}
	if spec.EventEnum != "" && entry.Event == "" {
		return fmt.Errorf("transitions %s has blank event in event-driven transition", spec.TransitionName)
	}
	if spec.EventEnum == "" && entry.Event != "" {
		return fmt.Errorf("transitions %s has event without event-enum", spec.TransitionName)
	}
	return nil
}

func graphEntryKey(entry fsmTransitionEntry) string {
	return entry.From + "\x00" + entry.To + "\x00" + entry.Event
}
