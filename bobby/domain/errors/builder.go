package errors

import (
	"errors"
	"strconv"

	"github.com/steve-care-software/products/bobby/domain/resources"
	"github.com/steve-care-software/products/libs/hash"
)

type builder struct {
	hashAdapter      hash.Adapter
	immutableBuilder resources.ImmutableBuilder
	message          string
	code             uint
	parent           Error
}

func createBuilder(
	hashAdapter hash.Adapter,
	immutableBuilder resources.ImmutableBuilder,
) Builder {
	out := builder{
		hashAdapter:      hashAdapter,
		immutableBuilder: immutableBuilder,
		message:          "",
		code:             0,
		parent:           nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(app.hashAdapter, app.immutableBuilder)
}

// WithMessage adds a message to the builder
func (app *builder) WithMessage(message string) Builder {
	app.message = message
	return app
}

// WithCode adds a code to the builder
func (app *builder) WithCode(code uint) Builder {
	app.code = code
	return app
}

// WithParent adds a parent to the builder
func (app *builder) WithParent(parent Error) Builder {
	app.parent = parent
	return app
}

// Now builds a new Error instance
func (app *builder) Now() (Error, error) {
	if app.message == "" {
		return nil, errors.New("the message is mandatory in order to build an Error instance")
	}

	if app.code <= 0 {
		return nil, errors.New("the code must be greater than zero (0) in order to build an Error instance")
	}

	hash, err := app.hashAdapter.FromMultiBytes([][]byte{
		[]byte(app.message),
		[]byte(strconv.Itoa(int(app.code))),
	})

	if err != nil {
		return nil, err
	}

	immutable, err := app.immutableBuilder.Create().WithHash(*hash).Now()
	if err != nil {
		return nil, err
	}

	if app.parent != nil {
		return createErrorWithParent(immutable, app.message, app.code, app.parent), nil
	}

	return createError(immutable, app.message, app.code), nil
}
