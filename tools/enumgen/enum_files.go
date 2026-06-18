package main

import (
	"bytes"
	"fmt"
	"strings"
)

type enumFileSection struct {
	Name  string
	Write func(*bytes.Buffer, enumSpec)
}

func generateEnumFiles(enum enumSpec) (map[string][]byte, error) {
	files := map[string][]byte{}
	for _, section := range enumFileSections(enum) {
		name := enumOutputPath(enum, section.Name)
		body, err := formatEnumSection(name, enum, section.Write)
		if err != nil {
			return nil, err
		}
		files[name] = body
	}
	return files, nil
}

func enumFileSections(enum enumSpec) []enumFileSection {
	sections := []enumFileSection{
		{"types", writeEnumTypes},
		{"code_consts", writeEnumCodeConsts},
		{"string_consts", writeEnumStringConsts},
		{"domain_consts", writeEnumDomainConsts},
		{"parse_map", writeEnumParseMap},
		{"parse", writeEnumParseMethods},
		{"string_map", writeEnumStringMap},
		{"code_methods", writeEnumCodeMethods},
		{"domain_methods", writeEnumDomainMethods},
		{"all_codes", writeEnumAllCodeValues},
		{"all", writeEnumAllMethods},
	}
	if enumHasPredicateSection(enum) {
		sections = append(sections, enumFileSection{"predicates", writeEnumPredicates})
	}
	return sections
}

func enumOutputPath(enum enumSpec, section string) string {
	return fmt.Sprintf("%s/%s_enum_%s_gen.go", enum.Package, snake(enum.Type), section)
}

func formatEnumSection(
	name string,
	enum enumSpec,
	write func(*bytes.Buffer, enumSpec),
) ([]byte, error) {
	var b bytes.Buffer
	writeHeader(&b, enum.Package)
	write(&b, enum)
	return formatSource(name, b.Bytes())
}

func enumPrivateName(enum enumSpec, suffix string) string {
	name := enum.Type[:1] + enum.Type[1:]
	return strings.ToLower(name[:1]) + name[1:] + suffix
}
