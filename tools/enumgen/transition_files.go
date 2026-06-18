package main

import (
	"bytes"
	"fmt"
	"strings"
)

type transitionModel struct {
	Spec  transitionSpec
	From  enumSpec
	To    enumSpec
	Event enumSpec
}

type transitionFileSection struct {
	Name  string
	Write func(*bytes.Buffer, transitionModel)
}

func generateTransitionFiles(
	transitions transitionSpec,
	enums map[string]enumSpec,
) (map[string][]byte, error) {
	model := transitionModel{Spec: transitions, From: enums[transitions.FromEnum], To: enums[transitions.ToEnum]}
	if transitions.EventEnum != "" {
		model.Event = enums[transitions.EventEnum]
	}
	files := map[string][]byte{}
	for _, section := range transitionFileSections() {
		name := transitionOutputPath(transitions, section.Name)
		body, err := formatTransitionSection(name, model, section.Write)
		if err != nil {
			return nil, err
		}
		files[name] = body
	}
	return files, nil
}

func transitionFileSections() []transitionFileSection {
	return []transitionFileSection{
		{"types", writeTransitionTypes},
		{"all", writeTransitionAll},
		{"allowed", writeTransitionAllowed},
		{"validate", writeTransitionValidate},
	}
}

func transitionOutputPath(transitions transitionSpec, section string) string {
	return fmt.Sprintf("%s/%s_enum_%s_gen.go", transitions.Package, snake(transitions.Name), section)
}

func transitionPrivateName(transitions transitionSpec, suffix string) string {
	return strings.ToLower(transitions.Name[:1]) + transitions.Name[1:] + suffix
}
