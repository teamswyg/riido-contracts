package main

import "github.com/teamswyg/riido-contracts/ir"

func transitionEvents(events []ir.EventType) []ir.EventType {
	var out []ir.EventType
	for _, event := range events {
		if event.IsTransition() {
			out = append(out, event)
		}
	}
	return out
}

func eventNames(events []ir.EventType) []string {
	out := make([]string, len(events))
	for i, event := range events {
		out[i] = string(event)
	}
	return out
}
