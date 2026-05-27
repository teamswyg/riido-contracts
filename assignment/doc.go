// Package assignment owns the shared C10 SaaS assignment polling contract.
//
// The package is intentionally limited to DTOs, schema identifiers, enum
// values, task event type values, and pure transition predicates that must be
// shared by riido-daemon and riido-control-plane. Store actors, HTTP handlers,
// SSE fan-out, metrics routes, health routes, authorization, provider process
// execution, DynamoDB adapters, Terraform, and deployment evidence stay in the
// runtime repositories that own those behaviors.
package assignment
