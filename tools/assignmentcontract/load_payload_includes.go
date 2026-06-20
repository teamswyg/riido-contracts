package main

import "path/filepath"

func loadPayloadIncludes(base string, c *contract) error {
	for _, file := range c.PayloadFieldFiles {
		var doc payloadFieldDocument
		path := filepath.Join(base, file)
		if err := loadInclude(path, "payload field", payloadSchema, &doc); err != nil {
			return err
		}
		c.AssignmentPayloadFields = append(c.AssignmentPayloadFields, doc.Field)
	}
	return nil
}
