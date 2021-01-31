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
	return createFactory(encPkFactory, sigPkFactory)
}

// Factory represents an access factory
type Factory interface {
	Create() (Access, error)
}

// Access represents a user access
type Access interface {
	ID() *uuid.UUID
	Signature() signature.PrivateKey
	Encryption() encryption.PrivateKey
}
