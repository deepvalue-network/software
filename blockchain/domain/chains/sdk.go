package chains

import (
	"time"

	block_mined "github.com/deepvalue-network/software/blockchain/domain/blocks/mined"
	mined_block "github.com/deepvalue-network/software/blockchain/domain/blocks/mined"
	"github.com/deepvalue-network/software/blockchain/domain/chains/peers"
	"github.com/deepvalue-network/software/blockchain/domain/genesis"
	link_mined "github.com/deepvalue-network/software/blockchain/domain/links/mined"
	mined_link "github.com/deepvalue-network/software/blockchain/domain/links/mined"
	uuid "github.com/satori/go.uuid"
)

// NewValidator creates a new validator instance
func NewValidator(
	minedBlockValidator mined_block.Validator,
	minedLinkValidator mined_link.Validator,
	chainRepository Repository,
) Validator {
	return createValidator(minedBlockValidator, minedLinkValidator, chainRepository)
}

// NewBuilder creates a new builder instance
func NewBuilder(peerSyncInterval time.Duration) Builder {
	peersBuilder := peers.NewBuilder()
	return createBuilder(
		peersBuilder,
		peerSyncInterval,
	)
}

// NewPointer creates a new chain pointer
func NewPointer() *chain {
	return new(chain)
}

// Validator represents a chain validator
type Validator interface {
	Execute(chain Chain) error
}

// Builder represents a chain builder
type Builder interface {
	Create() Builder
	WithID(id *uuid.UUID) Builder
	WithPeers(peers peers.Peers) Builder
	WithOriginal(original Chain) Builder
	WithGenesis(gen genesis.Genesis) Builder
	WithRoot(root block_mined.Block) Builder
	WithHead(head link_mined.Link) Builder
	CreatedOn(createdOn time.Time) Builder
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
