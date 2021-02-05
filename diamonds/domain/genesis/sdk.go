package genesis

import (
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/chains"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder(minPubKeysInOwner uint) Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter, minPubKeysInOwner)
}

// Builder represents a genesis builder
type Builder interface {
	Create() Builder
	WithAmount(amount uint64) Builder
	WithChain(chain chains.Chain) Builder
	WithHashedPubKeysOwner(hashedPubKeysOwner []hash.Hash) Builder
	CreatedOn(createdOn time.Time) Builder
	ActiveOn(activeOn time.Time) Builder
	Now() (Genesis, error)
}

// Genesis represents a genesis diamond
type Genesis interface {
	Hash() hash.Hash
	Amount() uint64
	Chain() chains.Chain
	Owner() []hash.Hash
	CreatedOn() time.Time
	ActiveOn() time.Time
}

// Repository represents a genesis repository
type Repository interface {
	List() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (Genesis, error)
}

// Service represents a genesis service
type Service interface {
	Insert(gen Genesis) error
	Delete(gen Genesis) error
}
