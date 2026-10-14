package main

import "fmt"

func parseFSMTransitionEntry(spec fsmMetadata, item node) (fsmTransitionEntry, error) {
	if len(item.list) != 3 && len(item.list) != 4 {
		return fsmTransitionEntry{}, fmt.Errorf("transitions %s has invalid transition entry", spec.TransitionName)
	}
	entry := fsmTransitionEntry{
		From: atom(item.list[1]),
		To:   atom(item.list[2]),
	}
	if len(item.list) == 4 {
		entry.Event = atom(item.list[3])
	}
	return entry, nil
}
