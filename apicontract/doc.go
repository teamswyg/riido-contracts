// Package apicontract owns shared API contract projection fixtures.
//
// The API contract flow is:
//
//	Domain DSL -> canonical API IR -> OpenAPI projection
//
// The DSL and IR are the contract truth sources. OpenAPI is a generated
// adapter-facing projection for web clients, HTTP test servers, and black-box HTTP
// tests. This package does not own runtime HTTP handlers, authorization
// implementation, stores, persistence, or deployment configuration.
package apicontract
