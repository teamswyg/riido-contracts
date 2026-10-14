// Package catalog owns the shared provider-kind vocabulary and default model
// suppression rules used by Riido runtime integrations.
//
// This package intentionally does not select adapters. Adapter selection is
// still governed by provider/capability.ProtocolKind and daemon-local runtime
// wiring.
package catalog
