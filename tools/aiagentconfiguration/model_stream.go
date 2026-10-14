package main

func hasStreamVariant(sumTypes []sumType, variant string) bool {
	for _, sum := range sumTypes {
		if sum.Name != "ClientStreamEvent" {
			continue
		}
		for _, item := range sum.Variants {
			if item.Kind == variant {
				return true
			}
		}
	}
	return false
}
