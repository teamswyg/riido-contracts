// Package runstate owns the provider-neutral run-scope sub-state vocabulary.
//
// The daemon owns the runtime reducer that moves through these states. This
// package only owns the shared names, integer codes, and terminal predicate used
// for conformance between generated contracts and runtime implementations.
package runstate
