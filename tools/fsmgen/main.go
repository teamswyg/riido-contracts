package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"go/format"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"

	"github.com/teamswyg/riido-contracts/assignment"
	"github.com/teamswyg/riido-contracts/ir"
	"github.com/teamswyg/riido-contracts/task"
)

const (
	modulePath  = "github.com/teamswyg/riido-contracts"
	sourcePath  = "enumgen/enums.lisp"
	generatedBy = "fsm gen"
	readmePath  = "README.md"
)

type generatedArtifact struct {
	Path string
	Body []byte
}

type readmeSection struct {
	ID      string
	Content string
}

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "fsmgen:", err)
		os.Exit(1)
	}
}

func run(args []string, out io.Writer) error {
	command := "verify"
	if len(args) > 0 && !strings.HasPrefix(args[0], "-") {
		command = args[0]
		args = args[1:]
	}
	fs := flag.NewFlagSet("fsmgen", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	if err := fs.Parse(args); err != nil {
		return err
	}
	root, err := findRepoRoot()
	if err != nil {
		return err
	}
	files, err := generatedFiles()
	if err != nil {
		return err
	}
	sections := generatedReadmeSections()
	switch command {
	case "generate":
		for _, file := range files {
			path := filepath.Join(root, filepath.FromSlash(file.Path))
			if err := os.WriteFile(path, file.Body, 0o644); err != nil {
				return fmt.Errorf("write %s: %w", file.Path, err)
			}
		}
		if err := writeReadmeSections(root, sections); err != nil {
			return err
		}
		fmt.Fprintf(out, "fsmgen: generated %d files and %d README sections\n", len(files), len(sections))
		return nil
	case "verify":
		for _, file := range files {
			path := filepath.Join(root, filepath.FromSlash(file.Path))
			got, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("read %s: %w", file.Path, err)
			}
			if !bytes.Equal(got, file.Body) {
				return fmt.Errorf("%s drifted; run go run ./tools/fsmgen generate", file.Path)
			}
		}
		if err := verifyReadmeSections(root, sections); err != nil {
			return err
		}
		fmt.Fprintf(out, "fsmgen: verified %d files and %d README sections\n", len(files), len(sections))
		return nil
	default:
		return errors.New("usage: go run ./tools/fsmgen [verify|generate]")
	}
}

func findRepoRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		body, err := os.ReadFile(filepath.Join(dir, "go.mod"))
		if err == nil && strings.Contains(string(body), "module "+modulePath) {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", errors.New("cannot find riido-contracts repo root")
		}
		dir = parent
	}
}

func generatedFiles() ([]generatedArtifact, error) {
	taskBody, err := generateTaskFSMFile()
	if err != nil {
		return nil, err
	}
	assignmentBody, err := generateAssignmentFSMFile()
	if err != nil {
		return nil, err
	}
	return []generatedArtifact{
		{Path: "task/task_fsm_gen.go", Body: taskBody},
		{Path: "assignment/assignment_fsm_gen.go", Body: assignmentBody},
	}, nil
}

func generateTaskFSMFile() ([]byte, error) {
	var b bytes.Buffer
	writeHeader(&b, "task")
	fmt.Fprintln(&b, "import (")
	fmt.Fprintln(&b, "\t\"github.com/teamswyg/riido-contracts/ir\"")
	fmt.Fprintln(&b, ")")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "type TaskFSM interface {")
	fmt.Fprintln(&b, "\tName() string")
	fmt.Fprintln(&b, "\tStates() []TaskStateCode")
	fmt.Fprintln(&b, "\tTerminalStates() []TaskStateCode")
	fmt.Fprintln(&b, "\tTransitions() []TaskTransitionCode")
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
	fmt.Fprintln(&b, "\treturn \"task\"")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "func (generatedTaskFSM) States() []TaskStateCode {")
	fmt.Fprintln(&b, "\treturn AllTaskStateCodes()")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "func (generatedTaskFSM) TerminalStates() []TaskStateCode {")
	writeTaskStateCodeReturn(&b, terminalTaskStates())
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "func (generatedTaskFSM) Transitions() []TaskTransitionCode {")
	fmt.Fprintln(&b, "\treturn LegalTransitionCodes()")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "func (generatedTaskFSM) CanTransition(from TaskStateCode, to TaskStateCode, trigger ir.EventTypeCode) bool {")
	fmt.Fprintln(&b, "\treturn ValidateTransitionCode(from, to, trigger)")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	writeTaskNextStates(&b)
	fmt.Fprintf(&b, "const TaskFSMMermaid = `%s`\n\n", taskMermaid())
	fmt.Fprintln(&b, "func (generatedTaskFSM) Mermaid() string {")
	fmt.Fprintln(&b, "\treturn TaskFSMMermaid")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	return formatSource("task/task_fsm_gen.go", b.Bytes())
}

