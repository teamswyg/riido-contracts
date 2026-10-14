package main

import (
	"fmt"
	"strings"
)

func parseFSMMetadata(form node) (fsmMetadata, error) {
	var spec fsmMetadata
	for i := 1; i < len(form.list); i++ {
		item := form.list[i]
		if item.isAtom() && strings.HasPrefix(item.atom, ":") {
			next, err := applyFSMMetadataField(&spec, form.list, i)
			if err != nil {
				return fsmMetadata{}, err
			}
			i = next
			continue
		}
		if isFSMTransitionEntry(item) {
			entry, err := parseFSMTransitionEntry(spec, item)
			if err != nil {
				return fsmMetadata{}, err
			}
			spec.Entries = append(spec.Entries, entry)
		}
	}
	return validateFSMMetadata(spec)
}

func isFSMTransitionEntry(form node) bool {
	return !form.isAtom() && len(form.list) > 0 && atom(form.list[0]) == "transition"
}

func missingTransitionPropertyError(item node) error {
	return fmt.Errorf("transition property %s missing value", item.atom)
}
