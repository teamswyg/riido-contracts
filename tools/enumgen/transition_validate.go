package main

import "fmt"

func validateTransitionRefs(transitions transitionSpec, enums map[string]enumSpec) error {
	from, ok := enums[transitions.FromEnum]
	if !ok {
		return fmt.Errorf("transitions %s unknown from enum %s", transitions.Name, transitions.FromEnum)
	}
	to, ok := enums[transitions.ToEnum]
	if !ok {
		return fmt.Errorf("transitions %s unknown to enum %s", transitions.Name, transitions.ToEnum)
	}
	event, err := transitionEventEnum(transitions, enums)
	if err != nil {
		return err
	}
	for _, entry := range transitions.Entries {
		if err := validateTransitionEntryRefs(transitions, from, to, event, entry); err != nil {
			return err
		}
	}
	return nil
}

func transitionEventEnum(transitions transitionSpec, enums map[string]enumSpec) (enumSpec, error) {
	if transitions.EventEnum == "" {
		return enumSpec{}, nil
	}
	event, ok := enums[transitions.EventEnum]
	if !ok {
		return enumSpec{}, fmt.Errorf("transitions %s unknown event enum %s", transitions.Name, transitions.EventEnum)
	}
	return event, nil
}
