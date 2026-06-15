package main

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/teamswyg/riido-contracts/assignment"
	"github.com/teamswyg/riido-contracts/ir"
	"github.com/teamswyg/riido-contracts/task"
)

func generatedFiles(root string, metadata map[string]fsmMetadata, patternDocs map[string]patternDocument) ([]generatedArtifact, error) {
	taskMeta, err := requireFSMMetadata(metadata, "task", "TaskTransitionCode")
	if err != nil {
		return nil, err
	}
	assignmentMeta, err := requireFSMMetadata(metadata, "assignment", "AssignmentTransitionCode")
	if err != nil {
		return nil, err
	}
	taskBody, err := generateTaskFSMFile(taskMeta)
	if err != nil {
		return nil, err
	}
	assignmentBody, err := generateAssignmentFSMFile(assignmentMeta)
	if err != nil {
		return nil, err
	}
	patternFiles, err := generatePatternFiles(root, patternDocs)
	if err != nil {
		return nil, err
	}
	files := []generatedArtifact{
		{Path: "task/task_fsm_gen.go", Body: taskBody},
		{Path: "assignment/assignment_fsm_gen.go", Body: assignmentBody},
	}
	files = append(patternFiles, files...)
	return files, nil
}

func generateTaskFSMFile(meta fsmMetadata) ([]byte, error) {
	startStates, err := taskStateCodesFromConsts(meta.StartPoints)
	if err != nil {
		return nil, err
	}
	endStates, err := taskStateCodesFromConsts(meta.EndPoints)
	if err != nil {
		return nil, err
	}
	var b bytes.Buffer
	writeHeader(&b, "task")
	fmt.Fprintln(&b, "import (")
	fmt.Fprintln(&b, "\t\"github.com/teamswyg/riido-contracts/ir\"")
	fmt.Fprintln(&b, ")")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "type TaskFSMTypeUnion string")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "const (")
	fmt.Fprintf(&b, "\tTaskFSMTypeUnion%s TaskFSMTypeUnion = %q\n", meta.TypeUnion, meta.TypeUnion)
	fmt.Fprintln(&b, ")")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "type TaskFSMPointKind uint8")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "const (")
	fmt.Fprintln(&b, "\tTaskFSMPointUnknown TaskFSMPointKind = iota")
	fmt.Fprintln(&b, "\tTaskFSMPointStart")
	fmt.Fprintln(&b, "\tTaskFSMPointIntermediate")
	fmt.Fprintln(&b, "\tTaskFSMPointEnd")
	fmt.Fprintln(&b, ")")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "type TaskFSM interface {")
	fmt.Fprintln(&b, "\tName() string")
	fmt.Fprintln(&b, "\tTypeUnion() TaskFSMTypeUnion")
	fmt.Fprintln(&b, "\tStates() []TaskStateCode")
	fmt.Fprintln(&b, "\tStartStates() []TaskStateCode")
	fmt.Fprintln(&b, "\tEndStates() []TaskStateCode")
	fmt.Fprintln(&b, "\tTerminalStates() []TaskStateCode")
	fmt.Fprintln(&b, "\tTransitions() []TaskTransitionCode")
	fmt.Fprintln(&b, "\tPointKind(state TaskStateCode) TaskFSMPointKind")
	fmt.Fprintln(&b, "\tIsStartState(state TaskStateCode) bool")
	fmt.Fprintln(&b, "\tIsEndState(state TaskStateCode) bool")
	fmt.Fprintln(&b, "\tCanTransition(from TaskStateCode, to TaskStateCode, trigger ir.EventTypeCode) bool")
	fmt.Fprintln(&b, "\tNextStates(from TaskStateCode, trigger ir.EventTypeCode) []TaskStateCode")
	fmt.Fprintln(&b, "\tMermaid() string")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "type TaskFSMServiceProvider interface {")
	fmt.Fprintln(&b, "\tTaskFSM() TaskFSM")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "type generatedTaskFSM struct{}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "type generatedTaskFSMServiceProvider struct{}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "func GeneratedTaskFSM() TaskFSM {")
	fmt.Fprintln(&b, "\treturn generatedTaskFSM{}")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "func GeneratedTaskFSMServiceProvider() TaskFSMServiceProvider {")
	fmt.Fprintln(&b, "\treturn generatedTaskFSMServiceProvider{}")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "func (generatedTaskFSMServiceProvider) TaskFSM() TaskFSM {")
	fmt.Fprintln(&b, "\treturn generatedTaskFSM{}")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "func (generatedTaskFSM) Name() string {")
	fmt.Fprintf(&b, "\treturn %q\n", meta.ReadmeSection)
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "func (generatedTaskFSM) TypeUnion() TaskFSMTypeUnion {")
	fmt.Fprintf(&b, "\treturn TaskFSMTypeUnion%s\n", meta.TypeUnion)
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "func (generatedTaskFSM) States() []TaskStateCode {")
	fmt.Fprintln(&b, "\treturn AllTaskStateCodes()")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "func (generatedTaskFSM) StartStates() []TaskStateCode {")
	writeTaskStateCodeReturn(&b, startStates)
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "func (generatedTaskFSM) EndStates() []TaskStateCode {")
	writeTaskStateCodeReturn(&b, endStates)
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "func (generatedTaskFSM) TerminalStates() []TaskStateCode {")
	writeTaskStateCodeReturn(&b, endStates)
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "func (generatedTaskFSM) Transitions() []TaskTransitionCode {")
	fmt.Fprintln(&b, "\treturn LegalTransitionCodes()")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	writeTaskPointMethods(&b, startStates, endStates)
	fmt.Fprintln(&b, "func (generatedTaskFSM) CanTransition(from TaskStateCode, to TaskStateCode, trigger ir.EventTypeCode) bool {")
	fmt.Fprintln(&b, "\treturn ValidateTransitionCode(from, to, trigger)")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	writeTaskNextStates(&b)
	fmt.Fprintf(&b, "const TaskFSMMermaid = `%s`\n\n", taskMermaid(startStates, endStates))
	fmt.Fprintln(&b, "func (generatedTaskFSM) Mermaid() string {")
	fmt.Fprintln(&b, "\treturn TaskFSMMermaid")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	return formatSource("task/task_fsm_gen.go", b.Bytes())
}

