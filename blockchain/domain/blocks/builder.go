package blocks

import (
	"errors"

	"github.com/deepvalue-network/software/libs/hash"
	"github.com/deepvalue-network/software/libs/hashtree"
)

type builder struct {
	hashTreeBuilder hashtree.Builder
	hashes          []hash.Hash
}

func createBuilder(
	hashTreeBuilder hashtree.Builder,
) Builder {
	out := builder{
		hashTreeBuilder: hashTreeBuilder,
		hashes:          nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashTreeBuilder)
}

// WithHashes add hashes to the builder
func (app *builder) WithHashes(hashes []hash.Hash) Builder {
	app.hashes = hashes
	return app
}

// Now builds a new Block instance
func (app *builder) Now() (Block, error) {
	if app.hashes != nil && len(app.hashes) <= 0 {
		app.hashes = nil
	}

	if app.hashes == nil {
		return nil, errors.New("the hashes are mandatory in order to buiild a Block instance")
	}

	blocks := [][]byte{}
	for _, oneHash := range app.hashes {
		blocks = append(blocks, oneHash.Bytes())
	}

	tree, err := app.hashTreeBuilder.Create().WithBlocks(blocks).Now()
	if err != nil {
		return nil, err
	}

	return createBlock(tree, app.hashes), nil
}
