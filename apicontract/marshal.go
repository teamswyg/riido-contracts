package apicontract

import (
	"encoding/json"
)

func MarshalCanonical(value any) ([]byte, error) {
	out, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return nil, err
	}
	return append(out, '\n'), nil
}
