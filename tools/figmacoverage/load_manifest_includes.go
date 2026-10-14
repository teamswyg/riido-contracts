package main

import "path/filepath"

func loadManifestIncludes(base string, m *manifest) error {
	if err := loadPolicyIncludes(base, m); err != nil {
		return err
	}
	if err := loadCoverageIncludes(base, m); err != nil {
		return err
	}
	if err := loadEvidenceNodeIncludes(base, m); err != nil {
		return err
	}
	return loadAnnotationIncludes(base, m)
}

func loadPolicyIncludes(base string, m *manifest) error {
	for _, file := range m.ToolLimitationFiles {
		limitation, err := loadToolLimitationDocument(filepath.Join(base, file))
		if err != nil {
			return err
		}
		m.SupportingToolLimitations = append(m.SupportingToolLimitations, limitation)
	}
	return nil
}
