package public

import (
	"crypto/rsa"

	"github.com/deepvalue-network/software/libs/hash"
)

// NewAdapter returns a new encryption's public adapter
func NewAdapter() Adapter {
	hashAdapter := hash.NewAdapter()
	builder := NewBuilder()
	return createAdapter(hashAdapter, builder)
}

// NewBuilder returns a new encryption's public builder
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents a public key adapter
type Adapter interface {
	FromBytes(input []byte) (Key, error)
	FromEncoded(encoded string) (Key, error)
	ToBytes(key Key) []byte
	ToEncoded(key Key) string
	ToHash(key Key) (*hash.Hash, error)
}

// Builder represents a publicKey builder
type Builder interface {
	Create() Builder
	WithKey(key rsa.PublicKey) Builder
	Now() (Key, error)
}

// Key represents an encryption public key
type Key interface {
	Key() rsa.PublicKey
	Encrypt(msg []byte) ([]byte, error)
}
