package main

func build(root string, opt options) (manifest, contract, string, error) {
	m, err := loadJSON[manifest](resolve(root, opt.manifest), "manifest")
	if err != nil {
		return manifest{}, contract{}, "", err
	}
	if err := verifyManifest(m); err != nil {
		return manifest{}, contract{}, "", err
	}
	if err := verifyWorkflow(root, m); err != nil {
		return manifest{}, contract{}, "", err
	}
	c, err := loadContract(resolve(root, m.Contract))
	if err != nil {
		return manifest{}, contract{}, "", err
	}
	return m, c, renderDoc(m, c), nil
}
