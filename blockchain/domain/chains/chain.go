package chains

import (
	"time"

	uuid "github.com/satori/go.uuid"
	block_mined "github.com/steve-care-software/products/blockchain/domain/blocks/mined"
	"github.com/steve-care-software/products/blockchain/domain/chains/peers"
	"github.com/steve-care-software/products/blockchain/domain/genesis"
	link_mined "github.com/steve-care-software/products/blockchain/domain/links/mined"
)

type chain struct {
	id            *uuid.UUID
	peers         peers.Peers
	gen           genesis.Genesis
	root          block_mined.Block
	totalHashes   uint
	height        uint
	createdOn     time.Time
	lastUpdatedOn time.Time
	head          link_mined.Link
}

func createChain(
	id *uuid.UUID,
	peers peers.Peers,
	gen genesis.Genesis,
	root block_mined.Block,
	totalHashes uint,
	height uint,
	createdOn time.Time,
	lastUpdatedOn time.Time,
) Chain {
	return createChainInternally(
		id,
		peers,
		gen,
		root,
		totalHashes,
		height,
		createdOn,
		lastUpdatedOn,
		nil,
	)
}

func createChainWithHead(
	id *uuid.UUID,
	peers peers.Peers,
	gen genesis.Genesis,
	root block_mined.Block,
	totalHashes uint,
	height uint,
	createdOn time.Time,
	lastUpdatedOn time.Time,
	head link_mined.Link,
) Chain {
	return createChainInternally(
		id,
		peers,
		gen,
		root,
		totalHashes,
		height,
		createdOn,
		lastUpdatedOn,
		head,
	)
}

func createChainInternally(
	id *uuid.UUID,
	peers peers.Peers,
	gen genesis.Genesis,
	root block_mined.Block,
	totalHashes uint,
	height uint,
	createdOn time.Time,
	lastUpdatedOn time.Time,
	head link_mined.Link,
) Chain {
	out := chain{
		id:            id,
		peers:         peers,
		gen:           gen,
		root:          root,
		totalHashes:   totalHashes,
		height:        height,
		createdOn:     createdOn,
		lastUpdatedOn: lastUpdatedOn,
		head:          head,
	}

	return &out
}

// ID returns the id
func (obj *chain) ID() *uuid.UUID {
	return obj.id
}

// Peers returns the peers
func (obj *chain) Peers() peers.Peers {
	return obj.peers
}

// Genesis returns the genesis
func (obj *chain) Genesis() genesis.Genesis {
	return obj.gen
}

// Root returns the root block
func (obj *chain) Root() block_mined.Block {
	return obj.root
}

// TotalHashes returns the total hashes
func (obj *chain) TotalHashes() uint {
	return obj.totalHashes
}

// Height returns the height
func (obj *chain) Height() uint {
	return obj.height
}

// CreatedOn returns the creation time
func (obj *chain) CreatedOn() time.Time {
	return obj.createdOn
}

// LastUpdatedOn returns the lastUpdatedOn time
func (obj *chain) LastUpdatedOn() time.Time {
	return obj.lastUpdatedOn
}

// HasHead returns true if there is a head, false otherwise
func (obj *chain) HasHead() bool {
	return obj.head != nil
}

// Head returns the head, if any
func (obj *chain) Head() link_mined.Link {
	return obj.head
}
