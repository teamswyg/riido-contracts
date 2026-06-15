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
)

const (
	modulePath  = "github.com/teamswyg/riido-contracts"
	sourcePath  = "enumgen/enums.lisp"
	generatedBy = "enum gen"
)

type document struct {
	Enums       []enumSpec
	Transitions []transitionSpec
}

type enumSpec struct {
	Package     string
	Type        string
	CodeType    string
	StringType  string
	ConstPrefix string
	AllFunc     string
	CodeAllFunc string
	Values      []enumValue
}

type enumValue struct {
	Const string
	Value string
	Attrs map[string]string
}

type transitionSpec struct {
	Package   string
	Name      string
	FromEnum  string
	ToEnum    string
	EventEnum string
	AllFunc   string
	Validate  string
	AllowSame bool
	Entries   []transitionEntry
}

type transitionEntry struct {
	From  string
	To    string
	Event string
}

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "enumgen:", err)
		os.Exit(1)
	}
}

func run(args []string, out io.Writer) error {
	command := "verify"
	if len(args) > 0 && !strings.HasPrefix(args[0], "-") {
		command = args[0]
		args = args[1:]
	}
	fs := flag.NewFlagSet("enumgen", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	if err := fs.Parse(args); err != nil {
		return err
	}
	root, err := findRepoRoot()
	if err != nil {
		return err
	}
	doc, err := loadDocument(filepath.Join(root, filepath.FromSlash(sourcePath)))
	if err != nil {
		return err
	}
	files, err := generatedFiles(doc)
	if err != nil {
		return err
	}
	switch command {
	case "generate":
		for name, body := range files {
			path := filepath.Join(root, filepath.FromSlash(name))
			if err := os.WriteFile(path, body, 0o644); err != nil {
				return fmt.Errorf("write %s: %w", name, err)
			}
		}
		fmt.Fprintf(out, "enumgen: generated %d files\n", len(files))
		return nil
	case "verify":
		for name, want := range files {
			path := filepath.Join(root, filepath.FromSlash(name))
			got, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("read %s: %w", name, err)
			}
			if !bytes.Equal(got, want) {
				return fmt.Errorf("%s drifted; run go run ./tools/enumgen generate", name)
			}
		}
		fmt.Fprintf(out, "enumgen: verified %d files\n", len(files))
		return nil
	default:
		return errors.New("usage: go run ./tools/enumgen [verify|generate]")
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

func loadDocument(path string) (document, error) {
	body, err := os.ReadFile(path)
	if err != nil {
		return document{}, fmt.Errorf("read %s: %w", path, err)
	}
	root, err := parseSExpr(string(body))
	if err != nil {
		return document{}, err
	}
	return documentFromNode(root)
}

type node struct {
	atom string
	list []node
}

func (n node) isAtom() bool {
	return n.list == nil
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

func documentFromNode(root node) (document, error) {
	if root.isAtom() || len(root.list) == 0 || atom(root.list[0]) != "enum-gen" {
		return document{}, errors.New("root form must be (enum-gen ...)")
	}
	var doc document
	for _, form := range root.list[1:] {
		if form.isAtom() || len(form.list) == 0 {
			return document{}, errors.New("enum-gen children must be lists")
		}
		switch atom(form.list[0]) {
		case "enum":
			spec, err := parseEnum(form)
			if err != nil {
				return document{}, err
			}
			doc.Enums = append(doc.Enums, spec)
		case "transitions":
			spec, err := parseTransitions(form)
			if err != nil {
				return document{}, err
			}
			doc.Transitions = append(doc.Transitions, spec)
		default:
			return document{}, fmt.Errorf("unknown form %q", atom(form.list[0]))
		}
	}
	return doc, validateDocument(doc)
}

func parseEnum(form node) (enumSpec, error) {
	props := map[string]string{}
	values := []enumValue{}
	for i := 1; i < len(form.list); i++ {
		item := form.list[i]
		if item.isAtom() && strings.HasPrefix(item.atom, ":") {
			if i+1 >= len(form.list) {
				return enumSpec{}, fmt.Errorf("enum property %s missing value", item.atom)
			}
			props[strings.TrimPrefix(item.atom, ":")] = atom(form.list[i+1])
			i++
			continue
		}
		if item.isAtom() || len(item.list) == 0 || atom(item.list[0]) != "value" {
			return enumSpec{}, errors.New("enum entries must be (value ...)")
		}
		value, err := parseEnumValue(item)
		if err != nil {
			return enumSpec{}, err
		}
		values = append(values, value)
	}
	spec := enumSpec{
		Package:     props["package"],
		Type:        props["type"],
		CodeType:    props["code-type"],
		StringType:  props["string-type"],
		ConstPrefix: props["const-prefix"],
		AllFunc:     props["all"],
		CodeAllFunc: props["code-all"],
		Values:      values,
	}
	if spec.Package == "" || spec.Type == "" || spec.CodeType == "" || spec.StringType == "" || spec.AllFunc == "" || spec.CodeAllFunc == "" {
		return enumSpec{}, fmt.Errorf("enum %q is missing required properties", spec.Type)
	}
	if len(spec.Values) == 0 {
		return enumSpec{}, fmt.Errorf("enum %s has no values", spec.Type)
	}
	return spec, nil
}

func parseEnumValue(form node) (enumValue, error) {
	if len(form.list) < 3 {
		return enumValue{}, errors.New("value requires const and string")
	}
	value := enumValue{
		Const: atom(form.list[1]),
		Value: atom(form.list[2]),
		Attrs: map[string]string{},
	}
	for i := 3; i < len(form.list); i++ {
		key := atom(form.list[i])
		if !strings.HasPrefix(key, ":") {
			return enumValue{}, fmt.Errorf("value %s has invalid attr %q", value.Const, key)
		}
		if i+1 >= len(form.list) {
			return enumValue{}, fmt.Errorf("value %s attr %s missing value", value.Const, key)
		}
		value.Attrs[strings.TrimPrefix(key, ":")] = atom(form.list[i+1])
		i++
	}
	return value, nil
}

func parseTransitions(form node) (transitionSpec, error) {
	props := map[string]string{}
	entries := []transitionEntry{}
	for i := 1; i < len(form.list); i++ {
		item := form.list[i]
		if item.isAtom() && strings.HasPrefix(item.atom, ":") {
			if i+1 >= len(form.list) {
				return transitionSpec{}, fmt.Errorf("transition property %s missing value", item.atom)
			}
			props[strings.TrimPrefix(item.atom, ":")] = atom(form.list[i+1])
			i++
			continue
		}
		if item.isAtom() || len(item.list) == 0 || atom(item.list[0]) != "transition" {
			return transitionSpec{}, errors.New("transition entries must be (transition ...)")
		}
		entry, err := parseTransitionEntry(item)
		if err != nil {
			return transitionSpec{}, err
		}
		entries = append(entries, entry)
	}
	spec := transitionSpec{
		Package:   props["package"],
		Name:      props["name"],
		FromEnum:  props["from-enum"],
		ToEnum:    props["to-enum"],
		EventEnum: props["event-enum"],
		AllFunc:   props["all"],
		Validate:  props["validate"],
		AllowSame: props["allow-same"] == "true",
		Entries:   entries,
	}
	if spec.Package == "" || spec.Name == "" || spec.FromEnum == "" || spec.ToEnum == "" || spec.AllFunc == "" || spec.Validate == "" {
		return transitionSpec{}, fmt.Errorf("transitions %q missing required properties", spec.Name)
	}
	if len(spec.Entries) == 0 {
		return transitionSpec{}, fmt.Errorf("transitions %s has no entries", spec.Name)
	}
	return spec, nil
}

func parseTransitionEntry(form node) (transitionEntry, error) {
	if len(form.list) != 3 && len(form.list) != 4 {
		return transitionEntry{}, errors.New("transition requires from, to, and optional event")
	}
	entry := transitionEntry{From: atom(form.list[1]), To: atom(form.list[2])}
	if len(form.list) == 4 {
		entry.Event = atom(form.list[3])
	}
	return entry, nil
}

func atom(n node) string {
	if !n.isAtom() {
		return ""
	}
	return n.atom
}

func validateDocument(doc document) error {
	enums := map[string]enumSpec{}
	for _, enum := range doc.Enums {
		ref := enum.Package + "." + enum.Type
		if _, ok := enums[ref]; ok {
			return fmt.Errorf("duplicate enum %s", ref)
		}
		seenConsts := map[string]bool{}
		seenValues := map[string]bool{}
		for _, value := range enum.Values {
			if value.Const == "" || value.Value == "" {
				return fmt.Errorf("enum %s has empty const or value", ref)
			}
			if seenConsts[value.Const] {
				return fmt.Errorf("enum %s duplicate const %s", ref, value.Const)
			}
			if seenValues[value.Value] {
				return fmt.Errorf("enum %s duplicate value %s", ref, value.Value)
			}
			seenConsts[value.Const] = true
			seenValues[value.Value] = true
		}
		enums[ref] = enum
	}
	for _, transitions := range doc.Transitions {
		from, ok := enums[transitions.FromEnum]
		if !ok {
			return fmt.Errorf("transitions %s unknown from enum %s", transitions.Name, transitions.FromEnum)
		}
		to, ok := enums[transitions.ToEnum]
		if !ok {
			return fmt.Errorf("transitions %s unknown to enum %s", transitions.Name, transitions.ToEnum)
		}
		var event enumSpec
		if transitions.EventEnum != "" {
			var ok bool
			event, ok = enums[transitions.EventEnum]
			if !ok {
				return fmt.Errorf("transitions %s unknown event enum %s", transitions.Name, transitions.EventEnum)
			}
		}
		for _, entry := range transitions.Entries {
			if !from.hasConst(entry.From) {
				return fmt.Errorf("transitions %s unknown from const %s", transitions.Name, entry.From)
			}
			if !to.hasConst(entry.To) {
				return fmt.Errorf("transitions %s unknown to const %s", transitions.Name, entry.To)
			}
			if transitions.EventEnum != "" && !event.hasConst(entry.Event) {
				return fmt.Errorf("transitions %s unknown event const %s", transitions.Name, entry.Event)
			}
			if transitions.EventEnum == "" && entry.Event != "" {
				return fmt.Errorf("transitions %s does not declare event enum", transitions.Name)
			}
		}
	}
	return nil
}

func (e enumSpec) hasConst(name string) bool {
	for _, value := range e.Values {
		if value.Const == name {
			return true
		}
	}
	return false
}

func generatedFiles(doc document) (map[string][]byte, error) {
	files := map[string][]byte{}
	enums := map[string]enumSpec{}
	for _, enum := range doc.Enums {
		enums[enum.Package+"."+enum.Type] = enum
		body, err := generateEnumFile(enum)
		if err != nil {
			return nil, err
		}
		files[enumOutputPath(enum)] = body
	}
	for _, transitions := range doc.Transitions {
		body, err := generateTransitionFile(transitions, enums)
		if err != nil {
			return nil, err
		}
		files[transitionOutputPath(transitions)] = body
	}
	return files, nil
}

func enumOutputPath(enum enumSpec) string {
	return enum.Package + "/" + snake(enum.Type) + "_enum_gen.go"
}

func transitionOutputPath(transitions transitionSpec) string {
	return transitions.Package + "/" + snake(transitions.Name) + "_enum_gen.go"
}

func generateEnumFile(enum enumSpec) ([]byte, error) {
	var b bytes.Buffer
	writeHeader(&b, enum.Package)
	fmt.Fprintf(&b, "type %s uint16\n\n", enum.CodeType)
	fmt.Fprintf(&b, "type %s string\n\n", enum.StringType)

	fmt.Fprintln(&b, "const (")
	fmt.Fprintf(&b, "\t%sUnknown %s = iota\n", enum.CodeType, enum.CodeType)
	for _, value := range enum.Values {
		fmt.Fprintf(&b, "\t%s\n", enum.codeConst(value.Const))
	}
	fmt.Fprintln(&b, ")")
	fmt.Fprintln(&b)

	fmt.Fprintln(&b, "const (")
	for _, value := range enum.Values {
		fmt.Fprintf(&b, "\t%s %s = %q\n", enum.stringConst(value.Const), enum.StringType, value.Value)
	}
	fmt.Fprintln(&b, ")")
	fmt.Fprintln(&b)

	fmt.Fprintln(&b, "const (")
	for _, value := range enum.Values {
		fmt.Fprintf(&b, "\t%s %s = %s(%s)\n", value.Const, enum.Type, enum.Type, enum.stringConst(value.Const))
	}
	fmt.Fprintln(&b, ")")
	fmt.Fprintln(&b)

	fmt.Fprintf(&b, "func (value %s) Code() %s {\n", enum.Type, enum.CodeType)
	fmt.Fprintf(&b, "\treturn Parse%sCode(string(value))\n", enum.Type)
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)

	fmt.Fprintf(&b, "func Parse%sCode(value string) %s {\n", enum.Type, enum.CodeType)
	fmt.Fprintf(&b, "\tswitch %s(value) {\n", enum.StringType)
	for _, value := range enum.Values {
		fmt.Fprintf(&b, "\tcase %s:\n\t\treturn %s\n", enum.stringConst(value.Const), enum.codeConst(value.Const))
	}
	fmt.Fprintln(&b, "\tdefault:")
	fmt.Fprintf(&b, "\t\treturn %sUnknown\n", enum.CodeType)
	fmt.Fprintln(&b, "\t}")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)

	fmt.Fprintf(&b, "func (code %s) IsKnown() bool {\n", enum.CodeType)
	fmt.Fprintf(&b, "\treturn code >= %s && code <= %s\n", enum.codeConst(enum.Values[0].Const), enum.codeConst(enum.Values[len(enum.Values)-1].Const))
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)

	fmt.Fprintf(&b, "func (code %s) StringValue() %s {\n", enum.CodeType, enum.StringType)
	fmt.Fprintln(&b, "\tswitch code {")
	for _, value := range enum.Values {
		fmt.Fprintf(&b, "\tcase %s:\n\t\treturn %s\n", enum.codeConst(value.Const), enum.stringConst(value.Const))
	}
	fmt.Fprintln(&b, "\tdefault:")
	fmt.Fprintf(&b, "\t\treturn %s(\"\")\n", enum.StringType)
	fmt.Fprintln(&b, "\t}")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)

	fmt.Fprintf(&b, "func (code %s) String() string {\n", enum.CodeType)
	fmt.Fprintln(&b, "\treturn string(code.StringValue())")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)

	fmt.Fprintf(&b, "func (code %s) %s() %s {\n", enum.CodeType, enum.Type, enum.Type)
	fmt.Fprintf(&b, "\treturn %s(code.StringValue())\n", enum.Type)
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)

	fmt.Fprintf(&b, "func (value %s) Valid() bool {\n", enum.Type)
	fmt.Fprintln(&b, "\treturn value.Code().IsKnown()")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)

	fmt.Fprintf(&b, "func %s() []%s {\n", enum.CodeAllFunc, enum.CodeType)
	fmt.Fprintf(&b, "\treturn []%s{\n", enum.CodeType)
	for _, value := range enum.Values {
		fmt.Fprintf(&b, "\t\t%s,\n", enum.codeConst(value.Const))
	}
	fmt.Fprintln(&b, "\t}")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)

	fmt.Fprintf(&b, "func %s() []%s {\n", enum.AllFunc, enum.Type)
	fmt.Fprintf(&b, "\tcodes := %s()\n", enum.CodeAllFunc)
	fmt.Fprintf(&b, "\tout := make([]%s, len(codes))\n", enum.Type)
	fmt.Fprintln(&b, "\tfor index, code := range codes {")
	fmt.Fprintf(&b, "\t\tout[index] = code.%s()\n", enum.Type)
	fmt.Fprintln(&b, "\t}")
	fmt.Fprintln(&b, "\treturn out")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)

	writePredicate(&b, enum, "terminal", "IsTerminal")
	writePredicate(&b, enum, "active", "IsActive")
	writePredicate(&b, enum, "agent-active", "IsAgentActive")
	writePredicate(&b, enum, "transition", "IsTransition")
	writeNativeConfigRequirement(&b, enum)
	writePackagePredicate(&b, enum, "terminal", "IsTerminal")
	writePackagePredicate(&b, enum, "agent-active", "IsAgentActive")

	return formatSource(enumOutputPath(enum), b.Bytes())
}

