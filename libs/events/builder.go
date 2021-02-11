package events

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/deepvalue-network/software/libs/hash"
)

type builder struct {
	hashAdapter hash.Adapter
	identifier  int
	onEnter     EventFn
	onExit      EventFn
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		identifier:  -1,
		onEnter:     nil,
		onExit:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter)
}

// WithIdentifier adds an identifier to the builder
func (app *builder) WithIdentifier(identifier int) Builder {
	app.identifier = identifier
	return app
}

// OnEnter adds an onEnter func to the builder
func (app *builder) OnEnter(onEnter EventFn) Builder {
	app.onEnter = onEnter
	return app
}

// OnExit adds an onExit func to the builder
func (app *builder) OnExit(onExit EventFn) Builder {
	app.onExit = onExit
	return app
}

// Now builds a new Event instance
func (app *builder) Now() (Event, error) {
	if app.identifier < 0 {
		return nil, errors.New("the identifier is mandatory in order to build an Event instance")
	}

	data := [][]byte{
		[]byte(strconv.Itoa(app.identifier)),
	}

	if app.onEnter != nil {
		fn := app.onEnter
		str := fmt.Sprintf("%v", &fn)
		data = append(data, []byte(str))
	}

	if app.onExit != nil {
		fn := app.onExit
		str := fmt.Sprintf("%v", &fn)
		data = append(data, []byte(str))
	}

	hsh, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.onEnter != nil && app.onExit != nil {
		return createEventWithOnEnterAndOnExit(*hsh, app.identifier, app.onEnter, app.onExit), nil
	}

	if app.onEnter != nil {
		return createEventWithOnEnter(*hsh, app.identifier, app.onEnter), nil
	}

	if app.onExit != nil {
		return createEventWithOnExit(*hsh, app.identifier, app.onExit), nil
	}

	return nil, errors.New("the Event is invalid")
}
