package main

type model struct {
	Manifest          manifest
	Operations        []operationTuple
	V1Count           int
	V2Count           int
	V2Only            []operationTuple
	OpenAPIPathCount  int
	OpenAPIOpCount    int
	StreamVariants    []string
	DSLIRMatch        bool
	IROpenAPIMatch    bool
	V2CoversV1        bool
	StreamVariantPass bool
}
