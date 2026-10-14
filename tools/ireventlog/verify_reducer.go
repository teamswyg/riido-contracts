package main

import (
	"fmt"
	"strings"
)

func verifyReducerSurface(model model) error {
	want := []string{"LastEventID", "CurrentState", "Diagnostics", "Error"}
	if len(model.ReduceResultFieldNames) != len(want) {
		return fmt.Errorf("ReduceResult fields = %v", model.ReduceResultFieldNames)
	}
	for i, field := range want {
		if model.ReduceResultFieldNames[i] != field {
			return fmt.Errorf("ReduceResult field %d = %s, want %s", i, model.ReduceResultFieldNames[i], field)
		}
	}
	for _, field := range model.ReduceResultFieldNames {
		if strings.Contains(field, "Append") || strings.Contains(field, "Writer") {
			return fmt.Errorf("ReduceResult exposes side-effect field %s", field)
		}
	}
	return nil
}
