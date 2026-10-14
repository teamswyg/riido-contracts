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
	model := model{Manifest: m}
	model.FixtureSchema = findSchema(dsl.Schemas, "AgentOnboardingFixture")
	model.ListSchema = findSchema(dsl.Schemas, "AgentOnboardingFixtureListResponse")
	model.CreateRequestSchema = findSchema(dsl.Schemas, "CreateAgentConfigurationRequest")
	model.FixtureFields = propertyNames(model.FixtureSchema)
	model.CreateRequestFields = propertyNames(model.CreateRequestSchema)
	model.NoDiffPathsClean = noDiffPathsClean(openapi, m.NoDiffPathFragments)
	return addOperations(model, dsl, ir, openapi), nil
}
