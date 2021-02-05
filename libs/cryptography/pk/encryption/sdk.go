package encryption

import (
	"crypto/rsa"

	"github.com/deepvalue-network/software/libs/cryptography/pk/encryption/public"
)

// NewFactory returns a new encryption's privatekey factory
func NewFactory(bitrate int) Factory {
	builder := NewBuilder()
	return createFactory(builder, bitrate)
}

// NewAdapter returns a new encryption's privatekey adapter
func NewAdapter() Adapter {
	builder := NewBuilder()
	return createAdapter(builder)
}

// NewBuilder returns a new encryption's privatekey builder
func NewBuilder() Builder {
	pubKeyBuilder := public.NewBuilder()
	return createBuilder(pubKeyBuilder)
}

// Factory represents a privateKey factory
type Factory interface {
	Create() (PrivateKey, error)
}

// Adapter represents a privateKey adapter
type Adapter interface {
	FromBytes(bytes []byte) (PrivateKey, error)
	FromEncoded(encoded string) (PrivateKey, error)
	ToBytes(pk PrivateKey) []byte
	ToEncoded(pk PrivateKey) string
}

// Builder represents a privateKey builder
type Builder interface {
	Create() Builder
	WithPK(pk rsa.PrivateKey) Builder
	Now() (PrivateKey, error)
}

// PrivateKey represents an encryption private key
type PrivateKey interface {
	Key() rsa.PrivateKey
	Public() public.Key
	Decrypt(cipher []byte) ([]byte, error)
}
