package chains

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
	block_mined "github.com/deepvalue-network/software/blockchain/domain/blocks/mined"
	"github.com/deepvalue-network/software/blockchain/domain/chains/peers"
	"github.com/deepvalue-network/software/blockchain/domain/genesis"
	link_mined "github.com/deepvalue-network/software/blockchain/domain/links/mined"
)

type builder struct {
	peersBuilder     peers.Builder
	peerSyncInterval time.Duration
	id               *uuid.UUID
	original         Chain
	peers            peers.Peers
	gen              genesis.Genesis
	root             block_mined.Block
	head             link_mined.Link
	createdOn        *time.Time
}

func createBuilder(
	peersBuilder peers.Builder,
	peerSyncInterval time.Duration,
) Builder {
	out := builder{
		peersBuilder:     peersBuilder,
		peerSyncInterval: peerSyncInterval,
		id:               nil,
		peers:            nil,
		original:         nil,
		gen:              nil,
		root:             nil,
		head:             nil,
		createdOn:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.peersBuilder, app.peerSyncInterval)
}

// WithID adds an ID to the builder
func (app *builder) WithID(id *uuid.UUID) Builder {
	app.id = id
	return app
}

// WithPeers add peers to the builder
func (app *builder) WithPeers(peers peers.Peers) Builder {
	app.peers = peers
	return app
}

// WithOriginal adds an original chain to the builder
func (app *builder) WithOriginal(original Chain) Builder {
	app.original = original
	return app
}

// WithGenesis adds a genesis instance to the builder
func (app *builder) WithGenesis(gen genesis.Genesis) Builder {
	app.gen = gen
	return app
}

// WithRoot adds a root mined block instance to the builder
func (app *builder) WithRoot(root block_mined.Block) Builder {
	app.root = root
	return app
}

// WithHead adds a head mined link instance to the builder
func (app *builder) WithHead(head link_mined.Link) Builder {
	app.head = head
	return app
}

// CreatedOn adds a creation time to the builder
func (app *builder) CreatedOn(createdOn time.Time) Builder {
	app.createdOn = &createdOn
	return app
}

// Now builds a new Chain instance
func (app *builder) Now() (Chain, error) {
	if app.createdOn == nil {
		createdOn := time.Now().UTC()
		app.createdOn = &createdOn
	}

	if app.original != nil {
		if app.head == nil {
			return nil, errors.New("the head mined link is mandatory in order to build an updated Chain instance")
		}

		peers := app.original.Peers()
		if app.peers != nil {
			peers.Merge(app.peers)
		}

		id := app.original.ID()
		root := app.original.Root()
		gen := app.original.Genesis()
		totalHashes := app.original.TotalHashes() + uint(len(app.head.Link().NextBlock().Hashes()))
		height := app.original.Height() + 1
		return createChainWithHead(
			id,
			peers,
			gen,
			root,
			totalHashes,
			height,
			*app.createdOn,
			app.head,
		), nil
	}

	if app.id == nil {
		id := uuid.NewV4()
		app.id = &id
	}

	if app.gen == nil {
		return nil, errors.New("the genesis is mandatory in order to build a new Chain instance")
	}

	if app.root == nil {
		return nil, errors.New("the root mined block is mandatory in order to build a new Chain instance")
	}

	peers, err := app.peersBuilder.Create().WithSyncDuration(app.peerSyncInterval).Now()
	if err != nil {
		return nil, err
	}

	if app.peers != nil {
		peers.Merge(app.peers)
	}

	totalHashes := uint(len(app.root.Block().Hashes()))
	height := uint(0)
	return createChain(
		app.id,
		peers,
		app.gen,
		app.root,
		totalHashes,
		height,
		*app.createdOn,
	), nil
}
