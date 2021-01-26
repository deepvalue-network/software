package repositories

import (
	uuid "github.com/satori/go.uuid"
	"github.com/steve-care-software/products/blockchain/domain/blocks"
	mined_block "github.com/steve-care-software/products/blockchain/domain/blocks/mined"
	"github.com/steve-care-software/products/blockchain/domain/chains"
	"github.com/steve-care-software/products/blockchain/domain/chains/peers"
	"github.com/steve-care-software/products/blockchain/domain/links"
	mined_link "github.com/steve-care-software/products/blockchain/domain/links/mined"
	"github.com/steve-care-software/products/libs/hash"
)

// NewApplication creates a new application instance
func NewApplication(
	block Block,
	minedBlock MinedBlock,
	link Link,
	minedLink MinedLink,
	chain Chain,
) Application {
	return createApplication(
		block,
		minedBlock,
		link,
		minedLink,
		chain,
	)
}

// RemoteBuilder represents a remote application builder
type RemoteBuilder interface {
	Create() RemoteBuilder
	WithPeer(peer peers.Peer) RemoteBuilder
	Now() (Application, error)
}

// Application represents a remote application
type Application interface {
	Block() Block
	MinedBlock() MinedBlock
	Link() Link
	MinedLink() MinedLink
	Chain() Chain
}

// Block represents a block application
type Block interface {
	List() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (blocks.Block, error)
}

// MinedBlock represents the mined block application
type MinedBlock interface {
	List() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (mined_block.Block, error)
}

// Link represents the link application
type Link interface {
	List() ([]hash.Hash, error)
	Retrieve(hash hash.Hash) (links.Link, error)
}

// MinedLink represents the mined link application
type MinedLink interface {
	List() ([]hash.Hash, error)
	Head() (mined_link.Link, error)
	Retrieve(hash hash.Hash) (mined_link.Link, error)
}

// Chain represents a chain application
type Chain interface {
	List() ([]*uuid.UUID, error)
	Retrieve(id *uuid.UUID) (chains.Chain, error)
}
