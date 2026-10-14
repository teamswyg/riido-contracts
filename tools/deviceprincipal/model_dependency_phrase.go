package main

func factHasPhrase(fact dependencyFact, phrase string) bool {
	for _, ref := range fact.SourceRef {
		if ref.Repo == "riido-contracts" && ref.RequiredPhrase == phrase {
			return true
		}
	}
	return false
}
