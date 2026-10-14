package main

import "strings"

type patternTemplateValue struct {
	Const       string `json:"Const"`
	Value       string `json:"Value"`
	CodeConst   string `json:"CodeConst"`
	StringConst string `json:"StringConst"`
}

type patternTemplateData struct {
	SourcePath     string                 `json:"SourcePath"`
	Package        string                 `json:"Package"`
	Type           string                 `json:"Type"`
	PrivateName    string                 `json:"PrivateName"`
	CodeType       string                 `json:"CodeType"`
	StringType     string                 `json:"StringType"`
	Values         []patternTemplateValue `json:"Values"`
	FirstCodeConst string                 `json:"FirstCodeConst"`
	LastCodeConst  string                 `json:"LastCodeConst"`
}

func patternTemplateDataFrom(source string, patterns patternDocument) patternTemplateData {
	data := patternTemplateData{
		SourcePath:  source,
		Package:     patterns.SumType.Package,
		Type:        patterns.SumType.Type,
		PrivateName: lowerFirst(patterns.SumType.Type),
		CodeType:    patterns.SumType.CodeType,
		StringType:  patterns.SumType.StringType,
	}
	for _, value := range patterns.SumType.Values {
		data.Values = append(data.Values, patternTemplateValueFrom(patterns.SumType, value))
	}
	data.FirstCodeConst = data.Values[0].CodeConst
	data.LastCodeConst = data.Values[len(data.Values)-1].CodeConst
	return data
}

func patternTemplateValueFrom(sumType patternSumType, value patternValue) patternTemplateValue {
	suffix := patternConstSuffix(sumType, value.Const)
	return patternTemplateValue{
		Const:       value.Const,
		Value:       value.Value,
		CodeConst:   sumType.CodeType + suffix,
		StringConst: sumType.StringType + suffix,
	}
}

func lowerFirst(value string) string {
	if value == "" {
		return value
	}
	return strings.ToLower(value[:1]) + value[1:]
}
