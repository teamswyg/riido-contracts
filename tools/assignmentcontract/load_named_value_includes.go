package main

import "path/filepath"

func loadNamedValueIncludes(base string, c *contract) error {
	if err := loadNamedValues(base, c.PollActionFiles, pollActionSchema, &c.PollActions); err != nil {
		return err
	}
	return loadNamedValues(base, c.TaskEventFiles, taskEventSchema, &c.TaskEvents)
}

func loadNamedValues(base string, files []string, schema string, out *[]namedValue) error {
	for _, file := range files {
		var doc namedValueDocument
		path := filepath.Join(base, file)
		if err := loadInclude(path, "named value", schema, &doc); err != nil {
			return err
		}
		*out = append(*out, doc.Value)
	}
	return nil
}
