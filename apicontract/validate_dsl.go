package apicontract

func validateDSL(dsl DSLDocument) error {
	if err := validateDSLHeader(dsl); err != nil {
		return err
	}
	clientModules, err := validateClientModules(dsl.ClientModules)
	if err != nil {
		return err
	}
	index := newDSLValidationIndex()
	if err := validateDSLEnums(dsl.Enums, index.components); err != nil {
		return err
	}
	if err := validateDSLSumTypes(dsl.SumTypes, index.components); err != nil {
		return err
	}
	if err := validateDSLSchemas(dsl.Schemas, index); err != nil {
		return err
	}
	if err := validateDSLSumTypeVariantSchemas(dsl.SumTypes, index.schemas); err != nil {
		return err
	}
	cacheTags, err := validateDSLOperations(dsl.Operations, index.components, clientModules)
	if err != nil {
		return err
	}
	return validateDSLOperationInvalidations(dsl.Operations, clientModules, cacheTags)
}
