package main

import "fmt"

func verifyConformance(metadata map[string]fsmMetadata, patternDocs map[string]patternDocument) error {
	for _, spec := range metadata {
		source, err := cleanRepoRelativePath(spec.PatternSource)
		if err != nil {
			return fmt.Errorf("transitions %s pattern-source: %w", spec.TransitionName, err)
		}
		patterns, ok := patternDocs[source]
		if !ok {
			return fmt.Errorf("transitions %s imports unloaded pattern-source %s", spec.TransitionName, source)
		}
		if err := verifyFSMConformance(spec, patterns); err != nil {
			return err
		}
	}
	return nil
}

func profileCount(patternDocs map[string]patternDocument) int {
	count := 0
	for _, doc := range patternDocs {
		count += len(doc.Profiles)
	}
	return count
}

func verifyFSMConformance(spec fsmMetadata, patterns patternDocument) error {
	profile, ok := patterns.Profiles[spec.ConformanceProfile]
	if !ok {
		return fmt.Errorf("transitions %s references unknown conformance-profile %s", spec.TransitionName, spec.ConformanceProfile)
	}
	if err := verifyPatternNames(spec, patterns.SumType, profile); err != nil {
		return err
	}
	if err := verifyPatternConsistency(spec); err != nil {
		return err
	}
	return verifyGraphConformance(spec)
}
