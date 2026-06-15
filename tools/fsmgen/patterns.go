package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func loadPatternDocument(path string) (patternDocument, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return patternDocument{}, fmt.Errorf("read %s: %w", path, err)
	}
	root, err := parseSExpr(string(body))
	if err != nil {
		return patternDocument{}, err
	}
	if root.isAtom() || len(root.list) == 0 || atom(root.list[0]) != "fsm-pattern-gen" {
		return patternDocument{}, errors.New("pattern root form must be (fsm-pattern-gen ...)")
	}
	doc := patternDocument{Profiles: map[string]conformanceProfile{}}
	for _, form := range root.list[1:] {
		if form.isAtom() || len(form.list) == 0 {
			return patternDocument{}, errors.New("fsm-pattern-gen children must be lists")
		}
		switch atom(form.list[0]) {
		case "sum-type":
			sumType, err := parsePatternSumType(form)
			if err != nil {
				return patternDocument{}, err
			}
			doc.SumType = sumType
		case "conformance-profile":
			profile, err := parseConformanceProfile(form)
			if err != nil {
				return patternDocument{}, err
			}
			if _, ok := doc.Profiles[profile.Name]; ok {
				return patternDocument{}, fmt.Errorf("duplicate conformance profile %s", profile.Name)
			}
			doc.Profiles[profile.Name] = profile
		default:
			return patternDocument{}, fmt.Errorf("unknown fsm-pattern-gen form %q", atom(form.list[0]))
		}
	}
	if doc.SumType.Type == "" || len(doc.SumType.Values) == 0 {
		return patternDocument{}, errors.New("pattern sum-type is required")
	}
	if len(doc.Profiles) == 0 {
		return patternDocument{}, errors.New("at least one conformance-profile is required")
	}
	return doc, nil
}

func loadPatternDocuments(root string, metadata map[string]fsmMetadata) (map[string]patternDocument, error) {
	sources := map[string]bool{}
	for _, spec := range metadata {
		source, err := cleanRepoRelativePath(spec.PatternSource)
		if err != nil {
			return nil, fmt.Errorf("transitions %s pattern-source: %w", spec.TransitionName, err)
		}
		sources[source] = true
	}
	docs := make(map[string]patternDocument, len(sources))
	for source := range sources {
		doc, err := loadPatternDocument(filepath.Join(root, filepath.FromSlash(source)))
		if err != nil {
			return nil, err
		}
		docs[source] = doc
	}
	return docs, nil
}

func cleanRepoRelativePath(value string) (string, error) {
	normalized := filepath.ToSlash(value)
	if normalized == "" {
		return "", errors.New("path is empty")
	}
	if strings.ContainsRune(normalized, 0) {
		return "", errors.New("path contains NUL")
	}
	if strings.HasPrefix(normalized, "/") {
		return "", fmt.Errorf("path %s must be repository-relative", value)
	}
	clean := path.Clean(normalized)
	if clean == "." || clean == ".." || strings.HasPrefix(clean, "../") {
		return "", fmt.Errorf("path %s must stay inside repository", value)
	}
	return clean, nil
}

func parsePatternSumType(form node) (patternSumType, error) {
	props := map[string]string{}
	values := []patternValue{}
	for i := 1; i < len(form.list); i++ {
		item := form.list[i]
		if item.isAtom() && strings.HasPrefix(item.atom, ":") {
			if i+1 >= len(form.list) {
				return patternSumType{}, fmt.Errorf("sum-type property %s missing value", item.atom)
			}
			props[strings.TrimPrefix(item.atom, ":")] = atom(form.list[i+1])
			i++
			continue
		}
		if item.isAtom() || len(item.list) == 0 || atom(item.list[0]) != "variant" {
			return patternSumType{}, errors.New("sum-type entries must be (variant ...)")
		}
		if len(item.list) != 3 {
			return patternSumType{}, errors.New("variant requires const and string")
		}
		values = append(values, patternValue{Const: atom(item.list[1]), Value: atom(item.list[2])})
	}
	sumType := patternSumType{
		Package:     props["package"],
		Type:        props["type"],
		CodeType:    props["code-type"],
		StringType:  props["string-type"],
		ConstPrefix: props["const-prefix"],
		OutputPath:  props["output"],
		Values:      values,
	}
	if sumType.Package == "" || sumType.Type == "" || sumType.CodeType == "" || sumType.StringType == "" || sumType.OutputPath == "" {
		return patternSumType{}, errors.New("sum-type missing package, type, code-type, string-type, or output")
	}
	if len(sumType.Values) == 0 {
		return patternSumType{}, fmt.Errorf("sum-type %s has no variants", sumType.Type)
	}
	outputPath, err := cleanRepoRelativePath(sumType.OutputPath)
	if err != nil {
		return patternSumType{}, fmt.Errorf("sum-type %s output: %w", sumType.Type, err)
	}
	sumType.OutputPath = outputPath
	return sumType, nil
}

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
		value := form.list[i+1]
		switch key {
		case "name":
			profile.Name = atom(value)
		case "allowed-patterns":
			profile.Allowed = atomList(value)
		case "rejected-patterns":
			profile.Rejected = atomList(value)
		}
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
	if err := verifyGraphConformance(spec); err != nil {
		return err
	}
	return nil
}

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

