package main

func verifyRepository(root string, m manifest, checkDoc bool) error {
	if checkDoc {
		if err := verifyGeneratedDoc(root, m); err != nil {
			return err
		}
	}
	if err := verifyWorkflow(root, m); err != nil {
		return err
	}
	return verifyEntries(m)
}

func verifyGeneratedDoc(root string, m manifest) error {
	current, err := readRepoFile(root, m.GeneratedDoc)
	if err != nil {
		return err
	}
	if current != renderManifest(m) {
		return errOutOfDate(m.GeneratedDoc)
	}
	return nil
}
