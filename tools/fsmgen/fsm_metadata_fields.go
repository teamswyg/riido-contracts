package main

import "strings"

func applyFSMMetadataField(spec *fsmMetadata, list []node, index int) (int, error) {
	item := list[index]
	if index+1 >= len(list) {
		return 0, missingTransitionPropertyError(item)
	}
	key := strings.TrimPrefix(item.atom, ":")
	value := list[index+1]
	switch key {
	case "package":
		spec.Package = atom(value)
	case "name":
		spec.TransitionName = atom(value)
	case "from-enum":
		spec.FromEnum = atom(value)
	case "to-enum":
		spec.ToEnum = atom(value)
	case "event-enum":
		spec.EventEnum = atom(value)
	case "allow-same":
		spec.AllowSame = atom(value) == "true"
	case "fsm-name":
		spec.FSMName = atom(value)
	case "fsm-type-union":
		spec.TypeUnion = atom(value)
	case "pattern-source":
		spec.PatternSource = atom(value)
	case "conformance-profile":
		spec.ConformanceProfile = atom(value)
	case "patterns":
		spec.Patterns = atomList(value)
	case "start-points":
		spec.StartPoints = atomList(value)
	case "end-points":
		spec.EndPoints = atomList(value)
	case "readme-section":
		spec.ReadmeSection = atom(value)
	}
	return index + 1, nil
}
