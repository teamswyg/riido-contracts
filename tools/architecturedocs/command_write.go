package main

func runWriteDoc(args []string) error {
	fs := quietFlagSet("write-doc")
	manifestPath := manifestFlag(fs)
	if err := fs.Parse(args); err != nil {
		return err
	}
	root, m, err := loadDefaultedManifest(*manifestPath)
	if err != nil {
		return err
	}
	if err := writeRepoFile(root, m.GeneratedDocs.ModuleDecomposition, renderModuleDoc(m)); err != nil {
		return err
	}
	return writeRepoFile(root, m.GeneratedDocs.IntegrationMatrix, renderIntegrationDoc(m))
}