func generateAssignmentFSMFile(meta fsmMetadata) ([]byte, error) {
	startStates, err := assignmentStateCodesFromConsts(meta.StartPoints)
	if err != nil {
		return nil, err
	}
	endStates, err := assignmentStateCodesFromConsts(meta.EndPoints)
	if err != nil {
		return nil, err
	}
	var b bytes.Buffer
	writeHeader(&b, "assignment")
	fmt.Fprintln(&b, "type AssignmentFSMTypeUnion string")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "const (")
	fmt.Fprintf(&b, "\tAssignmentFSMTypeUnion%s AssignmentFSMTypeUnion = %q\n", meta.TypeUnion, meta.TypeUnion)
	fmt.Fprintln(&b, ")")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "type AssignmentFSMPointKind uint8")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "const (")
	fmt.Fprintln(&b, "\tAssignmentFSMPointUnknown AssignmentFSMPointKind = iota")
	fmt.Fprintln(&b, "\tAssignmentFSMPointStart")
	fmt.Fprintln(&b, "\tAssignmentFSMPointIntermediate")
	fmt.Fprintln(&b, "\tAssignmentFSMPointEnd")
	fmt.Fprintln(&b, ")")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "type AssignmentFSM interface {")
	fmt.Fprintln(&b, "\tName() string")
	fmt.Fprintln(&b, "\tTypeUnion() AssignmentFSMTypeUnion")
	fmt.Fprintln(&b, "\tStates() []AssignmentStateCode")
	fmt.Fprintln(&b, "\tStartStates() []AssignmentStateCode")
	fmt.Fprintln(&b, "\tEndStates() []AssignmentStateCode")
	fmt.Fprintln(&b, "\tTerminalStates() []AssignmentStateCode")
	fmt.Fprintln(&b, "\tTransitions() []AssignmentTransitionCode")
	fmt.Fprintln(&b, "\tPointKind(state AssignmentStateCode) AssignmentFSMPointKind")
	fmt.Fprintln(&b, "\tIsStartState(state AssignmentStateCode) bool")
	fmt.Fprintln(&b, "\tIsEndState(state AssignmentStateCode) bool")
	fmt.Fprintln(&b, "\tCanTransition(from AssignmentStateCode, to AssignmentStateCode) bool")
	fmt.Fprintln(&b, "\tNextStates(from AssignmentStateCode) []AssignmentStateCode")
	fmt.Fprintln(&b, "\tMermaid() string")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "type AssignmentFSMServiceProvider interface {")
	fmt.Fprintln(&b, "\tAssignmentFSM() AssignmentFSM")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "type generatedAssignmentFSM struct{}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "type generatedAssignmentFSMServiceProvider struct{}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "func GeneratedAssignmentFSM() AssignmentFSM {")
	fmt.Fprintln(&b, "\treturn generatedAssignmentFSM{}")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "func GeneratedAssignmentFSMServiceProvider() AssignmentFSMServiceProvider {")
	fmt.Fprintln(&b, "\treturn generatedAssignmentFSMServiceProvider{}")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "func (generatedAssignmentFSMServiceProvider) AssignmentFSM() AssignmentFSM {")
	fmt.Fprintln(&b, "\treturn generatedAssignmentFSM{}")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "func (generatedAssignmentFSM) Name() string {")
	fmt.Fprintf(&b, "\treturn %q\n", meta.ReadmeSection)
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "func (generatedAssignmentFSM) TypeUnion() AssignmentFSMTypeUnion {")
	fmt.Fprintf(&b, "\treturn AssignmentFSMTypeUnion%s\n", meta.TypeUnion)
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "func (generatedAssignmentFSM) States() []AssignmentStateCode {")
	fmt.Fprintln(&b, "\treturn AllAssignmentStateCodes()")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "func (generatedAssignmentFSM) StartStates() []AssignmentStateCode {")
	writeAssignmentStateCodeReturn(&b, startStates)
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "func (generatedAssignmentFSM) EndStates() []AssignmentStateCode {")
	writeAssignmentStateCodeReturn(&b, endStates)
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "func (generatedAssignmentFSM) TerminalStates() []AssignmentStateCode {")
	writeAssignmentStateCodeReturn(&b, endStates)
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "func (generatedAssignmentFSM) Transitions() []AssignmentTransitionCode {")
	fmt.Fprintln(&b, "\treturn AssignmentTransitionCodes()")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	writeAssignmentPointMethods(&b, startStates, endStates)
	fmt.Fprintln(&b, "func (generatedAssignmentFSM) CanTransition(from AssignmentStateCode, to AssignmentStateCode) bool {")
	fmt.Fprintln(&b, "\treturn CanTransitionCode(from, to)")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	writeAssignmentNextStates(&b)
	fmt.Fprintf(&b, "const AssignmentFSMMermaid = `%s`\n\n", assignmentMermaid(startStates, endStates))
	fmt.Fprintln(&b, "func (generatedAssignmentFSM) Mermaid() string {")
	fmt.Fprintln(&b, "\treturn AssignmentFSMMermaid")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	return formatSource("assignment/assignment_fsm_gen.go", b.Bytes())
}

