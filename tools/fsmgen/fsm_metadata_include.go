package main

import (
	"fmt"
	"path/filepath"
)

func isFSMIncludeForm(form node) bool {
	return !form.isAtom() && len(form.list) > 0 && atom(form.list[0]) == "include"
}

func loadFSMMetadataInclude(
	base string,
	form node,
	seen map[string]bool,
) (map[string]fsmMetadata, error) {
	path, err := fsmIncludePath(base, form)
	if err != nil {
		return nil, err
	}
	return loadFSMMetadataAt(path, filepath.Dir(path), seen)
}

func fsmIncludePath(base string, form node) (string, error) {
	if len(form.list) != 2 {
		return "", fmt.Errorf("include form requires exactly one path")
	}
	if base == "" {
		return "", fmt.Errorf("include %q has no base path", atom(form.list[1]))
	}
	path := filepath.Clean(filepath.Join(base, filepath.FromSlash(atom(form.list[1]))))
	return filepath.Abs(path)
}
