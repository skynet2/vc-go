/*
Copyright Gen Digital Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package jwt

import (
	"github.com/trustbloc/kms-go/doc/jose"
)

// NewJOSESigner wraps ProofCreator into jose signer.
func NewJOSESigner(params SignParameters, signer ProofCreator) (*JoseSigner, error) {
	headers, err := signer.CreateJWTHeaders(params)
	if err != nil {
		return nil, err
	}

	return &JoseSigner{
		signer:     signer,
		signParams: params,
		headers:    headers,
	}, nil
}

// JoseSigner implement jose.proofCreator interface.
type JoseSigner struct {
	signer     ProofCreator
	signParams SignParameters
	headers    jose.Headers
}

// Sign returns signature.
func (s JoseSigner) Sign(data []byte) ([]byte, error) {
	return s.signer.SignJWT(s.signParams, data)
}

// Headers returns headers.
func (s JoseSigner) Headers() jose.Headers {
	return s.headers
}

type joseVerifier struct {
	proofChecker ProofChecker
}

func (v *joseVerifier) Verify(joseHeaders jose.Headers, payload, signingInput, signature []byte) error {
	return v.proofChecker.CheckJWTProof(joseHeaders, payload, signingInput, signature)
}