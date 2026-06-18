package progressmessage

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func Catalog() (IRDocument, error) {
	var ir IRDocument
	dec := json.NewDecoder(bytes.NewReader(embeddedIR))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&ir); err != nil {
		return IRDocument{}, fmt.Errorf("progressmessage: decode embedded IR: %w", err)
	}
	if err := ValidateIR(ir); err != nil {
		return IRDocument{}, err
	}
	return ir, nil
}
