package main

func generatedTaskReadmeSection(metadata map[string]fsmMetadata) (readmeSection, error) {
	meta, err := requireFSMMetadata(metadata, "task", "TaskTransitionCode")
	if err != nil {
		return readmeSection{}, err
	}
	startStates, err := taskStateCodesFromConsts(meta.StartPoints)
	if err != nil {
		return readmeSection{}, err
	}
	endStates, err := taskStateCodesFromConsts(meta.EndPoints)
	if err != nil {
		return readmeSection{}, err
	}
	return readmeSection{
		ID:      meta.ReadmeSection,
		Content: mermaidFence(taskMermaid(startStates, endStates)),
	}, nil
}
