package main

func verifyRepository(root string, m manifest, checkDoc bool) error {
	if err := verifyRequiredDocs(root, m); err != nil {
		return err
	}
	if checkDoc {
		if err := verifyGeneratedDocs(root, m); err != nil {
			return err
		}
	}
	if err := verifyPackageCoverage(root, m); err != nil {
		return err
	}
	return verifyNoStaleRuntimeWords(root, m)
}

func verifyGeneratedDocs(root string, m manifest) error {
	checks := map[string]string{
		m.GeneratedDocs.ModuleDecomposition: renderModuleDoc(m),
		m.GeneratedDocs.IntegrationMatrix:   renderIntegrationDoc(m),
	}
	for path, expected := range checks {
		current, err := readRepoFile(root, path)
		if err != nil {
			return err
		}
		if current != expected {
			return errOutOfDate(path)
		}
	}
	return nil
}
