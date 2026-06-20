package main

import "fmt"

func verifyCounts(model model) error {
	m := model.Manifest
	checks := map[string][2]int{
		"operation count":    {len(model.Operations), m.ExpectedOperationCount},
		"v1 operation count": {model.V1Count, m.ExpectedV1Operations},
		"v2 operation count": {model.V2Count, m.ExpectedV2Operations},
		"v2-only count":      {len(model.V2Only), m.ExpectedV2OnlyOperations},
		"openapi paths":      {model.OpenAPIPathCount, m.ExpectedOpenAPIPaths},
		"openapi operations": {model.OpenAPIOpCount, m.ExpectedOpenAPIOperations},
		"stream variants":    {len(model.StreamVariants), m.ExpectedStreamVariantCount},
	}
	for name, values := range checks {
		if values[0] != values[1] {
			return fmt.Errorf("%s = %d, want %d", name, values[0], values[1])
		}
	}
	return nil
}
