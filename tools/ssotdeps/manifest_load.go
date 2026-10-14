package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func loadManifest(path string) (manifest, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return manifest{}, fmt.Errorf("read manifest: %w", err)
	}
	dec := json.NewDecoder(strings.NewReader(string(body)))
	dec.DisallowUnknownFields()
	var m manifest
	if err := dec.Decode(&m); err != nil {
		return manifest{}, fmt.Errorf("decode manifest: %w", err)
	}
	var extra struct{}
	if err := dec.Decode(&extra); err != io.EOF {
		return manifest{}, errors.New("decode manifest: trailing data")
	}
	if err := loadManifestIncludes(filepath.Dir(path), &m); err != nil {
		return manifest{}, err
	}
	return m, nil
}

func loadManifestIncludes(base string, m *manifest) error {
	for _, file := range m.FactFiles {
		fact, err := loadFactDocument(filepath.Join(base, file))
		if err != nil {
			return err
		}
		m.Facts = append(m.Facts, fact)
	}
	for _, file := range m.RepoDependencyFiles {
		dep, err := loadRepoDependencyDocument(filepath.Join(base, file))
		if err != nil {
			return err
		}
		m.RepoDependencies = append(m.RepoDependencies, dep)
	}
	return nil
}