func writePredicate(b *bytes.Buffer, enum enumSpec, attr, method string) {
	values := enum.valuesWithAttr(attr, "true")
	if len(values) == 0 {
		return
	}
	fmt.Fprintf(b, "func (code %s) %s() bool {\n", enum.CodeType, method)
	fmt.Fprintln(b, "\tswitch code {")
	writeCaseList(b, "\t", enumCodeRefs(enum, values))
	fmt.Fprintln(b, "\t\treturn true")
	fmt.Fprintln(b, "\tdefault:")
	fmt.Fprintln(b, "\t\treturn false")
	fmt.Fprintln(b, "\t}")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)

	fmt.Fprintf(b, "func (value %s) %s() bool {\n", enum.Type, method)
	fmt.Fprintf(b, "\treturn value.Code().%s()\n", method)
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
}

func writePackagePredicate(b *bytes.Buffer, enum enumSpec, attr, method string) {
	if enum.Package != "assignment" || len(enum.valuesWithAttr(attr, "true")) == 0 {
		return
	}
	fmt.Fprintf(b, "func %s(value %s) bool {\n", method, enum.Type)
	fmt.Fprintf(b, "\treturn value.Code().%s()\n", method)
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
}

func writeNativeConfigRequirement(b *bytes.Buffer, enum enumSpec) {
	if enum.Package != "ir" || enum.Type != "EventType" {
		return
	}
	groups := map[string][]enumValue{}
	for _, value := range enum.Values {
		requirement := value.Attrs["native-config"]
		if requirement == "" || requirement == "required" {
			continue
		}
		groups[requirement] = append(groups[requirement], value)
	}
	order := []struct {
		Key string
		Go  string
	}{
		{"pre-execute", "NativeConfigOptionalPreExecute"},
		{"phase-dependent", "NativeConfigPhaseDependent"},
		{"forbidden", "NativeConfigForbidden"},
	}
	fmt.Fprintf(b, "func (code %s) NativeConfigRequirement() NativeConfigRequirement {\n", enum.CodeType)
	fmt.Fprintln(b, "\tswitch code {")
	for _, item := range order {
		values := groups[item.Key]
		if len(values) == 0 {
			continue
		}
		writeCaseList(b, "\t", enumCodeRefs(enum, values))
		fmt.Fprintf(b, "\t\treturn %s\n", item.Go)
	}
	fmt.Fprintln(b, "\tdefault:")
	fmt.Fprintln(b, "\t\treturn NativeConfigRequired")
	fmt.Fprintln(b, "\t}")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)

	fmt.Fprintf(b, "func (value %s) NativeConfigRequirement() NativeConfigRequirement {\n", enum.Type)
	fmt.Fprintln(b, "\treturn value.Code().NativeConfigRequirement()")
	fmt.Fprintln(b, "}")
	fmt.Fprintln(b)
}

