package main

func buildModel(root string, m manifest) (model, error) {
	dsl, err := readLooseJSONFile[contractFixture](resolve(root, m.DSLFixture))
	if err != nil {
		return model{}, err
	}
	ir, err := readLooseJSONFile[contractFixture](resolve(root, m.IRFixture))
	if err != nil {
		return model{}, err
	}
	openapi, err := readLooseJSONFile[openAPIDoc](resolve(root, m.OpenAPIFixture))
	if err != nil {
		return model{}, err
	}
	model := model{Manifest: m, DSLIRMatch: true, OpenAPIMatch: true}
	model = addOperations(model, dsl, ir, openapi)
	model = addSchemas(model, dsl)
	model = addPolicies(model, dsl)
	model = addEnums(model, dsl)
	model.StreamPass = hasStreamVariant(dsl.SumTypes, m.RequiredStreamVariant)
	return model, nil
}
