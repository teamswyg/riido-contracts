package main

import (
	"errors"
	"fmt"
	"strings"
)

func verifyOwner(root string, owner ownerRef) error {
	if strings.TrimSpace(owner.Repo) == "" {
		return errors.New("owner.repo is required")
	}
	if strings.TrimSpace(owner.Path) == "" {
		return errors.New("owner.path is required")
	}
	if owner.Repo == localRepo {
		if _, err := readLocalRef(root, owner.Path); err != nil {
			return fmt.Errorf("owner %s: %w", owner.Path, err)
		}
	}
	return nil
}
