package main

import (
	"errors"
	"fmt"
)

func patternSumTypeFromProps(props map[string]string, values []patternValue) (patternSumType, error) {
	sumType := patternSumType{
		Package:     props["package"],
		Type:        props["type"],
		CodeType:    props["code-type"],
		StringType:  props["string-type"],
		ConstPrefix: props["const-prefix"],
		OutputPath:  props["output"],
		Values:      values,
	}
	if sumType.Package == "" || sumType.Type == "" || sumType.CodeType == "" || sumType.StringType == "" || sumType.OutputPath == "" {
		return patternSumType{}, errors.New("sum-type missing package, type, code-type, string-type, or output")
	}
	if len(sumType.Values) == 0 {
		return patternSumType{}, fmt.Errorf("sum-type %s has no variants", sumType.Type)
	}
	outputPath, err := cleanRepoRelativePath(sumType.OutputPath)
	if err != nil {
		return patternSumType{}, fmt.Errorf("sum-type %s output: %w", sumType.Type, err)
	}
	sumType.OutputPath = outputPath
	return sumType, nil
}