func verifyPatternConsistency(spec fsmMetadata) error {
	patterns := stringSet(spec.Patterns)
	if !patterns["PatternFlat"] {
		return fmt.Errorf("transitions %s must declare PatternFlat", spec.TransitionName)
	}
	if !patterns["PatternExplicitBoundary"] {
		return fmt.Errorf("transitions %s must declare PatternExplicitBoundary", spec.TransitionName)
	}
	if spec.EventEnum == "" {
		if !patterns["PatternStateDriven"] {
			return fmt.Errorf("transitions %s has no event-enum and must declare PatternStateDriven", spec.TransitionName)
		}
		if patterns["PatternEventDriven"] {
			return fmt.Errorf("transitions %s cannot declare PatternEventDriven without event-enum", spec.TransitionName)
		}
	} else {
		if !patterns["PatternEventDriven"] {
			return fmt.Errorf("transitions %s has event-enum and must declare PatternEventDriven", spec.TransitionName)
		}
		if patterns["PatternStateDriven"] {
			return fmt.Errorf("transitions %s cannot declare PatternStateDriven with event-enum", spec.TransitionName)
		}
	}
	if len(spec.EndPoints) > 1 && !patterns["PatternMultiTerminal"] {
		return fmt.Errorf("transitions %s has multiple end-points and must declare PatternMultiTerminal", spec.TransitionName)
	}
	if spec.AllowSame && !patterns["PatternSameStateAllowed"] {
		return fmt.Errorf("transitions %s allows same-state transitions and must declare PatternSameStateAllowed", spec.TransitionName)
	}
	if !spec.AllowSame && patterns["PatternSameStateAllowed"] {
		return fmt.Errorf("transitions %s declares PatternSameStateAllowed without allow-same", spec.TransitionName)
	}
	hasMultiTargetEvent := specHasMultiTargetEvent(spec)
	if hasMultiTargetEvent && !patterns["PatternMultiTargetEvent"] {
		return fmt.Errorf("transitions %s has multi-target events and must declare PatternMultiTargetEvent", spec.TransitionName)
	}
	if !hasMultiTargetEvent && patterns["PatternMultiTargetEvent"] {
		return fmt.Errorf("transitions %s declares PatternMultiTargetEvent without multi-target event transitions", spec.TransitionName)
	}
	return nil
}

func verifyGraphConformance(spec fsmMetadata) error {
	vertices := map[string]bool{}
	outgoing := map[string][]string{}
	seenTransitions := map[string]bool{}
	for _, entry := range spec.Entries {
		if entry.From == "" || entry.To == "" {
			return fmt.Errorf("transitions %s has blank from/to", spec.TransitionName)
		}
		if spec.EventEnum != "" && entry.Event == "" {
			return fmt.Errorf("transitions %s has blank event in event-driven transition", spec.TransitionName)
		}
		if spec.EventEnum == "" && entry.Event != "" {
			return fmt.Errorf("transitions %s has event without event-enum", spec.TransitionName)
		}
		key := entry.From + "\x00" + entry.To + "\x00" + entry.Event
		if seenTransitions[key] {
			return fmt.Errorf("transitions %s has duplicate transition %s -> %s", spec.TransitionName, entry.From, entry.To)
		}
		seenTransitions[key] = true
		vertices[entry.From] = true
		vertices[entry.To] = true
		outgoing[entry.From] = append(outgoing[entry.From], entry.To)
	}
	for _, state := range spec.StartPoints {
		if !vertices[state] {
			return fmt.Errorf("transitions %s start-point %s is not part of graph", spec.TransitionName, state)
		}
	}
	for _, state := range spec.EndPoints {
		if !vertices[state] {
			return fmt.Errorf("transitions %s end-point %s is not part of graph", spec.TransitionName, state)
		}
		if len(outgoing[state]) > 0 {
			return fmt.Errorf("transitions %s end-point %s must not have outgoing transitions", spec.TransitionName, state)
		}
	}
	reachable := reachableVertices(spec.StartPoints, outgoing)
	for state := range vertices {
		if !reachable[state] {
			return fmt.Errorf("transitions %s state %s is unreachable from start-points", spec.TransitionName, state)
		}
	}
	endSet := stringSet(spec.EndPoints)
	for state := range vertices {
		if endSet[state] {
			continue
		}
		if len(outgoing[state]) == 0 {
			return fmt.Errorf("transitions %s non-end state %s has no outgoing transitions", spec.TransitionName, state)
		}
	}
	return nil
}

func specHasMultiTargetEvent(spec fsmMetadata) bool {
	if spec.EventEnum == "" {
		return false
	}
	targetsByFromEvent := map[string]map[string]bool{}
	for _, entry := range spec.Entries {
		key := entry.From + "\x00" + entry.Event
		if _, ok := targetsByFromEvent[key]; !ok {
			targetsByFromEvent[key] = map[string]bool{}
		}
		targetsByFromEvent[key][entry.To] = true
		if len(targetsByFromEvent[key]) > 1 {
			return true
		}
	}
	return false
}

func reachableVertices(starts []string, outgoing map[string][]string) map[string]bool {
	reachable := map[string]bool{}
	queue := append([]string(nil), starts...)
	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]
		if reachable[state] {
			continue
		}
		reachable[state] = true
		queue = append(queue, outgoing[state]...)
	}
	return reachable
}

func stringSet(values []string) map[string]bool {
	out := make(map[string]bool, len(values))
	for _, value := range values {
		out[value] = true
	}
	return out
}
