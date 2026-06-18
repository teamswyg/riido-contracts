package main

import "bytes"

func enumHasPredicateSection(enum enumSpec) bool {
	return hasAttr(enum, "terminal") ||
		hasAttr(enum, "active") ||
		hasAttr(enum, "agent-active") ||
		hasAttr(enum, "transition") ||
		enumHasNativeConfigRequirement(enum)
}

func writeEnumPredicates(b *bytes.Buffer, enum enumSpec) {
	writePredicate(b, enum, "terminal", "IsTerminal")
	writePredicate(b, enum, "active", "IsActive")
	writePredicate(b, enum, "agent-active", "IsAgentActive")
	writePredicate(b, enum, "transition", "IsTransition")
	writeNativeConfigRequirement(b, enum)
	writePackagePredicate(b, enum, "terminal", "IsTerminal")
	writePackagePredicate(b, enum, "agent-active", "IsAgentActive")
}

func hasAttr(enum enumSpec, attr string) bool {
	return len(enum.valuesWithAttr(attr, "true")) > 0
}

func enumHasNativeConfigRequirement(enum enumSpec) bool {
	return enum.Package == "ir" && enum.Type == "EventType"
}
