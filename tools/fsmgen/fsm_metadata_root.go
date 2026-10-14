package main

func fsmMetadataFromLoadedNode(
	root node,
	base string,
	seen map[string]bool,
) (map[string]fsmMetadata, error) {
	if root.isAtom() || len(root.list) == 0 || atom(root.list[0]) != "enum-gen" {
		return fsmMetadataFromSingleForm(root)
	}
	return fsmMetadataFromRoot(root, base, seen)
}

func fsmMetadataFromSingleForm(root node) (map[string]fsmMetadata, error) {
	metadata := map[string]fsmMetadata{}
	if isTransitionForm(root) {
		if err := addFSMMetadataForm(metadata, root); err != nil {
			return nil, err
		}
	}
	return metadata, nil
}

func isTransitionForm(form node) bool {
	return !form.isAtom() && len(form.list) > 0 && atom(form.list[0]) == "transitions"
}
