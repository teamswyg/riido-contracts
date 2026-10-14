package apicontract

func validateIR(ir IRDocument) error {
	if err := validateIRHeader(ir); err != nil {
		return err
	}
	index := buildIRValidationIndex(ir)
	if err := validateIRSumTypeVariants(ir.SumTypes, index.schemas); err != nil {
		return err
	}
	if err := validateIROperationRefs(ir.Operations, index.components); err != nil {
		return err
	}
	return validateIRClientMetadata(ir.ClientModules, ir.Operations)
}
