package main

import (
	"fmt"
	"strings"
)

func verifyAnnotationInventoryItem(policy annotationContentPolicy, item annotationInventory) error {
	if blank(item.UIArea) || blank(item.FigmaGeneratedPath) || blank(item.CanonicalGeneratedPath) {
		return fmt.Errorf("annotation inventory item requires ui_area and generated paths")
	}
	if item.CategoryID != policy.CategoryID || item.CategoryLabel != policy.CategoryLabel {
		return fmt.Errorf("annotation %s category mismatch", item.FigmaGeneratedPath)
	}
	if !strings.HasPrefix(item.FigmaGeneratedPath, "riido.v2.") {
		return fmt.Errorf("annotation %s must use riido.v2 facade path", item.FigmaGeneratedPath)
	}
	if !strings.HasPrefix(item.CanonicalGeneratedPath, "aiAgent.") {
		return fmt.Errorf("annotation %s must use canonical aiAgent path", item.FigmaGeneratedPath)
	}
	if !validOperationKind(item.OperationKind) {
		return fmt.Errorf("annotation %s unsupported operation_kind %q", item.FigmaGeneratedPath, item.OperationKind)
	}
	if blank(item.Background) || item.AnnotationCount <= 0 || len(item.Sources) == 0 {
		return fmt.Errorf("annotation %s requires background, positive count, and sources", item.FigmaGeneratedPath)
	}
	return nil
}
