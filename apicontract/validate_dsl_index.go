package apicontract

type dslValidationIndex struct {
	components map[string]struct{}
	schemas    map[string]struct{}
}

func newDSLValidationIndex() dslValidationIndex {
	return dslValidationIndex{
		components: map[string]struct{}{},
		schemas:    map[string]struct{}{},
	}
}
