package worlds

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes"
)

type builder struct {
	id                *uuid.UUID
	currentSceneIndex uint
	scenes            []scenes.Scene
}

func createBuilder() Builder {
	out := builder{
		id:                nil,
		currentSceneIndex: 0,
		scenes:            nil,
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

// WithCurrentSceneIndex adds a currentSceneIndex to the builder
func (app *builder) WithCurrentSceneIndex(currentSceneIndex uint) Builder {
	app.currentSceneIndex = currentSceneIndex
	return app
}

// WithScenes add scenes to the builder
func (app *builder) WithScenes(scenes []scenes.Scene) Builder {
	app.scenes = scenes
	return app
}

// Now builds a new World instance
func (app *builder) Now() (World, error) {
	if app.id == nil {
		return nil, errors.New("the ID is mandatory in order to build a World instance")
	}

	if app.scenes != nil && len(app.scenes) <= 0 {
		app.scenes = nil
	}

	if app.scenes == nil {
		return nil, errors.New("there must be at least 1 Scene in order to build a World instance")
	}

	return createWorld(app.id, app.currentSceneIndex, app.scenes), nil
}
