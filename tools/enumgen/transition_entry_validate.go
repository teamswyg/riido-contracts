package main

import "fmt"

func validateTransitionEntryRefs(
	transitions transitionSpec,
	from enumSpec,
	to enumSpec,
	event enumSpec,
	entry transitionEntry,
) error {
	if !from.hasConst(entry.From) {
		return fmt.Errorf("transitions %s unknown from const %s", transitions.Name, entry.From)
	}
	if !to.hasConst(entry.To) {
		return fmt.Errorf("transitions %s unknown to const %s", transitions.Name, entry.To)
	}
	if transitions.EventEnum != "" && !event.hasConst(entry.Event) {
		return fmt.Errorf("transitions %s unknown event const %s", transitions.Name, entry.Event)
	}
	if transitions.EventEnum == "" && entry.Event != "" {
		return fmt.Errorf("transitions %s does not declare event enum", transitions.Name)
	}
	return nil
}
