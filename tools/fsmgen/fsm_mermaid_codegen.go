package main

import (
	"fmt"
	"strings"

	"github.com/teamswyg/riido-contracts/assignment"
	"github.com/teamswyg/riido-contracts/task"
)

func taskMermaid(startStates, endStates []task.TaskStateCode) string {
	var b strings.Builder
	b.WriteString("stateDiagram-v2\n")
	for _, state := range startStates {
		fmt.Fprintf(&b, "  [*] --> %s\n", mermaidNode(state.String()))
	}
	for _, transition := range task.LegalTransitionCodes() {
		fmt.Fprintf(&b, "  %s --> %s : %s\n", mermaidNode(transition.From.String()), mermaidNode(transition.To.String()), transition.Trigger.String())
	}
	for _, state := range endStates {
		fmt.Fprintf(&b, "  %s --> [*]\n", mermaidNode(state.String()))
	}
	return b.String()
}

func assignmentMermaid(startStates, endStates []assignment.AssignmentStateCode) string {
	var b strings.Builder
	b.WriteString("stateDiagram-v2\n")
	for _, state := range startStates {
		fmt.Fprintf(&b, "  [*] --> %s\n", mermaidNode(state.String()))
	}
	for _, transition := range assignment.AssignmentTransitionCodes() {
		fmt.Fprintf(&b, "  %s --> %s\n", mermaidNode(transition.From.String()), mermaidNode(transition.To.String()))
	}
	for _, state := range endStates {
		fmt.Fprintf(&b, "  %s --> [*]\n", mermaidNode(state.String()))
	}
	return b.String()
}
