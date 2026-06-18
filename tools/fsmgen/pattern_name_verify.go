package main

import "fmt"

func verifyPatternNames(spec fsmMetadata, sumType patternSumType, profile conformanceProfile) error {
	known := map[string]bool{}
	for _, value := range sumType.Values {
		known[value.Const] = true
	}
	allowed := stringSet(profile.Allowed)
	rejected := stringSet(profile.Rejected)
	for _, pattern := range spec.Patterns {
		if !known[pattern] {
			return fmt.Errorf("transitions %s references unknown pattern %s", spec.TransitionName, pattern)
		}
		if rejected[pattern] {
			return fmt.Errorf("transitions %s uses rejected pattern %s", spec.TransitionName, pattern)
		}
		if !allowed[pattern] {
			return fmt.Errorf("transitions %s uses pattern %s outside profile %s", spec.TransitionName, pattern, profile.Name)
		}
	}
	for _, pattern := range profile.Rejected {
		if !known[pattern] {
			return fmt.Errorf("profile %s rejects unknown pattern %s", profile.Name, pattern)
		}
	}
	return nil
}
