package links

import (
	"errors"
	"strconv"

	"github.com/deepvalue-network/software/blockchain/domain/blocks"
	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	hashAdapter   hash.Adapter
	index         uint
	prevMinedLink *hash.Hash
	nextBlock     blocks.Block
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter:   hashAdapter,
		index:         0,
		prevMinedLink: nil,
		nextBlock:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithIndex adds an index to the builder
func (app *builder) WithIndex(index uint) Builder {
	app.index = index
	return app
}

// WithPreviousMinedLink adds a previous mined link linkHash to the builder
func (app *builder) WithPreviousMinedLink(prevMinedLink hash.Hash) Builder {
	app.prevMinedLink = &prevMinedLink
	return app
}

// WithNextBlock adds a next block to the builder
func (app *builder) WithNextBlock(block blocks.Block) Builder {
	app.nextBlock = block
	return app
}

// Now builds a new Link instance
func (app *builder) Now() (Link, error) {
	if app.prevMinedLink == nil {
		return nil, errors.New("the previous mined link hash is mandatory in order to build a Link instance")
	}

	if app.nextBlock == nil {
		return nil, errors.New("the next block is mandatory in order to build a Link instance")
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(strconv.Itoa(int(app.index))),
		app.prevMinedLink.Bytes(),
		app.nextBlock.Tree().Head().Bytes(),
	})

	if err != nil {
		return nil, err
	}

	return createLink(*hash, app.index, *app.prevMinedLink, app.nextBlock), nil
}
