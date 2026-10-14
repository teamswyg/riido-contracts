package main

func streamVariants(doc apiFixture) []string {
	for _, sumType := range doc.SumTypes {
		if sumType.Name == "ClientStreamEvent" {
			out := make([]string, 0, len(sumType.Variants))
			for _, variant := range sumType.Variants {
				out = append(out, variant.Kind)
			}
			return out
		}
	}
	return nil
}

func hasRequiredVariants(got, want []string) bool {
	gotSet := map[string]bool{}
	for _, value := range got {
		gotSet[value] = true
	}
	for _, value := range want {
		if !gotSet[value] {
			return false
		}
	}
	return true
}