func generateAssignmentFSMFile() ([]byte, error) {
	var b bytes.Buffer
	writeHeader(&b, "assignment")
	fmt.Fprintln(&b, "type AssignmentFSM interface {")
	fmt.Fprintln(&b, "\tName() string")
	fmt.Fprintln(&b, "\tStates() []AssignmentStateCode")
	fmt.Fprintln(&b, "\tTerminalStates() []AssignmentStateCode")
	fmt.Fprintln(&b, "\tTransitions() []AssignmentTransitionCode")
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
	fmt.Fprintln(&b, "\treturn \"assignment\"")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "func (generatedAssignmentFSM) States() []AssignmentStateCode {")
	fmt.Fprintln(&b, "\treturn AllAssignmentStateCodes()")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "func (generatedAssignmentFSM) TerminalStates() []AssignmentStateCode {")
	writeAssignmentStateCodeReturn(&b, terminalAssignmentStates())
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "func (generatedAssignmentFSM) Transitions() []AssignmentTransitionCode {")
	fmt.Fprintln(&b, "\treturn AssignmentTransitionCodes()")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	fmt.Fprintln(&b, "func (generatedAssignmentFSM) CanTransition(from AssignmentStateCode, to AssignmentStateCode) bool {")
	fmt.Fprintln(&b, "\treturn CanTransitionCode(from, to)")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)
	writeAssignmentNextStates(&b)
	fmt.Fprintf(&b, "const AssignmentFSMMermaid = `%s`\n\n", assignmentMermaid())
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

func terminalTaskStates() []task.TaskStateCode {
	out := []task.TaskStateCode{}
	for _, state := range task.AllTaskStateCodes() {
		if state.IsTerminal() {
			out = append(out, state)
		}
	}
	return out
}

func terminalAssignmentStates() []assignment.AssignmentStateCode {
	out := []assignment.AssignmentStateCode{}
	for _, state := range assignment.AllAssignmentStateCodes() {
		if state.IsTerminal() {
			out = append(out, state)
		}
	}
	return out
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

func taskMermaid() string {
	var b strings.Builder
	b.WriteString("stateDiagram-v2\n")
	if states := task.AllTaskStateCodes(); len(states) > 0 {
		fmt.Fprintf(&b, "  [*] --> %s\n", mermaidNode(states[0].String()))
	}
	for _, transition := range task.LegalTransitionCodes() {
		fmt.Fprintf(&b, "  %s --> %s : %s\n", mermaidNode(transition.From.String()), mermaidNode(transition.To.String()), transition.Trigger.String())
	}
	for _, state := range terminalTaskStates() {
		fmt.Fprintf(&b, "  %s --> [*]\n", mermaidNode(state.String()))
	}
	return b.String()
}

func assignmentMermaid() string {
	var b strings.Builder
	b.WriteString("stateDiagram-v2\n")
	if states := assignment.AllAssignmentStateCodes(); len(states) > 0 {
		fmt.Fprintf(&b, "  [*] --> %s\n", mermaidNode(states[0].String()))
	}
	for _, transition := range assignment.AssignmentTransitionCodes() {
		fmt.Fprintf(&b, "  %s --> %s\n", mermaidNode(transition.From.String()), mermaidNode(transition.To.String()))
	}
	for _, state := range terminalAssignmentStates() {
		fmt.Fprintf(&b, "  %s --> [*]\n", mermaidNode(state.String()))
	}
	return b.String()
}

func generatedReadmeSections() []readmeSection {
	return []readmeSection{
		{
			ID:      "task",
			Content: mermaidFence(taskMermaid()),
		},
		{
			ID:      "assignment",
			Content: mermaidFence(assignmentMermaid()),
		},
	}
}

func mermaidFence(body string) string {
	return "```mermaid\n" + body + "```\n"
}

func writeReadmeSections(root string, sections []readmeSection) error {
	path := filepath.Join(root, readmePath)
	body, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", readmePath, err)
	}
	updated := string(body)
	for _, section := range sections {
		var replaceErr error
		updated, replaceErr = replaceSection(updated, section)
		if replaceErr != nil {
			return replaceErr
		}
	}
	if err := os.WriteFile(path, []byte(updated), 0o644); err != nil {
		return fmt.Errorf("write %s: %w", readmePath, err)
	}
	return nil
}

