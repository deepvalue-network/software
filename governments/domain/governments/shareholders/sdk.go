package shareholders

import (
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/chains"
	"github.com/deepvalue-network/software/libs/hash"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewShareHolderBuilder creates a new shareHolder builder instance
func NewShareHolderBuilder(minHashesInShareHolders uint) ShareHolderBuilder {
	hashAdapter := hash.NewAdapter()
	return createShareHolderBuilder(hashAdapter, minHashesInShareHolders)
}

// Builder represents a shareholders builder
type Builder interface {
	Create() Builder
	WithShareHolders(shareHolders []ShareHolder) Builder
	Now() (ShareHolders, error)
}

// ShareHolders represents shareholders
type ShareHolders interface {
	Hash() hash.Hash
	All() []ShareHolder
	Same(pubKeyHashes []hash.Hash) bool
}

// ShareHolderBuilder represents a shareholder builder
type ShareHolderBuilder interface {
	Create() ShareHolderBuilder
	WithChain(chain chains.Chain) ShareHolderBuilder
	WithKeys(keys []hash.Hash) ShareHolderBuilder
	WithPower(power uint) ShareHolderBuilder
	CreatedOn(createdOn time.Time) ShareHolderBuilder
	Now() (ShareHolder, error)
}

// ShareHolder represents a shareholder
type ShareHolder interface {
	Hash() hash.Hash
	Chain() chains.Chain
	Keys() []hash.Hash
	Power() uint
	CreatedOn() time.Time
	Same(pubKeyHashes []hash.Hash) bool
	Contains(hashedPubKey hash.Hash) bool
}