func writeTaskNextStates(b *bytes.Buffer) {
	groups := taskTransitionsByFromAndTrigger()
	fromStates := task.AllTaskStateCodes()
	fmt.Fprintln(b, "func (generatedTaskFSM) NextStates(from TaskStateCode, trigger ir.EventTypeCode) []TaskStateCode {")
	fmt.Fprintln(b, "\tswitch from {")
	for _, from := range fromStates {
		byTrigger := groups[from]
		if len(byTrigger) == 0 {
			continue
		}
		fmt.Fprintf(b, "\tcase %s:\n", taskStateCodeRef(from))
		fmt.Fprintln(b, "\t\tswitch trigger {")
		triggers := sortedTaskTriggers(byTrigger)
		for _, trigger := range triggers {
			fmt.Fprintf(b, "\t\tcase %s:\n", eventTypeCodeRef(trigger))
			writeTaskStateCodeReturnWithIndent(b, "\t\t\t", byTrigger[trigger])
		}
		fmt.Fprintln(b, "\t\t}")
	}
	fmt.Fprintln(b, "\t}")
	fmt.Fprintln(b, "\treturn nil")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
}

func writeAssignmentNextStates(b *bytes.Buffer) {
	groups := assignmentTransitionsByFrom()
	fmt.Fprintln(b, "func (generatedAssignmentFSM) NextStates(from AssignmentStateCode) []AssignmentStateCode {")
	fmt.Fprintln(b, "\tswitch from {")
	for _, from := range assignment.AllAssignmentStateCodes() {
		next := groups[from]
		if !from.IsKnown() {
			continue
		}
		next = append([]assignment.AssignmentStateCode{from}, next...)
		sort.SliceStable(next, func(i, j int) bool {
			return next[i] < next[j]
		})
		fmt.Fprintf(b, "\tcase %s:\n", assignmentStateCodeRef(from))
		writeAssignmentStateCodeReturnWithIndent(b, "\t\t", next)
	}
	fmt.Fprintln(b, "\t}")
	fmt.Fprintln(b, "\treturn nil")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
}