func generateTransitionFile(transitions transitionSpec, enums map[string]enumSpec) ([]byte, error) {
	from := enums[transitions.FromEnum]
	to := enums[transitions.ToEnum]
	event := enumSpec{}
	if transitions.EventEnum != "" {
		event = enums[transitions.EventEnum]
	}

	var b bytes.Buffer
	writeHeader(&b, transitions.Package)
	imports := transitionImports(transitions, from, to, event)
	if len(imports) > 0 {
		fmt.Fprintln(&b, "import (")
		for _, imp := range imports {
			fmt.Fprintf(&b, "\t%q\n", imp)
		}
		fmt.Fprintln(&b, ")")
		fmt.Fprintln(&b)
	}
	fmt.Fprintf(&b, "type %s struct {\n", transitions.Name)
	fmt.Fprintf(&b, "\tFrom %s\n", typeRef(from, transitions.Package))
	fmt.Fprintf(&b, "\tTo %s\n", typeRef(to, transitions.Package))
	if transitions.EventEnum != "" {
		fmt.Fprintf(&b, "\tTrigger %s\n", typeRef(event, transitions.Package))
	}
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)

	fmt.Fprintf(&b, "func %s() []%s {\n", transitions.AllFunc, transitions.Name)
	fmt.Fprintf(&b, "\treturn []%s{\n", transitions.Name)
	for _, entry := range transitions.Entries {
		fmt.Fprintf(&b, "\t\t{From: %s, To: %s", codeRef(from, transitions.Package, entry.From), codeRef(to, transitions.Package, entry.To))
		if transitions.EventEnum != "" {
			fmt.Fprintf(&b, ", Trigger: %s", codeRef(event, transitions.Package, entry.Event))
		}
		fmt.Fprintln(&b, "},")
	}
	fmt.Fprintln(&b, "\t}")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)

	fmt.Fprintf(&b, "func %s(from %s, to %s", transitions.Validate, typeRef(from, transitions.Package), typeRef(to, transitions.Package))
	if transitions.EventEnum != "" {
		fmt.Fprintf(&b, ", trigger %s", typeRef(event, transitions.Package))
	}
	fmt.Fprintln(&b, ") bool {")
	if transitions.AllowSame {
		fmt.Fprintln(&b, "\tif from == to && from.IsKnown() {")
		fmt.Fprintln(&b, "\t\treturn true")
		fmt.Fprintln(&b, "\t}")
	}
	fmt.Fprintln(&b, "\tswitch from {")
	byFrom := map[string][]transitionEntry{}
	for _, entry := range transitions.Entries {
		byFrom[entry.From] = append(byFrom[entry.From], entry)
	}
	fromKeys := make([]string, 0, len(byFrom))
	for key := range byFrom {
		fromKeys = append(fromKeys, key)
	}
	sort.SliceStable(fromKeys, func(i, j int) bool {
		return from.indexOfConst(fromKeys[i]) < from.indexOfConst(fromKeys[j])
	})
	for _, fromConst := range fromKeys {
		fmt.Fprintf(&b, "\tcase %s:\n", codeRef(from, transitions.Package, fromConst))
		fmt.Fprintln(&b, "\t\tswitch to {")
		groups := groupTransitionTargets(byFrom[fromConst], to, event)
		for _, group := range groups {
			fmt.Fprintf(&b, "\t\tcase %s:\n", codeRef(to, transitions.Package, group.To))
			if transitions.EventEnum == "" {
				fmt.Fprintln(&b, "\t\t\treturn true")
				continue
			}
			fmt.Fprint(&b, "\t\t\treturn ")
			for index, trigger := range group.Events {
				if index > 0 {
					fmt.Fprint(&b, " || ")
				}
				fmt.Fprintf(&b, "trigger == %s", codeRef(event, transitions.Package, trigger))
			}
			fmt.Fprintln(&b)
		}
		fmt.Fprintln(&b, "\t\t}")
	}
	fmt.Fprintln(&b, "\t}")
	fmt.Fprintln(&b, "\treturn false")
	fmt.Fprintln(&b, "}")
	fmt.Fprintln(&b)

	return formatSource(transitionOutputPath(transitions), b.Bytes())
}

