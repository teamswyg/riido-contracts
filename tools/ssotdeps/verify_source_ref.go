package main

import (
	"errors"
	"fmt"
	"strings"
)

func verifySourceRef(root string, ref sourceRef) error {
	if strings.TrimSpace(ref.Repo) == "" {
		return errors.New("source_ref.repo is required")
	}
	if strings.TrimSpace(ref.Path) == "" {
		return errors.New("source_ref.path is required")
	}
	if strings.TrimSpace(ref.RequiredPhrase) == "" {
		return fmt.Errorf("source_ref %s:%s required_phrase is required", ref.Repo, ref.Path)
	}
	if ref.Repo != localRepo {
		return nil
	}
	body, err := readLocalRef(root, ref.Path)
	if err != nil {
		return fmt.Errorf("source_ref %s: %w", ref.Path, err)
	}
	if !strings.Contains(body, ref.RequiredPhrase) {
		return fmt.Errorf("source_ref %s does not contain phrase %q", ref.Path, ref.RequiredPhrase)
	}
	return nil
}
