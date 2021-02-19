package governments

import (
	"errors"

	"github.com/deepvalue-network/software/libs/hash"
	uuid "github.com/satori/go.uuid"
)

type builder struct {
	hashAdapter hash.Adapter
	id          *uuid.UUID
	current     Content
	prev        Government
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		id:          nil,
		current:     nil,
		prev:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithID adds an ID to the builder
func (app *builder) WithID(id *uuid.UUID) Builder {
	app.id = id
	return app
}

// WithCurrent adds a current content to the builder
func (app *builder) WithCurrent(current Content) Builder {
	app.current = current
	return app
}

// WithPrevious adds a previous governement to the builder
func (app *builder) WithPrevious(prev Government) Builder {
	app.prev = prev
	return app
}

// Now builds a new Government instance
func (app *builder) Now() (Government, error) {
	if app.current == nil {
		return nil, errors.New("the current Content is mandatory in order to build a Government instance")
	}

	data := [][]byte{
		app.current.Hash().Bytes(),
	}

	if app.prev != nil {
		data = append(data, app.prev.Hash().Bytes())
	}

	if app.prev == nil {
		if app.id == nil {
			id := uuid.NewV4()
			app.id = &id
		}

		data = append(data, app.id.Bytes())
	}

	hash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.prev != nil {
		return createGovernmentWithPrevious(*hash, app.current, app.prev), nil
	}

	return createGovernment(*hash, app.id, app.current), nil
}
