package main

import (
	"bytes"
	"fmt"
	"sort"

	"github.com/teamswyg/riido-contracts/assignment"
)

func writeAssignmentFSMNextStateTable(b *bytes.Buffer, _ assignmentFSMModel) {
	fmt.Fprintln(b, "var assignmentFSMNextStates = map[AssignmentStateCode][]AssignmentStateCode{")
	for _, row := range assignmentFSMNextStateRows() {
		fmt.Fprintf(b, "\t%s: ", assignmentStateCodeRef(row.From))
		writeInlineAssignmentStateSlice(b, row.Next)
	}
	fmt.Fprintln(b, "}")
}

func writeAssignmentFSMNextStates(b *bytes.Buffer, _ assignmentFSMModel) {
	fmt.Fprintln(b, "func (generatedAssignmentFSM) CanTransition(from AssignmentStateCode, to AssignmentStateCode) bool {")
	fmt.Fprintln(b, "\treturn CanTransitionCode(from, to)")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func (generatedAssignmentFSM) NextStates(from AssignmentStateCode) []AssignmentStateCode {")
	fmt.Fprintln(b, "\tstates := assignmentFSMNextStates[from]")
	fmt.Fprintln(b, "\treturn append([]AssignmentStateCode(nil), states...)")
	fmt.Fprintln(b, "}")
}

type assignmentFSMNextStateRow struct {
	From assignment.AssignmentStateCode
	Next []assignment.AssignmentStateCode
}

func assignmentFSMNextStateRows() []assignmentFSMNextStateRow {
	groups := assignmentTransitionsByFrom()
	rows := []assignmentFSMNextStateRow{}
	for _, from := range assignment.AllAssignmentStateCodes() {
		next := append([]assignment.AssignmentStateCode{from}, groups[from]...)
		sortAssignmentStateCodes(next)
		rows = append(rows, assignmentFSMNextStateRow{From: from, Next: next})
	}
	return rows
}

func writeInlineAssignmentStateSlice(b *bytes.Buffer, states []assignment.AssignmentStateCode) {
	fmt.Fprint(b, "[]AssignmentStateCode{")
	for index, state := range states {
		if index > 0 {
			fmt.Fprint(b, ", ")
		}
		fmt.Fprint(b, assignmentStateCodeRef(state))
	}
	fmt.Fprintln(b, "},")
}

func sortAssignmentStateCodes(states []assignment.AssignmentStateCode) {
	sort.SliceStable(states, func(i, j int) bool {
		return states[i] < states[j]
	})
}
