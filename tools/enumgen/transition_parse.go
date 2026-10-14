package main

import (
	"errors"
	"fmt"
	"strings"
)

func parseTransitions(form node) (transitionSpec, error) {
	props := map[string]string{}
	entries := []transitionEntry{}
	for i := 1; i < len(form.list); i++ {
		item := form.list[i]
		if item.isAtom() && strings.HasPrefix(item.atom, ":") {
			if i+1 >= len(form.list) {
				return transitionSpec{}, fmt.Errorf("transition property %s missing value", item.atom)
			}
			props[strings.TrimPrefix(item.atom, ":")] = atom(form.list[i+1])
			i++
			continue
		}
		if item.isAtom() || len(item.list) == 0 || atom(item.list[0]) != "transition" {
			return transitionSpec{}, errors.New("transition entries must be (transition ...)")
		}
		entry, err := parseTransitionEntry(item)
		if err != nil {
			return transitionSpec{}, err
		}
		entries = append(entries, entry)
	}
	return validateTransitionSpec(transitionSpec{
		Package:   props["package"],
		Name:      props["name"],
		FromEnum:  props["from-enum"],
		ToEnum:    props["to-enum"],
		EventEnum: props["event-enum"],
		AllFunc:   props["all"],
		Validate:  props["validate"],
		AllowSame: props["allow-same"] == "true",
		Entries:   entries,
	})
}

func validateTransitionSpec(spec transitionSpec) (transitionSpec, error) {
	if spec.Package == "" || spec.Name == "" || spec.FromEnum == "" ||
		spec.ToEnum == "" || spec.AllFunc == "" || spec.Validate == "" {
		return transitionSpec{}, fmt.Errorf("transitions %q missing required properties", spec.Name)
	}
	if len(spec.Entries) == 0 {
		return transitionSpec{}, fmt.Errorf("transitions %s has no entries", spec.Name)
	}
	return spec, nil
}