func writeTaskPointMethods(b *bytes.Buffer, startStates, endStates []task.TaskStateCode) {
	fmt.Fprintln(b, "func (fsm generatedTaskFSM) PointKind(state TaskStateCode) TaskFSMPointKind {")
	fmt.Fprintln(b, "\tswitch {")
	fmt.Fprintln(b, "\tcase fsm.IsStartState(state):")
	fmt.Fprintln(b, "\t\treturn TaskFSMPointStart")
	fmt.Fprintln(b, "\tcase fsm.IsEndState(state):")
	fmt.Fprintln(b, "\t\treturn TaskFSMPointEnd")
	fmt.Fprintln(b, "\tcase state.IsKnown():")
	fmt.Fprintln(b, "\t\treturn TaskFSMPointIntermediate")
	fmt.Fprintln(b, "\tdefault:")
	fmt.Fprintln(b, "\t\treturn TaskFSMPointUnknown")
	fmt.Fprintln(b, "\t}")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func (generatedTaskFSM) IsStartState(state TaskStateCode) bool {")
	writeTaskStateSwitchReturn(b, startStates)
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func (generatedTaskFSM) IsEndState(state TaskStateCode) bool {")
	writeTaskStateSwitchReturn(b, endStates)
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
}

func writeAssignmentPointMethods(b *bytes.Buffer, startStates, endStates []assignment.AssignmentStateCode) {
	fmt.Fprintln(b, "func (fsm generatedAssignmentFSM) PointKind(state AssignmentStateCode) AssignmentFSMPointKind {")
	fmt.Fprintln(b, "\tswitch {")
	fmt.Fprintln(b, "\tcase fsm.IsStartState(state):")
	fmt.Fprintln(b, "\t\treturn AssignmentFSMPointStart")
	fmt.Fprintln(b, "\tcase fsm.IsEndState(state):")
	fmt.Fprintln(b, "\t\treturn AssignmentFSMPointEnd")
	fmt.Fprintln(b, "\tcase state.IsKnown():")
	fmt.Fprintln(b, "\t\treturn AssignmentFSMPointIntermediate")
	fmt.Fprintln(b, "\tdefault:")
	fmt.Fprintln(b, "\t\treturn AssignmentFSMPointUnknown")
	fmt.Fprintln(b, "\t}")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func (generatedAssignmentFSM) IsStartState(state AssignmentStateCode) bool {")
	writeAssignmentStateSwitchReturn(b, startStates)
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
	fmt.Fprintln(b, "func (generatedAssignmentFSM) IsEndState(state AssignmentStateCode) bool {")
	writeAssignmentStateSwitchReturn(b, endStates)
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
}

func writeTaskStateSwitchReturn(b *bytes.Buffer, states []task.TaskStateCode) {
	fmt.Fprintln(b, "\tswitch state {")
	for _, state := range states {
		fmt.Fprintf(b, "\tcase %s:\n", taskStateCodeRef(state))
		fmt.Fprintln(b, "\t\treturn true")
	}
	fmt.Fprintln(b, "\tdefault:")
	fmt.Fprintln(b, "\t\treturn false")
	fmt.Fprintln(b, "\t}")
}

func writeAssignmentStateSwitchReturn(b *bytes.Buffer, states []assignment.AssignmentStateCode) {
	fmt.Fprintln(b, "\tswitch state {")
	for _, state := range states {
		fmt.Fprintf(b, "\tcase %s:\n", assignmentStateCodeRef(state))
		fmt.Fprintln(b, "\t\treturn true")
	}
	fmt.Fprintln(b, "\tdefault:")
	fmt.Fprintln(b, "\t\treturn false")
	fmt.Fprintln(b, "\t}")
}

func taskTransitionsByFromAndTrigger() map[task.TaskStateCode]map[ir.EventTypeCode][]task.TaskStateCode {
	groups := map[task.TaskStateCode]map[ir.EventTypeCode][]task.TaskStateCode{}
	for _, transition := range task.LegalTransitionCodes() {
		if _, ok := groups[transition.From]; !ok {
			groups[transition.From] = map[ir.EventTypeCode][]task.TaskStateCode{}
		}
		groups[transition.From][transition.Trigger] = append(groups[transition.From][transition.Trigger], transition.To)
	}
	for _, byTrigger := range groups {
		for trigger := range byTrigger {
			sort.SliceStable(byTrigger[trigger], func(i, j int) bool {
				return byTrigger[trigger][i] < byTrigger[trigger][j]
			})
		}
	}
	return groups
}

func assignmentTransitionsByFrom() map[assignment.AssignmentStateCode][]assignment.AssignmentStateCode {
	groups := map[assignment.AssignmentStateCode][]assignment.AssignmentStateCode{}
	for _, transition := range assignment.AssignmentTransitionCodes() {
		groups[transition.From] = append(groups[transition.From], transition.To)
	}
	for from := range groups {
		sort.SliceStable(groups[from], func(i, j int) bool {
			return groups[from][i] < groups[from][j]
		})
	}
	return groups
}

func sortedTaskTriggers(groups map[ir.EventTypeCode][]task.TaskStateCode) []ir.EventTypeCode {
	triggers := make([]ir.EventTypeCode, 0, len(groups))
	for trigger := range groups {
		triggers = append(triggers, trigger)
	}
	sort.SliceStable(triggers, func(i, j int) bool {
		return triggers[i] < triggers[j]
	})
	return triggers
}

func taskStateCodesFromConsts(names []string) ([]task.TaskStateCode, error) {
	out := make([]task.TaskStateCode, 0, len(names))
	for _, name := range names {
		code, ok := taskStateCodeFromConst(name)
		if !ok {
			return nil, fmt.Errorf("unknown task state const %s", name)
		}
		out = append(out, code)
	}
	return out, nil
}

func taskStateCodeFromConst(name string) (task.TaskStateCode, bool) {
	for _, code := range task.AllTaskStateCodes() {
		if taskStateConstName(code) == name {
			return code, true
		}
	}
	return task.TaskStateCodeUnknown, false
}

func taskStateConstName(code task.TaskStateCode) string {
	return "State" + pascalCodeSuffix(code.String())
}

func assignmentStateCodesFromConsts(names []string) ([]assignment.AssignmentStateCode, error) {
	out := make([]assignment.AssignmentStateCode, 0, len(names))
	for _, name := range names {
		code, ok := assignmentStateCodeFromConst(name)
		if !ok {
			return nil, fmt.Errorf("unknown assignment state const %s", name)
		}
		out = append(out, code)
	}
	return out, nil
}

func assignmentStateCodeFromConst(name string) (assignment.AssignmentStateCode, bool) {
	for _, code := range assignment.AllAssignmentStateCodes() {
		if assignmentStateConstName(code) == name {
			return code, true
		}
	}
	return assignment.AssignmentStateCodeUnknown, false
}

func assignmentStateConstName(code assignment.AssignmentStateCode) string {
	return "Assignment" + pascalCodeSuffix(code.String())
}

func writeTaskStateCodeReturn(b *bytes.Buffer, states []task.TaskStateCode) {
	writeTaskStateCodeReturnWithIndent(b, "\t", states)
}

func writeTaskStateCodeReturnWithIndent(b *bytes.Buffer, indent string, states []task.TaskStateCode) {
	fmt.Fprintf(b, "%sreturn []TaskStateCode{\n", indent)
	for _, state := range states {
		fmt.Fprintf(b, "%s\t%s,\n", indent, taskStateCodeRef(state))
	}
	fmt.Fprintf(b, "%s}\n", indent)
}

func writeAssignmentStateCodeReturn(b *bytes.Buffer, states []assignment.AssignmentStateCode) {
	writeAssignmentStateCodeReturnWithIndent(b, "\t", states)
}

func writeAssignmentStateCodeReturnWithIndent(b *bytes.Buffer, indent string, states []assignment.AssignmentStateCode) {
	fmt.Fprintf(b, "%sreturn []AssignmentStateCode{\n", indent)
	for _, state := range states {
		fmt.Fprintf(b, "%s\t%s,\n", indent, assignmentStateCodeRef(state))
	}
	fmt.Fprintf(b, "%s}\n", indent)
}

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