func verifyReadmeSections(root string, sections []readmeSection) error {
	path := filepath.Join(root, readmePath)
	body, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", readmePath, err)
	}
	current := string(body)
	for _, section := range sections {
		got, err := extractSection(current, section.ID)
		if err != nil {
			return err
		}
		if got != section.Content {
			return fmt.Errorf("%s section %q drifted; run go run ./tools/fsmgen generate", readmePath, section.ID)
		}
	}
	return nil
}

func replaceSection(body string, section readmeSection) (string, error) {
	start := sectionStart(section.ID)
	end := sectionEnd(section.ID)
	startIndex := strings.Index(body, start)
	if startIndex < 0 {
		return "", fmt.Errorf("%s missing marker %s", readmePath, start)
	}
	contentStart := startIndex + len(start)
	endIndex := strings.Index(body[contentStart:], end)
	if endIndex < 0 {
		return "", fmt.Errorf("%s missing marker %s", readmePath, end)
	}
	contentEnd := contentStart + endIndex
	replacement := "\n" + section.Content
	return body[:contentStart] + replacement + body[contentEnd:], nil
}

func extractSection(body, id string) (string, error) {
	start := sectionStart(id)
	end := sectionEnd(id)
	startIndex := strings.Index(body, start)
	if startIndex < 0 {
		return "", fmt.Errorf("%s missing marker %s", readmePath, start)
	}
	contentStart := startIndex + len(start)
	endIndex := strings.Index(body[contentStart:], end)
	if endIndex < 0 {
		return "", fmt.Errorf("%s missing marker %s", readmePath, end)
	}
	contentEnd := contentStart + endIndex
	content := body[contentStart:contentEnd]
	return strings.TrimPrefix(content, "\n"), nil
}

func sectionStart(id string) string {
	return "<!-- fsmgen:" + id + ":start -->"
}

func sectionEnd(id string) string {
	return "<!-- fsmgen:" + id + ":end -->"
}

func taskStateCodeRef(code task.TaskStateCode) string {
	return "TaskStateCode" + pascalCodeSuffix(code.String())
}

func eventTypeCodeRef(code ir.EventTypeCode) string {
	return "ir.EventTypeCode" + pascalCodeSuffix(code.String())
}

func assignmentStateCodeRef(code assignment.AssignmentStateCode) string {
	return "AssignmentStateCode" + pascalCodeSuffix(code.String())
}

func pascalCodeSuffix(value string) string {
	parts := strings.FieldsFunc(value, func(r rune) bool {
		return r == '_' || r == '-' || unicode.IsSpace(r)
	})
	for index, part := range parts {
		if part == "" {
			continue
		}
		runes := []rune(part)
		runes[0] = unicode.ToUpper(runes[0])
		parts[index] = string(runes)
	}
	return strings.Join(parts, "")
}

func mermaidNode(value string) string {
	if value == "" {
		return "unknown"
	}
	var b strings.Builder
	for index, r := range value {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			if index == 0 {
				b.WriteRune(unicode.ToUpper(r))
			} else {
				b.WriteRune(r)
			}
			continue
		}
		b.WriteByte('_')
	}
	return b.String()
}

func writeHeader(b *bytes.Buffer, packageName string) {
	fmt.Fprintf(b, "// Code generated by %s from %s; DO NOT EDIT.\n\n", generatedBy, sourcePath)
	fmt.Fprintf(b, "package %s\n\n", packageName)
}

func formatSource(name string, source []byte) ([]byte, error) {
	out, err := format.Source(source)
	if err != nil {
		return nil, fmt.Errorf("format %s: %w\n%s", name, err, source)
	}
	return out, nil
}
