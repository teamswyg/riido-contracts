package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func loadDocManifest(path string) (docManifest, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return docManifest{}, err
	}
	dec := json.NewDecoder(bytes.NewReader(body))
	dec.DisallowUnknownFields()
	var m docManifest
	if err := dec.Decode(&m); err != nil {
		return docManifest{}, fmt.Errorf("decode doc manifest: %w", err)
	}
	if err := dec.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return docManifest{}, errors.New("decode doc manifest: trailing data")
	}
	return m, nil
}

func writeFile(path string, body []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, body, 0o644)
}
