package main

import "strings"

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
