// Package discovery defines the signed control-plane wire contract for
// provider executable discovery hints.
//
// The contract never carries customer absolute paths or commands. A daemon
// resolves bounded roots on its own host and remains responsible for checking
// that a candidate is an executable supported by its local provider adapter.
package discovery
