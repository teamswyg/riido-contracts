package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

func loadManifest(path string) (manifest, error) {
	file, err := os.Open(path)
	if err != nil {
		return manifest{}, err
	}
	defer file.Close()
	dec := json.NewDecoder(file)
	dec.DisallowUnknownFields()
	var m manifest
	if err := dec.Decode(&m); err != nil {
		return manifest{}, err
	}
	var extra any
	err = dec.Decode(&extra)
	if !errors.Is(err, io.EOF) {
		if err == nil {
			err = fmt.Errorf("trailing JSON document")
		}
		return manifest{}, fmt.Errorf("%s has trailing JSON documents: %w", path, err)
	}
	return m, nil
}
