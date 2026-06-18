package main

import "fmt"

func validateDocument(doc document) error {
	enums := map[string]enumSpec{}
	for _, enum := range doc.Enums {
		ref := enum.Package + "." + enum.Type
		if _, ok := enums[ref]; ok {
			return fmt.Errorf("duplicate enum %s", ref)
		}
		if err := validateEnumValues(ref, enum.Values); err != nil {
			return err
		}
		enums[ref] = enum
	}
	for _, transitions := range doc.Transitions {
		if err := validateTransitionRefs(transitions, enums); err != nil {
			return err
		}
	}
	return nil
}

func validateEnumValues(ref string, values []enumValue) error {
	seenConsts := map[string]bool{}
	seenValues := map[string]bool{}
	for _, value := range values {
		if value.Const == "" || value.Value == "" {
			return fmt.Errorf("enum %s has empty const or value", ref)
		}
		if seenConsts[value.Const] {
			return fmt.Errorf("enum %s duplicate const %s", ref, value.Const)
		}
		if seenValues[value.Value] {
			return fmt.Errorf("enum %s duplicate value %s", ref, value.Value)
		}
		seenConsts[value.Const] = true
		seenValues[value.Value] = true
	}
	return nil
}
