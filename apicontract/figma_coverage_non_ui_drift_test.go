package apicontract

import (
	"strconv"
	"strings"
)

func figmaNonUIInventoryDriftDocumented(limitations []figmaSupportingToolLimitation, pageID string, knownInventoryCount, childCount int) bool {
	for _, limitation := range limitations {
		if limitation.ID != "figma-onboarding-page-load-timeout.v1" {
			continue
		}
		if !stringSliceContains(limitation.AuthoritativeResult, pageID) {
			continue
		}
		if !stringSliceContains(limitation.AuthoritativeResult, "child_count="+strconv.Itoa(childCount)) {
			continue
		}
		if !stringSliceContains(limitation.AuthoritativeResult, "known_inventory_count="+strconv.Itoa(knownInventoryCount)) {
			continue
		}
		if !stringSliceContains(limitation.AuthoritativeResult, "unresolved_extra_top_level_node="+strconv.Itoa(childCount-knownInventoryCount)) {
			continue
		}
		return strings.Contains(strings.ToLower(limitation.Rule), "known_inventory_count may lag expected_pages.child_count")
	}
	return false
}
