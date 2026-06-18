package main

import (
	"bytes"
	"fmt"

	"github.com/teamswyg/riido-contracts/task"
)

func generatedTaskFSMFiles(metadata map[string]fsmMetadata) ([]generatedArtifact, error) {
	meta, err := requireFSMMetadata(metadata, "task", "TaskTransitionCode")
	if err != nil {
		return nil, err
	}
	startStates, err := taskStateCodesFromConsts(meta.StartPoints)
	if err != nil {
		return nil, err
	}
	endStates, err := taskStateCodesFromConsts(meta.EndPoints)
	if err != nil {
		return nil, err
	}
	model := taskFSMModel{Meta: meta, StartStates: startStates, EndStates: endStates}
	return formatTaskFSMSections(model)
}

func formatTaskFSMSections(model taskFSMModel) ([]generatedArtifact, error) {
	sections := []fsmSection[taskFSMModel]{
		{"task/task_fsm_types_gen.go", "task", writeTaskFSMTypes, []string{modulePath + "/ir"}},
		{"task/task_fsm_provider_gen.go", "task", writeTaskFSMProvider, nil},
		{"task/task_fsm_states_gen.go", "task", writeTaskFSMStates, nil},
		{"task/task_fsm_points_gen.go", "task", writeTaskFSMPoints, nil},
		{"task/task_fsm_next_state_table_gen.go", "task", writeTaskFSMNextStateTable, []string{modulePath + "/ir"}},
		{"task/task_fsm_next_states_gen.go", "task", writeTaskFSMNextStates, []string{modulePath + "/ir"}},
		{"task/task_fsm_mermaid_gen.go", "task", writeTaskFSMMermaid, nil},
	}
	files := make([]generatedArtifact, 0, len(sections))
	for _, section := range sections {
		file, err := formatFSMSection(section, model)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	return files, nil
}

func writeTaskStateSliceVar(b *bytes.Buffer, name string, states []task.TaskStateCode) {
	fmt.Fprintf(b, "var %s = []TaskStateCode{\n", name)
	for _, state := range states {
		fmt.Fprintf(b, "\t%s,\n", taskStateCodeRef(state))
	}
	fmt.Fprintln(b, "}")
}
