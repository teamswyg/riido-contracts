package main

import (
	"errors"
	"fmt"
	"strings"
)

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
