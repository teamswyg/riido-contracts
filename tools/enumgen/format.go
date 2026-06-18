package main

import (
	"fmt"
	"go/format"
)

func formatSource(name string, source []byte) ([]byte, error) {
	out, err := format.Source(source)
	if err != nil {
		return nil, fmt.Errorf("format %s: %w\n%s", name, err, source)
	}
	return out, nil
}
