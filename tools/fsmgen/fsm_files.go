package main

import (
	"bytes"
	"fmt"

	"github.com/teamswyg/riido-contracts/assignment"
	"github.com/teamswyg/riido-contracts/task"
)

type taskFSMModel struct {
	Meta        fsmMetadata
	StartStates []task.TaskStateCode
	EndStates   []task.TaskStateCode
}

type assignmentFSMModel struct {
	Meta        fsmMetadata
	StartStates []assignment.AssignmentStateCode
	EndStates   []assignment.AssignmentStateCode
}

type fsmSection[T any] struct {
	Path    string
	Package string
	Write   func(*bytes.Buffer, T)
	Imports []string
}

func generatedFSMFiles(metadata map[string]fsmMetadata) ([]generatedArtifact, error) {
	taskFiles, err := generatedTaskFSMFiles(metadata)
	if err != nil {
		return nil, err
	}
	assignmentFiles, err := generatedAssignmentFSMFiles(metadata)
	if err != nil {
		return nil, err
	}
	return append(taskFiles, assignmentFiles...), nil
}

func formatFSMSection[T any](section fsmSection[T], model T) (generatedArtifact, error) {
	var b bytes.Buffer
	writeHeader(&b, section.Package)
	writeImports(&b, section.Imports)
	section.Write(&b, model)
	body, err := formatSource(section.Path, b.Bytes())
	if err != nil {
		return generatedArtifact{}, err
	}
	return generatedArtifact{Path: section.Path, Body: body}, nil
}

func writeImports(b *bytes.Buffer, imports []string) {
	if len(imports) == 0 {
		return
	}
	fmt.Fprintln(b, "import (")
	for _, imp := range imports {
		fmt.Fprintf(b, "\t%q\n", imp)
	}
	fmt.Fprintln(b, ")")
	fmt.Fprintln(b)
}
