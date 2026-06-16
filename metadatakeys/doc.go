// Package metadatakeys owns shared string keys used in Riido runtime metadata
// maps.
//
// The values are intentionally plain strings at storage boundaries so existing
// DynamoDB and JSON records remain compatible. Code should consume these names
// through this vocabulary instead of redefining literals in each layer.
package metadatakeys
