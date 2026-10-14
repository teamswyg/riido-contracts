package main

func generatedFiles(doc document) (map[string][]byte, error) {
	files := map[string][]byte{}
	enums := map[string]enumSpec{}
	for _, enum := range doc.Enums {
		enums[enum.Package+"."+enum.Type] = enum
		enumFiles, err := generateEnumFiles(enum)
		if err != nil {
			return nil, err
		}
		for name, body := range enumFiles {
			files[name] = body
		}
	}
	for _, transitions := range doc.Transitions {
		transitionFiles, err := generateTransitionFiles(transitions, enums)
		if err != nil {
			return nil, err
		}
		for name, body := range transitionFiles {
			files[name] = body
		}
	}
	return files, nil
}