type transitionTargetGroup struct {
	To     string
	Events []string
}

func groupTransitionTargets(entries []transitionEntry, to, event enumSpec) []transitionTargetGroup {
	byTo := map[string][]string{}
	for _, entry := range entries {
		byTo[entry.To] = append(byTo[entry.To], entry.Event)
	}
	toKeys := make([]string, 0, len(byTo))
	for key := range byTo {
		toKeys = append(toKeys, key)
	}
	sort.SliceStable(toKeys, func(i, j int) bool {
		return to.indexOfConst(toKeys[i]) < to.indexOfConst(toKeys[j])
	})
	groups := make([]transitionTargetGroup, 0, len(toKeys))
	for _, toConst := range toKeys {
		events := byTo[toConst]
		if event.Type != "" {
			sort.SliceStable(events, func(i, j int) bool {
				return event.indexOfConst(events[i]) < event.indexOfConst(events[j])
			})
		}
		groups = append(groups, transitionTargetGroup{To: toConst, Events: events})
	}
	return groups
}

func transitionImports(transitions transitionSpec, enums ...enumSpec) []string {
	seen := map[string]bool{}
	var imports []string
	for _, enum := range enums {
		if enum.Package == "" || enum.Package == transitions.Package {
			continue
		}
		path := modulePath + "/" + enum.Package
		if !seen[path] {
			seen[path] = true
			imports = append(imports, path)
		}
	}
	sort.Strings(imports)
	return imports
}

