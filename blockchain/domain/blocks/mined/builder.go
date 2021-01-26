package mined

import (
	"errors"
	"strconv"
	"time"

	"github.com/steve-care-software/products/blockchain/domain/blocks"
	"github.com/steve-care-software/products/libs/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	block       blocks.Block
	results     string
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		block:       nil,
		results:     "",
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithBlock adds a block to the builder
func (app *builder) WithBlock(block blocks.Block) Builder {
	app.block = block
	return app
}

// WithResults add results to the builder
func (app *builder) WithResults(results string) Builder {
	app.results = results
	return app
}

// Now builds a new Block instance
func (app *builder) Now() (Block, error) {
	if app.block == nil {
		return nil, errors.New("the block is mandatory in order to build a mined Block instance")
	}

	if app.results == "" {
		return nil, errors.New("the results are mandatory in order to build a mined Block instance")
	}

	createdOn := time.Now().UTC()
	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		app.block.Tree().Head().Bytes(),
		[]byte(app.results),
		[]byte(strconv.Itoa(createdOn.Nanosecond())),
	})

	if err != nil {
		return nil, err
	}

	return createBlock(*hash, app.block, app.results, createdOn), nil
}
