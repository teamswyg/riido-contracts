package main

import "strings"

const schemaVersion = "riido-contracts-ci-evidence.v1"

func buildEvidence(id, workflow, text string) evidence {
	commands := commandRecords(requiredCommands(workflow, id), text)
	return evidence{
		SchemaVersion: schemaVersion,
		ID:            id,
		Status:        status(commands),
		Workflow:      workflow,
		Commands:      commands,
		Loop:          ciLoop(),
	}
}

func commandRecords(required []string, text string) []commandRecord {
	records := make([]commandRecord, 0, len(required))
	for _, command := range required {
		records = append(records, commandRecord{
			Command: command,
			Found:   strings.Contains(text, command),
		})
	}
	return records
}

func status(records []commandRecord) string {
	if len(records) == 0 {
		return "failed"
	}
	for _, record := range records {
		if !record.Found {
			return "failed"
		}
	}
	return "verified"
}
