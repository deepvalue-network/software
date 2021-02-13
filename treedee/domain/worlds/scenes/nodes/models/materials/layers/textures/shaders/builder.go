package shaders

import (
	"errors"

	uuid "github.com/satori/go.uuid"
)

type builder struct {
	id        *uuid.UUID
	code      string
	isDynamic bool
}

func createBuilder() Builder {
	out := builder{
		id:        nil,
		code:      "",
		isDynamic: false,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithID adds an ID to the builder
func (app *builder) WithID(id *uuid.UUID) Builder {
	app.id = id
	return app
}

// WithCode adds code to the builder
func (app *builder) WithCode(code string) Builder {
	app.code = code
	return app
}

// IsDynamic flags the builder as dynamic
func (app *builder) IsDynamic() Builder {
	app.isDynamic = true
	return app
}

// Now builds a new Shader instance
func (app *builder) Now() (Shader, error) {
	if app.id == nil {
		return nil, errors.New("the id is mandatory in order to build a Shader instance")
	}

	if app.code == "" {
		return nil, errors.New("the code is mandatory in order to build a Shader instance")
	}

	return createShader(app.id, app.code, app.isDynamic), nil
}
