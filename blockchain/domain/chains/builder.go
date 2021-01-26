package chains

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
	block_mined "github.com/steve-care-software/products/blockchain/domain/blocks/mined"
	"github.com/steve-care-software/products/blockchain/domain/chains/peers"
	"github.com/steve-care-software/products/blockchain/domain/genesis"
	link_mined "github.com/steve-care-software/products/blockchain/domain/links/mined"
)

type builder struct {
	peersBuilder     peers.Builder
	peerSyncInterval time.Duration
	original         Chain
	gen              genesis.Genesis
	root             block_mined.Block
	head             link_mined.Link
}

func createBuilder(
	peersBuilder peers.Builder,
	peerSyncInterval time.Duration,
) Builder {
	out := builder{
		peersBuilder:     peersBuilder,
		peerSyncInterval: peerSyncInterval,
		original:         nil,
		gen:              nil,
		root:             nil,
		head:             nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.peersBuilder, app.peerSyncInterval)
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

// Now builds a new Chain instance
func (app *builder) Now() (Chain, error) {
	if app.original != nil {
		if app.head == nil {
			return nil, errors.New("the head mined link is mandatory in order to build an updated Chain instance")
		}

		id := app.original.ID()
		peers := app.original.Peers()
		root := app.original.Root()
		gen := app.original.Genesis()
		totalHashes := app.original.TotalHashes() + uint(len(app.head.Link().NextBlock().Hashes()))
		height := app.original.Height() + 1
		createdOn := app.original.CreatedOn()
		lastUpdatedOn := time.Now().UTC()
		return createChainWithHead(
			id,
			peers,
			gen,
			root,
			totalHashes,
			height,
			createdOn,
			lastUpdatedOn,
			app.head,
		), nil
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

	id := uuid.NewV4()
	totalHashes := uint(len(app.root.Block().Hashes()))
	height := uint(0)
	createdOn := time.Now().UTC()
	lastUpdatedOn := time.Now().UTC()
	return createChain(
		&id,
		peers,
		app.gen,
		app.root,
		totalHashes,
		height,
		createdOn,
		lastUpdatedOn,
	), nil
}
