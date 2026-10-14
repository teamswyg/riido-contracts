package capability

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

// ComputeCapabilityFingerprint produces the SHA-256 hex of the canonical
// JSON serialization of input. Same input from any daemon must yield the same
// fingerprint.
func ComputeCapabilityFingerprint(in CapabilityFingerprintInput) (CapabilityFingerprint, error) {
	payload, err := newFingerprintPayload(in)
	if err != nil {
		return "", err
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("marshal payload: %w", err)
	}
	sum := sha256.Sum256(b)
	return CapabilityFingerprint(hex.EncodeToString(sum[:])), nil
}
