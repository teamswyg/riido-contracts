package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"os"
)

func readManifest(path string) (manifest, error) {
	var m manifest
	body, err := os.ReadFile(path)
	if err != nil {
		return m, err
	}
	dec := json.NewDecoder(bytes.NewReader(body))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&m); err != nil {
		return m, err
	}
	if err := dec.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return m, errors.New("manifest has trailing JSON")
	}
	return m, nil
}

func writeFile(path string, body []byte) error {
	return os.WriteFile(path, body, 0o644)
}
