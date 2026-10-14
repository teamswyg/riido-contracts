package main

import "fmt"

func verifySourceRefs(root string, refs []sourceRef) error {
	seen := map[string]bool{}
	for _, ref := range refs {
		key := ref.Repo + ":" + ref.Path + ":" + ref.RequiredPhrase
		if seen[key] {
			return fmt.Errorf("duplicate source_ref %q", key)
		}
		seen[key] = true
		if err := verifySourceRef(root, ref); err != nil {
			return err
		}
	}
	return nil
}
