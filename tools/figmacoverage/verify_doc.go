package main

import "fmt"

func verifyRenderedDoc(root string, m manifest) error {
	current, err := readLocalRef(root, m.HumanDoc)
	if err != nil {
		return fmt.Errorf("read generated doc: %w", err)
	}
	expected := renderManifest(m)
	if current != expected {
		return fmt.Errorf("%s is out of date; regenerate with go run ./tools/figmacoverage render > %s", m.HumanDoc, m.HumanDoc)
	}
	return nil
}
