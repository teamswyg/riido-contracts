package main

import "fmt"

func verifyPatternConsistency(spec fsmMetadata) error {
	patterns := stringSet(spec.Patterns)
	if err := verifyBasePatternConsistency(spec, patterns); err != nil {
		return err
	}
	if err := verifyDriverPatternConsistency(spec, patterns); err != nil {
		return err
	}
	if err := verifyTerminalPatternConsistency(spec, patterns); err != nil {
		return err
	}
	return verifyMultiTargetPatternConsistency(spec, patterns)
}

func verifyBasePatternConsistency(spec fsmMetadata, patterns map[string]bool) error {
	if !patterns["PatternFlat"] {
		return fmt.Errorf("transitions %s must declare PatternFlat", spec.TransitionName)
	}
	if !patterns["PatternExplicitBoundary"] {
		return fmt.Errorf("transitions %s must declare PatternExplicitBoundary", spec.TransitionName)
	}
	return nil
}

func verifyDriverPatternConsistency(spec fsmMetadata, patterns map[string]bool) error {
	if spec.EventEnum == "" {
		if !patterns["PatternStateDriven"] {
			return fmt.Errorf("transitions %s has no event-enum and must declare PatternStateDriven", spec.TransitionName)
		}
		if patterns["PatternEventDriven"] {
			return fmt.Errorf("transitions %s cannot declare PatternEventDriven without event-enum", spec.TransitionName)
		}
		return nil
	}
	if !patterns["PatternEventDriven"] {
		return fmt.Errorf("transitions %s has event-enum and must declare PatternEventDriven", spec.TransitionName)
	}
	if patterns["PatternStateDriven"] {
		return fmt.Errorf("transitions %s cannot declare PatternStateDriven with event-enum", spec.TransitionName)
	}
	return nil
}
