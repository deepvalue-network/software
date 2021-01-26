package hashtree

import (
	"encoding/json"
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

// WithJS add js to the builder
func (app *builder) WithJS(js []byte) Builder {
	app.js = js
	return app
}

// Now builds a new HashTree instance
func (app *builder) Now() (HashTree, error) {
	if app.blocks != nil {
		return createHashTreeFromBlocks(app.blocks)
	}

	if app.js != nil {
		ptr := new(hashtree)
		err := json.Unmarshal(app.js, ptr)
		if err != nil {
			return nil, err
		}

		return ptr, nil
	}

	return nil, errors.New("the HashTree is invalid")
}
