package main

import "path/filepath"

func loadCoverageIncludes(base string, m *manifest) error {
	if err := loadExpectedNodeIncludes(base, m); err != nil {
		return err
	}
	if err := loadPageInventoryIncludes(base, m); err != nil {
		return err
	}
	return loadCoverageEntryIncludes(base, m)
}

func loadExpectedNodeIncludes(base string, m *manifest) error {
	for _, file := range m.ExpectedTopLevelNodeFiles {
		node, err := loadNodeDocument(filepath.Join(base, file))
		if err != nil {
			return err
		}
		m.ExpectedTopLevelNodes = append(m.ExpectedTopLevelNodes, node)
	}
	return nil
}

func loadPageInventoryIncludes(base string, m *manifest) error {
	for _, file := range m.PageInventoryFiles {
		inventory, err := loadPageInventoryDocument(filepath.Join(base, file))
		if err != nil {
			return err
		}
		m.NonUITopLevelInventory = append(m.NonUITopLevelInventory, inventory)
	}
	return nil
}
