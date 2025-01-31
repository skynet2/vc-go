/*
Copyright Gen Digital Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package api

import (
	"github.com/trustbloc/did-go/doc/ld/processor"
	"github.com/trustbloc/kms-go/doc/jose/jwk"
)

// SignatureSuite encapsulates signature suite methods required for signature verification.
type SignatureSuite interface {

	// GetCanonicalDocument will return normalized/canonical version of the document
	GetCanonicalDocument(doc map[string]interface{}, opts ...processor.Opts) ([]byte, error)

	// GetDigest returns document digest
	GetDigest(doc []byte) []byte

	// Verify will verify signature against public key
	Verify(pubKey *PublicKey, doc []byte, signature []byte) error

	// Accept registers this signature suite with the given signature type
	Accept(signatureType string) bool

	// CompactProof indicates weather to compact the proof doc before canonization
	CompactProof() bool
}

// PublicKey contains a result of public key resolution.
type PublicKey struct {
	Type  string
	Value []byte
	JWK   *jwk.JWK
}
