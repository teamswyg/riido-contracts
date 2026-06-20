package progressmessage

import "embed"

//go:embed catalog.ir.riido.json catalog.ir.messages/*.riido.json
var embeddedCatalog embed.FS
