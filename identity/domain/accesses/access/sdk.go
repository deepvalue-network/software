package access

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/products/libs/cryptography/pk/encryption"
	"github.com/steve-care-software/products/libs/cryptography/pk/signature"
)

// NewFactory creates a new factory instance
func NewFactory(encPkBitrate int) Factory {
	sigPkFactory := signature.NewPrivateKeyFactory()
	encPkFactory := encryption.NewFactory(encPkBitrate)
	builder := NewBuilder()
	return createFactory(encPkFactory, sigPkFactory, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewPointer returns a new access pointer
func NewPointer() interface{} {
	return new(access)
}

// Factory represents an access factory
type Factory interface {
	Create() (Access, error)
}

// Builder represents an access builder
type Builder interface {
	Create() Builder
	WithID(id *uuid.UUID) Builder
	WithSignature(sigPK signature.PrivateKey) Builder
	WithEncryption(encPK encryption.PrivateKey) Builder
	Now() (Access, error)
}

// Access represents a user access
type Access interface {
	ID() *uuid.UUID
	Signature() signature.PrivateKey
	Encryption() encryption.PrivateKey
}
