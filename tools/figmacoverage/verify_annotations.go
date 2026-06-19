package main

import (
	"errors"
	"fmt"
)

func verifyAnnotationPolicy(policy annotationContentPolicy) error {
	if blank(policy.CategoryID) || blank(policy.CategoryLabel) || blank(policy.Rule) {
		return errors.New("api annotation category_id, category_label, and rule are required")
	}
	if len(policy.LabelFormat) != 3 {
		return fmt.Errorf("label_format length=%d, want 3", len(policy.LabelFormat))
	}
	if policy.LiveInspection.TotalAPIGeneratedAnnotations <= 0 {
		return errors.New("live inspection total_api_generated_annotations must be positive")
	}
	return nil
}

func verifyAnnotationInventory(policy annotationContentPolicy, inventory []annotationInventory) error {
	if len(inventory) == 0 {
		return errors.New("api_generated_annotation_inventory is required")
	}
	total := 0
	for _, item := range inventory {
		if err := verifyAnnotationInventoryItem(policy, item); err != nil {
			return err
		}
		total += item.AnnotationCount
	}
	if total != policy.LiveInspection.TotalAPIGeneratedAnnotations {
		return fmt.Errorf("annotation_count sum=%d total_api_generated_annotations=%d", total, policy.LiveInspection.TotalAPIGeneratedAnnotations)
	}
	return nil
}

func validOperationKind(kind string) bool {
	return kind == "Query" || kind == "Mutation" || kind == "SSE Stream"
}
