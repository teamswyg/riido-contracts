package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func verifyManifestLoopSources(root string, sources []manifestLoopSource) error {
	for _, source := range sources {
		if err := verifyManifestLoopSource(root, source); err != nil {
			return err
		}
	}
	return nil
}

func verifyManifestLoopSource(root string, source manifestLoopSource) error {
	if !filled(source.ID, source.LoopSource) {
		return errors.New("manifest_loop_sources require id and loop_source")
	}
	if len(source.Paths) == 0 && len(source.PathPrefixes) == 0 {
		return fmt.Errorf("manifest_loop_source %s has no target paths", source.ID)
	}
	if !manifestSourceHasLoop(root, "", source.LoopSource) {
		return fmt.Errorf("manifest_loop_source %s has no loop source", source.ID)
	}
	for _, path := range source.Paths {
		if err := verifyManifestLoopSourcePath(root, path); err != nil {
			return fmt.Errorf("manifest_loop_source %s: %w", source.ID, err)
		}
	}
	for _, prefix := range source.PathPrefixes {
		if !strings.HasSuffix(prefix, "/") {
			return fmt.Errorf("manifest_loop_source %s prefix %s must end with /", source.ID, prefix)
		}
		if err := verifyManifestLoopSourcePath(root, prefix); err != nil {
			return fmt.Errorf("manifest_loop_source %s: %w", source.ID, err)
		}
	}
	return nil
}

func verifyManifestLoopSourcePath(root, path string) error {
	target := resolve(root, path)
	if !fileWithinRoot(root, target) {
		return fmt.Errorf("%s escapes repository root", path)
	}
	if _, err := os.Stat(target); err != nil {
		return err
	}
	return nil
}
