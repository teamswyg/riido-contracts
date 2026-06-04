// Package progressmessage owns the AI Agent runtime progress message catalog.
//
// The catalog is intentionally small, integer-coded, append-only, and rendered
// before public SSE delivery so frontend clients can keep consuming the
// existing message string surface.
package progressmessage
