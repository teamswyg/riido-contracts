package main

func typeRef(enum enumSpec, currentPackage string) string {
	if enum.Package == currentPackage {
		return enum.CodeType
	}
	return enum.Package + "." + enum.CodeType
}

func codeRef(enum enumSpec, currentPackage, constName string) string {
	name := enum.codeConst(constName)
	if enum.Package == currentPackage {
		return name
	}
	return enum.Package + "." + name
}
