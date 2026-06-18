package main

import "fmt"

func addFSMMetadataForm(metadata map[string]fsmMetadata, form node) error {
	spec, err := parseFSMMetadata(form)
	if err != nil {
		return err
	}
	key := fsmMetadataKey(spec.Package, spec.TransitionName)
	if _, ok := metadata[key]; ok {
		return fmt.Errorf("duplicate fsm metadata for %s", key)
	}
	metadata[key] = spec
	return nil
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
