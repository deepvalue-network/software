package hashtree

import (
	"errors"
)

type builder struct {
	blocks [][]byte
	js     []byte
}

func createBuilder() Builder {
	out := builder{
		blocks: nil,
		js:     nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithBlocks add blocks to the builder
func (app *builder) WithBlocks(blocks [][]byte) Builder {
	app.blocks = blocks
	return app
}

// Now builds a new HashTree instance
func (app *builder) Now() (HashTree, error) {
	if app.blocks != nil {
		return createHashTreeFromBlocks(app.blocks)
	}

	return nil, errors.New("the HashTree is invalid")
}
