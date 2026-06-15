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
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

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

type fsmMetadata struct {
	Package        string
	TransitionName string
	FSMName        string
	TypeUnion      string
	StartPoints    []string
	EndPoints      []string
	ReadmeSection  string
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
	metadata, err := loadFSMMetadata(filepath.Join(root, filepath.FromSlash(sourcePath)))
	if err != nil {
		return err
	}
	files, err := generatedFiles(metadata)
	if err != nil {
		return err
	}
	sections, err := generatedReadmeSections(metadata)
	if err != nil {
		return err
	}
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

type node struct {
	atom string
	list []node
}

func (n node) isAtom() bool {
	return n.list == nil
}

func loadFSMMetadata(path string) (map[string]fsmMetadata, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}
	root, err := parseSExpr(string(body))
	if err != nil {
		return nil, err
	}
	if root.isAtom() || len(root.list) == 0 || atom(root.list[0]) != "enum-gen" {
		return nil, errors.New("root form must be (enum-gen ...)")
	}
	metadata := map[string]fsmMetadata{}
	for _, form := range root.list[1:] {
		if form.isAtom() || len(form.list) == 0 || atom(form.list[0]) != "transitions" {
			continue
		}
		spec, err := parseFSMMetadata(form)
		if err != nil {
			return nil, err
		}
		key := fsmMetadataKey(spec.Package, spec.TransitionName)
		if _, ok := metadata[key]; ok {
			return nil, fmt.Errorf("duplicate fsm metadata for %s", key)
		}
		metadata[key] = spec
	}
	return metadata, nil
}

func parseFSMMetadata(form node) (fsmMetadata, error) {
	var spec fsmMetadata
	for i := 1; i < len(form.list); i++ {
		item := form.list[i]
		if !item.isAtom() || !strings.HasPrefix(item.atom, ":") {
			continue
		}
		if i+1 >= len(form.list) {
			return fsmMetadata{}, fmt.Errorf("transition property %s missing value", item.atom)
		}
		key := strings.TrimPrefix(item.atom, ":")
		value := form.list[i+1]
		switch key {
		case "package":
			spec.Package = atom(value)
		case "name":
			spec.TransitionName = atom(value)
		case "fsm-name":
			spec.FSMName = atom(value)
		case "fsm-type-union":
			spec.TypeUnion = atom(value)
		case "start-points":
			spec.StartPoints = atomList(value)
		case "end-points":
			spec.EndPoints = atomList(value)
		case "readme-section":
			spec.ReadmeSection = atom(value)
		}
		i++
	}
	if spec.Package == "" || spec.TransitionName == "" {
		return fsmMetadata{}, errors.New("transitions block missing package or name")
	}
	if spec.FSMName == "" || spec.TypeUnion == "" || spec.ReadmeSection == "" {
		return fsmMetadata{}, fmt.Errorf("transitions %s missing fsm-name, fsm-type-union, or readme-section", spec.TransitionName)
	}
	if len(spec.StartPoints) == 0 || len(spec.EndPoints) == 0 {
		return fsmMetadata{}, fmt.Errorf("transitions %s missing start-points or end-points", spec.TransitionName)
	}
	return spec, nil
}

func parseSExpr(source string) (node, error) {
	tokens, err := lex(source)
	if err != nil {
		return node{}, err
	}
	index := 0
	root, err := parseNode(tokens, &index)
	if err != nil {
		return node{}, err
	}
	if index != len(tokens) {
		return node{}, fmt.Errorf("unexpected trailing token %q", tokens[index])
	}
	return root, nil
}

func lex(source string) ([]string, error) {
	out := []string{}
	for i := 0; i < len(source); {
		r, width := utf8.DecodeRuneInString(source[i:])
		switch {
		case unicode.IsSpace(r):
			i += width
		case r == ';':
			for i < len(source) && source[i] != '\n' {
				i++
			}
		case r == '(' || r == ')':
			out = append(out, string(r))
			i += width
		case r == '"':
			value, next, err := readStringToken(source, i)
			if err != nil {
				return nil, err
			}
			out = append(out, value)
			i = next
		default:
			start := i
			for i < len(source) {
				r, width = utf8.DecodeRuneInString(source[i:])
				if unicode.IsSpace(r) || r == '(' || r == ')' || r == ';' {
					break
				}
				i += width
			}
			out = append(out, source[start:i])
		}
	}
	return out, nil
}

