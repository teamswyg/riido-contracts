package main

import (
	"errors"
	"fmt"
	"strings"
)

func parsePatternSumType(form node) (patternSumType, error) {
	props := map[string]string{}
	values := []patternValue{}
	for i := 1; i < len(form.list); i++ {
		item := form.list[i]
		if item.isAtom() && strings.HasPrefix(item.atom, ":") {
			if i+1 >= len(form.list) {
				return patternSumType{}, fmt.Errorf("sum-type property %s missing value", item.atom)
			}
			props[strings.TrimPrefix(item.atom, ":")] = atom(form.list[i+1])
			i++
			continue
		}
		value, err := parsePatternVariant(item)
		if err != nil {
			return patternSumType{}, err
		}
		values = append(values, value)
	}
	return patternSumTypeFromProps(props, values)
}

func parsePatternVariant(form node) (patternValue, error) {
	if form.isAtom() || len(form.list) == 0 || atom(form.list[0]) != "variant" {
		return patternValue{}, errors.New("sum-type entries must be (variant ...)")
	}
	if len(form.list) != 3 {
		return patternValue{}, errors.New("variant requires const and string")
	}
	return patternValue{Const: atom(form.list[1]), Value: atom(form.list[2])}, nil
}
