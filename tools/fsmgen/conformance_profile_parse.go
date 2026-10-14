package main

import (
	"errors"
	"fmt"
	"strings"
)

func parseConformanceProfile(form node) (conformanceProfile, error) {
	var profile conformanceProfile
	for i := 1; i < len(form.list); i++ {
		item := form.list[i]
		if !item.isAtom() || !strings.HasPrefix(item.atom, ":") {
			continue
		}
		if i+1 >= len(form.list) {
			return conformanceProfile{}, fmt.Errorf("conformance-profile property %s missing value", item.atom)
		}
		key := strings.TrimPrefix(item.atom, ":")
		applyConformanceProfileField(&profile, key, form.list[i+1])
		i++
	}
	if profile.Name == "" {
		return conformanceProfile{}, errors.New("conformance-profile name is required")
	}
	if len(profile.Allowed) == 0 {
		return conformanceProfile{}, fmt.Errorf("conformance-profile %s has no allowed-patterns", profile.Name)
	}
	return profile, nil
}

func applyConformanceProfileField(profile *conformanceProfile, key string, value node) {
	switch key {
	case "name":
		profile.Name = atom(value)
	case "allowed-patterns":
		profile.Allowed = atomList(value)
	case "rejected-patterns":
		profile.Rejected = atomList(value)
	}
}
