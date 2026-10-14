package main

import "strings"

type patternSection struct {
	Name     string
	Template string
}

func patternSections() []patternSection {
	return []patternSection{
		{"types", "tools/fsmgen/templates/fsm_pattern_types.go.gotmpl"},
		{"code_consts", "tools/fsmgen/templates/fsm_pattern_code_consts.go.gotmpl"},
		{"string_consts", "tools/fsmgen/templates/fsm_pattern_string_consts.go.gotmpl"},
		{"domain_consts", "tools/fsmgen/templates/fsm_pattern_domain_consts.go.gotmpl"},
		{"parse_map", "tools/fsmgen/templates/fsm_pattern_parse_map.go.gotmpl"},
		{"parse", "tools/fsmgen/templates/fsm_pattern_parse.go.gotmpl"},
		{"string_map", "tools/fsmgen/templates/fsm_pattern_string_map.go.gotmpl"},
		{"code_methods", "tools/fsmgen/templates/fsm_pattern_code_methods.go.gotmpl"},
		{"all_codes", "tools/fsmgen/templates/fsm_pattern_all_codes.go.gotmpl"},
		{"all", "tools/fsmgen/templates/fsm_pattern_all.go.gotmpl"},
	}
}

func patternSectionOutputPath(outputPath, section string) string {
	base := strings.TrimSuffix(outputPath, "_gen.go")
	return base + "_" + section + "_gen.go"
}
