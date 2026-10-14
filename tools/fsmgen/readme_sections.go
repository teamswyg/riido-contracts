package main

func generatedReadmeSections(metadata map[string]fsmMetadata) ([]readmeSection, error) {
	taskSection, err := generatedTaskReadmeSection(metadata)
	if err != nil {
		return nil, err
	}
	assignmentSection, err := generatedAssignmentReadmeSection(metadata)
	if err != nil {
		return nil, err
	}
	return []readmeSection{taskSection, assignmentSection}, nil
}

func mermaidFence(body string) string {
	return "```mermaid\n" + body + "```\n"
}
