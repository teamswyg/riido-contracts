package main

func buildModel(root string, m manifest) (model, error) {
	dsl, err := readLooseJSONFile[apiFixture](resolve(root, m.DSLFixture))
	if err != nil {
		return model{}, err
	}
	ir, err := readLooseJSONFile[apiFixture](resolve(root, m.IRFixture))
	if err != nil {
		return model{}, err
	}
	openapi, err := readLooseJSONFile[openAPIDoc](resolve(root, m.OpenAPIFixture))
	if err != nil {
		return model{}, err
	}
	return buildComparedModel(m, dsl, ir, openapi), nil
}

func buildComparedModel(m manifest, dsl, ir apiFixture, openapi openAPIDoc) model {
	ops := operationTuples(dsl.Operations)
	v1, v2, v2Only := splitSurface(ops)
	openapiOps := openAPITuples(openapi)
	return model{
		Manifest:          m,
		Operations:        ops,
		V1Count:           len(v1),
		V2Count:           len(v2),
		V2Only:            v2Only,
		OpenAPIPathCount:  len(openapi.Paths),
		OpenAPIOpCount:    len(openapiOps),
		StreamVariants:    streamVariants(dsl),
		DSLIRMatch:        sameTuples(ops, operationTuples(ir.Operations)),
		IROpenAPIMatch:    sameTuples(operationTuples(ir.Operations), openapiOps),
		V2CoversV1:        v2CoversV1(v1, v2),
		StreamVariantPass: hasRequiredVariants(streamVariants(dsl), m.RequiredStreamVariants),
	}
}
