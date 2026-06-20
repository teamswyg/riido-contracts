package main

import "path/filepath"

func loadCoverageEntryIncludes(base string, m *manifest) error {
	for _, file := range m.CoverageEntryFiles {
		entry, err := loadCoverageEntryDocument(filepath.Join(base, file))
		if err != nil {
			return err
		}
		m.Entries = append(m.Entries, entry)
	}
	for _, file := range m.NonUICoverageEntryFiles {
		entry, err := loadCoverageEntryDocument(filepath.Join(base, file))
		if err != nil {
			return err
		}
		m.NonUITopLevelNodes = append(m.NonUITopLevelNodes, entry)
	}
	return nil
}

func loadEvidenceNodeIncludes(base string, m *manifest) error {
	for _, file := range m.VerifiedEvidenceNodeFiles {
		node, err := loadNodeDocument(filepath.Join(base, file))
		if err != nil {
			return err
		}
		m.VerifiedEvidenceNodes = append(m.VerifiedEvidenceNodes, node)
	}
	return nil
}
