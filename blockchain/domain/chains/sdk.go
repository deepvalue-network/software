package chains

import (
	"time"

	uuid "github.com/satori/go.uuid"
	block_mined "github.com/steve-care-software/products/blockchain/domain/blocks/mined"
	"github.com/steve-care-software/products/blockchain/domain/chains/peers"
	"github.com/steve-care-software/products/blockchain/domain/genesis"
	link_mined "github.com/steve-care-software/products/blockchain/domain/links/mined"
)

// NewBuilder creates a new builder instance
func NewBuilder(peerSyncInterval time.Duration) Builder {
	peersBuilder := peers.NewBuilder()
	return createBuilder(
		peersBuilder,
		peerSyncInterval,
	)
}

// Validator represents a chain validator
type Validator interface {
	Validate(chain Chain) bool
}

// Builder represents a chain builder
type Builder interface {
	Create() Builder
	WithOriginal(original Chain) Builder
	WithGenesis(gen genesis.Genesis) Builder
	WithRoot(root block_mined.Block) Builder
	WithHead(head link_mined.Link) Builder
	Now() (Chain, error)
}

// Chain represents a chain
type Chain interface {
	ID() *uuid.UUID
	Peers() peers.Peers
	Genesis() genesis.Genesis
	Root() block_mined.Block
	TotalHashes() uint
	Height() uint
	CreatedOn() time.Time
	LastUpdatedOn() time.Time
	HasHead() bool
	Head() link_mined.Link
}

// Repository represents a chain repository
type Repository interface {
	List() ([]*uuid.UUID, error)
	Retrieve(id *uuid.UUID) (Chain, error)
}

// Service represents a chain service
type Service interface {
	Insert(chain Chain) error
	Update(original Chain, updated Chain) error
	Delete(chain Chain) error
}
