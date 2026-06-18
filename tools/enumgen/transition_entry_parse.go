package main

import "errors"

func parseTransitionEntry(form node) (transitionEntry, error) {
	if len(form.list) != 3 && len(form.list) != 4 {
		return transitionEntry{}, errors.New("transition requires from, to, and optional event")
	}
	entry := transitionEntry{From: atom(form.list[1]), To: atom(form.list[2])}
	if len(form.list) == 4 {
		entry.Event = atom(form.list[3])
	}
	return entry, nil
}
