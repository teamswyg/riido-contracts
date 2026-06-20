package progressmessage

func Catalog() (IRDocument, error) {
	return LoadIR(embeddedCatalog, "catalog.ir.riido.json")
}
