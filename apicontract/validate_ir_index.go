package apicontract

type irValidationIndex struct {
	components map[string]struct{}
	schemas    map[string]struct{}
}

func buildIRValidationIndex(ir IRDocument) irValidationIndex {
	index := irValidationIndex{
		components: map[string]struct{}{},
		schemas:    map[string]struct{}{},
	}
	for _, enum := range ir.Enums {
		index.components[enum.Name] = struct{}{}
	}
	for _, sumType := range ir.SumTypes {
		index.components[sumType.Name] = struct{}{}
	}
	for _, component := range ir.Components {
		index.schemas[component.Name] = struct{}{}
		index.components[component.Name] = struct{}{}
	}
	return index
}
