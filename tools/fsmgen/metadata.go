package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func loadFSMMetadata(path string) (map[string]fsmMetadata, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}
	root, err := parseSExpr(string(body))
	if err != nil {
		return nil, err
	}
	if root.isAtom() || len(root.list) == 0 || atom(root.list[0]) != "enum-gen" {
		return nil, errors.New("root form must be (enum-gen ...)")
	}
	metadata := map[string]fsmMetadata{}
	for _, form := range root.list[1:] {
		if form.isAtom() || len(form.list) == 0 || atom(form.list[0]) != "transitions" {
			continue
		}
		spec, err := parseFSMMetadata(form)
		if err != nil {
			return nil, err
		}
		key := fsmMetadataKey(spec.Package, spec.TransitionName)
		if _, ok := metadata[key]; ok {
			return nil, fmt.Errorf("duplicate fsm metadata for %s", key)
		}
		metadata[key] = spec
	}
	return metadata, nil
}

func parseFSMMetadata(form node) (fsmMetadata, error) {
	var spec fsmMetadata
	for i := 1; i < len(form.list); i++ {
		item := form.list[i]
		if item.isAtom() && strings.HasPrefix(item.atom, ":") {
			if i+1 >= len(form.list) {
				return fsmMetadata{}, fmt.Errorf("transition property %s missing value", item.atom)
			}
			key := strings.TrimPrefix(item.atom, ":")
			value := form.list[i+1]
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
			i++
			continue
		}
		if item.isAtom() || len(item.list) == 0 || atom(item.list[0]) != "transition" {
			continue
		}
		if len(item.list) != 3 && len(item.list) != 4 {
			return fsmMetadata{}, fmt.Errorf("transitions %s has invalid transition entry", spec.TransitionName)
		}
		entry := fsmTransitionEntry{
			From: atom(item.list[1]),
			To:   atom(item.list[2]),
		}
		if len(item.list) == 4 {
			entry.Event = atom(item.list[3])
		}
		spec.Entries = append(spec.Entries, entry)
	}
	if spec.Package == "" || spec.TransitionName == "" {
		return fsmMetadata{}, errors.New("transitions block missing package or name")
	}
	if spec.FromEnum == "" || spec.ToEnum == "" {
		return fsmMetadata{}, fmt.Errorf("transitions %s missing from-enum or to-enum", spec.TransitionName)
	}
	if spec.FSMName == "" || spec.TypeUnion == "" || spec.PatternSource == "" || spec.ConformanceProfile == "" || spec.ReadmeSection == "" {
		return fsmMetadata{}, fmt.Errorf("transitions %s missing fsm-name, fsm-type-union, pattern-source, conformance-profile, or readme-section", spec.TransitionName)
	}
	if len(spec.Patterns) == 0 || len(spec.StartPoints) == 0 || len(spec.EndPoints) == 0 {
		return fsmMetadata{}, fmt.Errorf("transitions %s missing patterns, start-points, or end-points", spec.TransitionName)
	}
	if len(spec.Entries) == 0 {
		return fsmMetadata{}, fmt.Errorf("transitions %s has no transition entries", spec.TransitionName)
	}
	return spec, nil
}

func fsmMetadataKey(packageName, transitionName string) string {
	return packageName + "." + transitionName
}

func requireFSMMetadata(metadata map[string]fsmMetadata, packageName, transitionName string) (fsmMetadata, error) {
	key := fsmMetadataKey(packageName, transitionName)
	spec, ok := metadata[key]
	if !ok {
		return fsmMetadata{}, fmt.Errorf("missing fsm metadata for %s", key)
	}
	return spec, nil
}
