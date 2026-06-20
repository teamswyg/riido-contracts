package main

func generate() error {
	root, err := resolveRoot(".")
	if err != nil {
		return err
	}
	body, err := generatedIR(root)
	if err != nil {
		return err
	}
	return writeFile(resolve(root, irPath), body)
}
