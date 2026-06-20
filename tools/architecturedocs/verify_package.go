package main

import (
	"fmt"
	"strings"
)

func verifyPackageCoverage(root string, m manifest) error {
	context, err := readRepoFile(root, m.ContextMapDoc)
	if err != nil {
		return err
	}
	module, err := readRepoFile(root, m.GeneratedDocs.ModuleDecomposition)
	if err != nil {
		return err
	}
	haystack := context + "\n" + module
	for _, pkg := range m.Packages {
		if !strings.Contains(haystack, pkg.Path) {
			return fmt.Errorf("package %q missing from architecture docs", pkg.Path)
		}
	}
	return nil
}
