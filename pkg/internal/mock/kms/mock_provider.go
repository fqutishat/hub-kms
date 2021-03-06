/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package kms

import (
	"github.com/hyperledger/aries-framework-go/pkg/crypto"
	"github.com/hyperledger/aries-framework-go/pkg/kms"
	mockcrypto "github.com/hyperledger/aries-framework-go/pkg/mock/crypto"
	mockkms "github.com/hyperledger/aries-framework-go/pkg/mock/kms"

	mockkeystore "github.com/trustbloc/hub-kms/pkg/internal/mock/keystore"
	"github.com/trustbloc/hub-kms/pkg/keystore"
)

// MockProvider is a mock Provider for KMS service.
type MockProvider struct {
	MockKeystore   *mockkeystore.MockRepository
	MockKeyManager *mockkms.KeyManager
	MockCrypto     *mockcrypto.Crypto
}

// NewMockProvider returns a new mock Provider for KMS service.
func NewMockProvider() *MockProvider {
	return &MockProvider{
		MockKeystore:   mockkeystore.NewMockRepository(),
		MockKeyManager: &mockkms.KeyManager{},
		MockCrypto:     &mockcrypto.Crypto{},
	}
}

// Keystore gets the keystore repository instance.
func (p *MockProvider) Keystore() keystore.Repository {
	return p.MockKeystore
}

// KeyManager gets the KeyManager instance.
func (p *MockProvider) KeyManager() kms.KeyManager {
	return p.MockKeyManager
}

// Crypto gets the Crypto instance.
func (p *MockProvider) Crypto() crypto.Crypto {
	return p.MockCrypto
}
