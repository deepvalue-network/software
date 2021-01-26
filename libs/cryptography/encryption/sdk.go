package encryption

import (
	"go.dedis.ch/kyber/v3/group/edwards25519"
)

var curve = edwards25519.NewBlakeSHA256Ed25519()

// NewEncryption creates a new encryption instance
func NewEncryption(pass string) Encryption {
	hasher := curve.Hash()
	hasher.Write([]byte(pass))
	return createEncryption(hasher.Sum(nil))
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the builder
type Builder interface {
	Create() Builder
	WithPassword(password []byte) Builder
	Now() (Encryption, error)
}

// Encryption represents an encryption
type Encryption interface {
	Encrypt(message []byte) (string, error)
	Decrypt(encryptedText string) ([]byte, error)
}
