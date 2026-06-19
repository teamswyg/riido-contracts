package main

import (
	"fmt"
	"path/filepath"
)

func verifyRenderedDoc(root string, m manifest) error {
	current, err := readLocalRef(root, m.HumanDoc)
	if err != nil {
		return fmt.Errorf("read generated doc: %w", err)
	}
	expected := renderManifest(m)
	if current != expected {
		return fmt.Errorf(
			"%s is out of date; regenerate with go run ./tools/ssotdeps render > %s",
			filepath.ToSlash(m.HumanDoc),
			filepath.ToSlash(m.HumanDoc),
		)
	}
	return nil
}
