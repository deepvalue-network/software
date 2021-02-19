package shareholders

import (
	"time"

	"github.com/deepvalue-network/software/blockchain/domain/chains"
	"github.com/deepvalue-network/software/libs/cryptography/pk/signature"
	"github.com/deepvalue-network/software/libs/hash"
)

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
	Validate(sig signature.RingSignature) bool
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
}
