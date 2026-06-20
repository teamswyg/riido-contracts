package main

func build(root string, opt options) (model, string, error) {
	m, err := readManifest(resolve(root, opt.manifest))
	if err != nil {
		return model{}, "", err
	}
	if err := verifyManifest(m); err != nil {
		return model{}, "", err
	}
	model := buildModel(m)
	if err := verifyModel(model); err != nil {
		return model, "", err
	}
	if err := verifyWorkflow(root, m); err != nil {
		return model, "", err
	}
	return model, renderDoc(model), nil
}
