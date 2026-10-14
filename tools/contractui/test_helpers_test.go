package main

func totalOperations(data bundle) int {
	total := 0
	for _, contract := range data.Contracts {
		total += contract.OperationCount
	}
	return total
}

func schemaHasProperty(data bundle, schemaName, propertyName string) bool {
	for _, contract := range data.Contracts {
		schema, ok := contract.Schemas[schemaName].(map[string]any)
		if !ok {
			continue
		}
		props, ok := schema["properties"].(map[string]any)
		if ok && props[propertyName] != nil {
			return true
		}
	}
	return false
}

func scenarioExists(data bundle, name string) bool {
	for _, contract := range data.Contracts {
		for _, operation := range contract.Operations {
			for _, scenario := range operation.Scenarios {
				if scenario.Name == name {
					return true
				}
			}
		}
	}
	return false
}
