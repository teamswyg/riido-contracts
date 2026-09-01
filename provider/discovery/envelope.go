package discovery

import (
	"bytes"
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"
)

func Sign(policy Policy, keyID string, privateKey ed25519.PrivateKey) (SignedEnvelope, error) {
	if err := policy.Validate(); err != nil {
		return SignedEnvelope{}, err
	}
	if !boundedID(keyID, 64) || len(privateKey) != ed25519.PrivateKeySize {
		return SignedEnvelope{}, errors.New("provider discovery: invalid signing key")
	}
	payload, err := json.Marshal(policy)
	if err != nil {
		return SignedEnvelope{}, fmt.Errorf("provider discovery: encode policy: %w", err)
	}
	return SignedEnvelope{
		SchemaVersion: EnvelopeSchemaVersion,
		KeyID:         keyID,
		Algorithm:     SignatureAlgorithm,
		Payload:       base64.StdEncoding.EncodeToString(payload),
		Signature:     base64.StdEncoding.EncodeToString(ed25519.Sign(privateKey, payload)),
	}, nil
}

func (e SignedEnvelope) Verify(publicKeys map[string]ed25519.PublicKey, now time.Time) (Policy, error) {
	if err := e.Validate(); err != nil {
		return Policy{}, err
	}
	payload, _ := base64.StdEncoding.DecodeString(e.Payload)
	signature, _ := base64.StdEncoding.DecodeString(e.Signature)
	publicKey := publicKeys[e.KeyID]
	if len(publicKey) != ed25519.PublicKeySize || !ed25519.Verify(publicKey, payload, signature) {
		return Policy{}, errors.New("provider discovery: signature verification failed")
	}
	decoder := json.NewDecoder(bytes.NewReader(payload))
	decoder.DisallowUnknownFields()
	var policy Policy
	if err := decoder.Decode(&policy); err != nil {
		return Policy{}, fmt.Errorf("provider discovery: decode policy: %w", err)
	}
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return Policy{}, errors.New("provider discovery: trailing policy data")
	}
	if err := policy.ValidateAt(now); err != nil {
		return Policy{}, err
	}
	return policy, nil
}
