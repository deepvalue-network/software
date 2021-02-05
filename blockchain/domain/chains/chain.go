package chains

import (
	"time"

	uuid "github.com/satori/go.uuid"
	block_mined "github.com/deepvalue-network/software/blockchain/domain/blocks/mined"
	"github.com/deepvalue-network/software/blockchain/domain/chains/peers"
	"github.com/deepvalue-network/software/blockchain/domain/genesis"
	link_mined "github.com/deepvalue-network/software/blockchain/domain/links/mined"
)

type chain struct {
	id          *uuid.UUID        `hydro:"ID, ID"`
	peers       peers.Peers       `hydro:"Peers, Peers"`
	gen         genesis.Genesis   `hydro:"Genesis, Genesis"`
	root        block_mined.Block `hydro:"Hash, Hash"`
	totalHashes uint              `hydro:"TotalHashes, TotalHashes"`
	height      uint              `hydro:"Height, Height"`
	createdOn   time.Time         `hydro:"CreatedOn, CreatedOn"`
	head        link_mined.Link   `hydro:"Head, Head"`
}

func createChain(
	id *uuid.UUID,
	peers peers.Peers,
	gen genesis.Genesis,
	root block_mined.Block,
	totalHashes uint,
	height uint,
	createdOn time.Time,
) Chain {
	return createChainInternally(
		id,
		peers,
		gen,
		root,
		totalHashes,
		height,
		createdOn,
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
	head link_mined.Link,
) Chain {
	out := chain{
		id:          id,
		peers:       peers,
		gen:         gen,
		root:        root,
		totalHashes: totalHashes,
		height:      height,
		createdOn:   createdOn,
		head:        head,
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

// HasHead returns true if there is a head, false otherwise
func (obj *chain) HasHead() bool {
	return obj.head != nil
}

// Head returns the head, if any
func (obj *chain) Head() link_mined.Link {
	return obj.head
}
