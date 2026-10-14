package apicontract

import "fmt"

func validateIRSumTypeVariants(sumTypes []SumType, schemas map[string]struct{}) error {
	for _, sumType := range sumTypes {
		for _, variant := range sumType.Variants {
			if _, ok := schemas[variant.Schema]; !ok {
				return fmt.Errorf("apicontract: IR sum_type %q variant schema %q is missing", sumType.Name, variant.Schema)
			}
		}
	}
	return nil
}
