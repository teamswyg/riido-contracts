package main

import (
	"bytes"
	"fmt"

	"github.com/teamswyg/riido-contracts/assignment"
)

func generatedAssignmentFSMFiles(metadata map[string]fsmMetadata) ([]generatedArtifact, error) {
	meta, err := requireFSMMetadata(metadata, "assignment", "AssignmentTransitionCode")
	if err != nil {
		return nil, err
	}
	startStates, err := assignmentStateCodesFromConsts(meta.StartPoints)
	if err != nil {
		return nil, err
	}
	endStates, err := assignmentStateCodesFromConsts(meta.EndPoints)
	if err != nil {
		return nil, err
	}
	model := assignmentFSMModel{Meta: meta, StartStates: startStates, EndStates: endStates}
	return formatAssignmentFSMSections(model)
}

func formatAssignmentFSMSections(model assignmentFSMModel) ([]generatedArtifact, error) {
	sections := []fsmSection[assignmentFSMModel]{
		{"assignment/assignment_fsm_types_gen.go", "assignment", writeAssignmentFSMTypes, nil},
		{"assignment/assignment_fsm_provider_gen.go", "assignment", writeAssignmentFSMProvider, nil},
		{"assignment/assignment_fsm_states_gen.go", "assignment", writeAssignmentFSMStates, nil},
		{"assignment/assignment_fsm_points_gen.go", "assignment", writeAssignmentFSMPoints, nil},
		{"assignment/assignment_fsm_next_state_table_gen.go", "assignment", writeAssignmentFSMNextStateTable, nil},
		{"assignment/assignment_fsm_next_states_gen.go", "assignment", writeAssignmentFSMNextStates, nil},
		{"assignment/assignment_fsm_mermaid_gen.go", "assignment", writeAssignmentFSMMermaid, nil},
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

func writeAssignmentStateSliceVar(b *bytes.Buffer, name string, states []assignment.AssignmentStateCode) {
	fmt.Fprintf(b, "var %s = []AssignmentStateCode{\n", name)
	for _, state := range states {
		fmt.Fprintf(b, "\t%s,\n", assignmentStateCodeRef(state))
	}
	fmt.Fprintln(b, "}")
}
