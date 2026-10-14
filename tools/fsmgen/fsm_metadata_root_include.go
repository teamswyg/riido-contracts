package main

func fsmMetadataFromRoot(
	root node,
	base string,
	seen map[string]bool,
) (map[string]fsmMetadata, error) {
	metadata := map[string]fsmMetadata{}
	for _, form := range root.list[1:] {
		if err := addFSMRootForm(metadata, base, seen, form); err != nil {
			return nil, err
		}
	}
	return metadata, nil
}

func addFSMRootForm(metadata map[string]fsmMetadata, base string, seen map[string]bool, form node) error {
	if isFSMIncludeForm(form) {
		included, err := loadFSMMetadataInclude(base, form, seen)
		if err != nil {
			return err
		}
		return mergeFSMMetadata(metadata, included)
	}
	if !isTransitionForm(form) {
		return nil
	}
	return addFSMMetadataForm(metadata, form)
}

func mergeFSMMetadata(target, source map[string]fsmMetadata) error {
	for key, value := range source {
		if _, ok := target[key]; ok {
			return duplicateFSMMetadataError(key)
		}
		target[key] = value
	}
	return nil
}
