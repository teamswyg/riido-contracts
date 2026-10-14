package main

import capabilitypkg "github.com/teamswyg/riido-contracts/provider/capability"

func protocolRows() []protocolRow {
	protocols := capabilitypkg.AllProtocolKinds()
	out := make([]protocolRow, len(protocols))
	for i, protocol := range protocols {
		out[i] = protocolRow{
			Kind: string(protocol),
			Args: capabilitypkg.ProtocolCriticalArgs(protocol),
		}
	}
	return out
}

func criticalArgSetCount() int {
	count := 0
	for _, row := range protocolRows() {
		if len(row.Args) > 0 {
			count++
		}
	}
	return count
}
