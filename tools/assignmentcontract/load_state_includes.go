package main

import "path/filepath"

func loadStateIncludes(base string, c *contract) error {
	for _, file := range c.AssignmentStateFiles {
		var doc stateDocument
		path := filepath.Join(base, file)
		if err := loadInclude(path, "assignment state", stateSchema, &doc); err != nil {
			return err
		}
		c.AssignmentStates = append(c.AssignmentStates, doc.State)
	}
	return nil
}