func readStringToken(source string, start int) (string, int, error) {
	var b strings.Builder
	b.WriteByte('"')
	escaped := false
	for i := start + 1; i < len(source); i++ {
		ch := source[i]
		b.WriteByte(ch)
		if escaped {
			escaped = false
			continue
		}
		if ch == '\\' {
			escaped = true
			continue
		}
		if ch == '"' {
			return b.String(), i + 1, nil
		}
	}
	return "", 0, errors.New("unterminated string literal")
}

func parseNode(tokens []string, index *int) (node, error) {
	if *index >= len(tokens) {
		return node{}, errors.New("unexpected end of input")
	}
	token := tokens[*index]
	*index++
	switch token {
	case "(":
		var list []node
		for {
			if *index >= len(tokens) {
				return node{}, errors.New("unterminated list")
			}
			if tokens[*index] == ")" {
				*index++
				return node{list: list}, nil
			}
			child, err := parseNode(tokens, index)
			if err != nil {
				return node{}, err
			}
			list = append(list, child)
		}
	case ")":
		return node{}, errors.New("unexpected )")
	default:
		if strings.HasPrefix(token, "\"") {
			value, err := strconv.Unquote(token)
			if err != nil {
				return node{}, fmt.Errorf("decode string %s: %w", token, err)
			}
			return node{atom: value}, nil
		}
		return node{atom: token}, nil
	}
}

func atom(n node) string {
	if !n.isAtom() {
		return ""
	}
	return n.atom
}

func atomList(n node) []string {
	if n.isAtom() {
		value := atom(n)
		if value == "" {
			return nil
		}
		return []string{value}
	}
	out := make([]string, 0, len(n.list))
	for _, item := range n.list {
		value := atom(item)
		if value != "" {
			out = append(out, value)
		}
	}
	return out
}

func fsmMetadataKey(packageName, transitionName string) string {
	return packageName + "." + transitionName
}

func requireFSMMetadata(metadata map[string]fsmMetadata, packageName, transitionName string) (fsmMetadata, error) {
	key := fsmMetadataKey(packageName, transitionName)
	spec, ok := metadata[key]
	if !ok {
		return fsmMetadata{}, fmt.Errorf("missing fsm metadata for %s", key)
	}
	return spec, nil
}

func generatedFiles(metadata map[string]fsmMetadata) ([]generatedArtifact, error) {
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
	return []generatedArtifact{
		{Path: "task/task_fsm_gen.go", Body: taskBody},
		{Path: "assignment/assignment_fsm_gen.go", Body: assignmentBody},
	}, nil
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

func generatedReadmeSections(metadata map[string]fsmMetadata) ([]readmeSection, error) {
	taskMeta, err := requireFSMMetadata(metadata, "task", "TaskTransitionCode")
	if err != nil {
		return nil, err
	}
	taskStart, err := taskStateCodesFromConsts(taskMeta.StartPoints)
	if err != nil {
		return nil, err
	}
	taskEnd, err := taskStateCodesFromConsts(taskMeta.EndPoints)
	if err != nil {
		return nil, err
	}
	assignmentMeta, err := requireFSMMetadata(metadata, "assignment", "AssignmentTransitionCode")
	if err != nil {
		return nil, err
	}
	assignmentStart, err := assignmentStateCodesFromConsts(assignmentMeta.StartPoints)
	if err != nil {
		return nil, err
	}
	assignmentEnd, err := assignmentStateCodesFromConsts(assignmentMeta.EndPoints)
	if err != nil {
		return nil, err
	}
	return []readmeSection{
		{
			ID:      taskMeta.ReadmeSection,
			Content: mermaidFence(taskMermaid(taskStart, taskEnd)),
		},
		{
			ID:      assignmentMeta.ReadmeSection,
			Content: mermaidFence(assignmentMermaid(assignmentStart, assignmentEnd)),
		},
	}, nil
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
