package main

import "sort"

func transitionImports(transitions transitionSpec, enums ...enumSpec) []string {
	seen := map[string]bool{}
	var imports []string
	for _, enum := range enums {
		if enum.Package == "" || enum.Package == transitions.Package {
			continue
		}
		path := modulePath + "/" + enum.Package
		if !seen[path] {
			seen[path] = true
			imports = append(imports, path)
		}
	}
	sort.Strings(imports)
	return imports
}
