package owners

import (
	"time"

	"github.com/steve-care-software/products/libs/cryptography/pk/encryption"
	"github.com/steve-care-software/products/libs/cryptography/pk/signature"
	"github.com/steve-care-software/products/libs/hash"
)

// Builder represents an owner builder
type Builder interface {
	Create() Builder
	WithSeed(seed string) Builder
	WithSignaturePK(sigPK signature.PrivateKey) Builder
	WithEncryptionPK(encPK encryption.PrivateKey) Builder
	CreatedOn(createdOn time.Time) Builder
	Now() (Owner, error)
}

// Owner represents an owner
type Owner interface {
	Hash() hash.Hash
	Seed() string
	Signature() signature.PrivateKey
	Encryption() encryption.PrivateKey
	CreatedOn() time.Time
}

// Repository represents an owner repository
type Repository interface {
	List() []hash.Hash
	Retrieve(hash hash.Hash, seed string, password string) (Owner, error)
}

// Service represents an owner service
type Service interface {
	Insert(owner Owner, password string) error
	Delete(owner Owner, password string) error
}
