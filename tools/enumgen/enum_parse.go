package main

import (
	"errors"
	"fmt"
	"strings"
)

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
	return validateEnumSpec(enumSpec{
		Package:     props["package"],
		Type:        props["type"],
		CodeType:    props["code-type"],
		StringType:  props["string-type"],
		ConstPrefix: props["const-prefix"],
		AllFunc:     props["all"],
		CodeAllFunc: props["code-all"],
		Values:      values,
	})
}

func validateEnumSpec(spec enumSpec) (enumSpec, error) {
	if spec.Package == "" || spec.Type == "" || spec.CodeType == "" ||
		spec.StringType == "" || spec.AllFunc == "" || spec.CodeAllFunc == "" {
		return enumSpec{}, fmt.Errorf("enum %q is missing required properties", spec.Type)
	}
	if len(spec.Values) == 0 {
		return enumSpec{}, fmt.Errorf("enum %s has no values", spec.Type)
	}
	return spec, nil
}
