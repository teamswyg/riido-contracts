package main

import (
	"errors"
	"fmt"
)

func verifyEntries(expected []node, entries []coverageEntry) error {
	if len(expected) == 0 || len(entries) == 0 {
		return errors.New("expected_top_level_nodes and entries are required")
	}
	if len(expected) != len(entries) {
		return fmt.Errorf("entries=%d expected_top_level_nodes=%d", len(entries), len(expected))
	}
	seen := map[string]bool{}
	for _, expectedNode := range expected {
		seen[expectedNode.NodeID] = false
	}
	for _, entry := range entries {
		if _, ok := seen[entry.NodeID]; !ok {
			return fmt.Errorf("entry %s is not an expected top-level node", entry.NodeID)
		}
		seen[entry.NodeID] = true
		if err := verifyEntry(entry); err != nil {
			return fmt.Errorf("entry %s: %w", entry.NodeID, err)
		}
	}
	for nodeID, found := range seen {
		if !found {
			return fmt.Errorf("expected top-level node %s has no entry", nodeID)
		}
	}
	return nil
}
