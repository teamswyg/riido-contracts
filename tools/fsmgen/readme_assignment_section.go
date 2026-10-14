package main

func generatedAssignmentReadmeSection(metadata map[string]fsmMetadata) (readmeSection, error) {
	meta, err := requireFSMMetadata(metadata, "assignment", "AssignmentTransitionCode")
	if err != nil {
		return readmeSection{}, err
	}
	startStates, err := assignmentStateCodesFromConsts(meta.StartPoints)
	if err != nil {
		return readmeSection{}, err
	}
	endStates, err := assignmentStateCodesFromConsts(meta.EndPoints)
	if err != nil {
		return readmeSection{}, err
	}
	return readmeSection{
		ID:      meta.ReadmeSection,
		Content: mermaidFence(assignmentMermaid(startStates, endStates)),
	}, nil
}
