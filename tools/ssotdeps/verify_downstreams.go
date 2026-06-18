package main

import (
	"errors"
	"fmt"
	"strings"
)

func verifyDownstreams(downstreams []downstream) error {
	seen := map[string]bool{}
	for _, downstream := range downstreams {
		if strings.TrimSpace(downstream.Repo) == "" {
			return errors.New("downstream repo is required")
		}
		if strings.TrimSpace(downstream.LocalScope) == "" {
			return fmt.Errorf("downstream %q local_scope is required", downstream.Repo)
		}
		if seen[downstream.Repo] {
			return fmt.Errorf("duplicate downstream repo %q", downstream.Repo)
		}
		seen[downstream.Repo] = true
	}
	return nil
}
