package main

func build(root string, opt options) (model, string, error) {
	m, err := readJSONFile[manifest](resolve(root, opt.manifest))
	if err != nil {
		return model{}, "", err
	}
	if err := verifyManifest(m); err != nil {
		return model{}, "", err
	}
	model, err := buildModel(root, m)
	if err != nil {
		return model, "", err
	}
	if err := verifyModel(model); err != nil {
		return model, "", err
	}
	return model, renderDoc(model), nil
}
