// Package contracts anchors shared Riido contract identifiers.
//
//go:generate go run ./tools/enumgen generate
//go:generate go run ./tools/fsmgen generate
package contracts

const (
	ModulePath         = "github.com/teamswyg/riido-contracts"
	ContractSetVersion = "riido-contracts.v0"
)