func typeRef(enum enumSpec, currentPackage string) string {
	if enum.Package == currentPackage {
		return enum.CodeType
	}
	return enum.Package + "." + enum.CodeType
}

func codeRef(enum enumSpec, currentPackage, constName string) string {
	name := enum.codeConst(constName)
	if enum.Package == currentPackage {
		return name
	}
	return enum.Package + "." + name
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

func (e enumSpec) suffix(constName string) string {
	if e.ConstPrefix != "" && strings.HasPrefix(constName, e.ConstPrefix) {
		suffix := strings.TrimPrefix(constName, e.ConstPrefix)
		if suffix != "" {
			return suffix
		}
	}
	return constName
}

func (e enumSpec) codeConst(constName string) string {
	return e.CodeType + e.suffix(constName)
}

func (e enumSpec) stringConst(constName string) string {
	return e.StringType + e.suffix(constName)
}

func (e enumSpec) valuesWithAttr(attr, want string) []enumValue {
	var out []enumValue
	for _, value := range e.Values {
		if value.Attrs[attr] == want {
			out = append(out, value)
		}
	}
	return out
}

func enumCodeRefs(enum enumSpec, values []enumValue) []string {
	out := make([]string, 0, len(values))
	for _, value := range values {
		out = append(out, enum.codeConst(value.Const))
	}
	return out
}

func writeCaseList(b *bytes.Buffer, indent string, refs []string) {
	fmt.Fprintf(b, "%scase ", indent)
	for index, ref := range refs {
		if index > 0 {
			fmt.Fprint(b, ", ")
		}
		fmt.Fprint(b, ref)
	}
	fmt.Fprintln(b, ":")
}

func (e enumSpec) indexOfConst(name string) int {
	for index, value := range e.Values {
		if value.Const == name {
			return index
		}
	}
	return len(e.Values)
}

func snake(value string) string {
	var b strings.Builder
	var prevLower bool
	for index, r := range value {
		if index > 0 && unicode.IsUpper(r) && prevLower {
			b.WriteByte('_')
		}
		b.WriteRune(unicode.ToLower(r))
		prevLower = unicode.IsLower(r) || unicode.IsDigit(r)
	}
	return b.String()
}
