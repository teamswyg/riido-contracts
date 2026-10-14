package main

import "fmt"

func verifyTerminalPatternConsistency(spec fsmMetadata, patterns map[string]bool) error {
	if len(spec.EndPoints) > 1 && !patterns["PatternMultiTerminal"] {
		return fmt.Errorf("transitions %s has multiple end-points and must declare PatternMultiTerminal", spec.TransitionName)
	}
	if spec.AllowSame && !patterns["PatternSameStateAllowed"] {
		return fmt.Errorf("transitions %s allows same-state transitions and must declare PatternSameStateAllowed", spec.TransitionName)
	}
	if !spec.AllowSame && patterns["PatternSameStateAllowed"] {
		return fmt.Errorf("transitions %s declares PatternSameStateAllowed without allow-same", spec.TransitionName)
	}
	return nil
}

func verifyMultiTargetPatternConsistency(spec fsmMetadata, patterns map[string]bool) error {
	hasMultiTargetEvent := specHasMultiTargetEvent(spec)
	if hasMultiTargetEvent && !patterns["PatternMultiTargetEvent"] {
		return fmt.Errorf("transitions %s has multi-target events and must declare PatternMultiTargetEvent", spec.TransitionName)
	}
	if !hasMultiTargetEvent && patterns["PatternMultiTargetEvent"] {
		return fmt.Errorf("transitions %s declares PatternMultiTargetEvent without multi-target event transitions", spec.TransitionName)
	}
	return nil
}
