package displays

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/cameras"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials"
	"github.com/deepvalue-network/software/treedee/domain/worlds/viewports"
)

type builder struct {
	id       *uuid.UUID
	index    uint
	viewport viewports.Viewport
	camera   cameras.Camera
	mat      materials.Material
}

func createBuilder() Builder {
	out := builder{
		id:       nil,
		index:    0,
		viewport: nil,
		camera:   nil,
		mat:      nil,
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

// WithIndex adds an index to the builder
func (app *builder) WithIndex(index uint) Builder {
	app.index = index
	return app
}

// WithViewport adds a viewport to the builder
func (app *builder) WithViewport(viewport viewports.Viewport) Builder {
	app.viewport = viewport
	return app
}

// WithCamera adds a camera to the builder
func (app *builder) WithCamera(camera cameras.Camera) Builder {
	app.camera = camera
	return app
}

// WithMaterial adds a material to the builder
func (app *builder) WithMaterial(mat materials.Material) Builder {
	app.mat = mat
	return app
}

// Now builds a new Display instance
func (app *builder) Now() (Display, error) {
	if app.id == nil {
		return nil, errors.New("the ID is mandatory in order to build a Display instance")
	}

	if app.viewport == nil {
		return nil, errors.New("the viewport is mandatory in order to build a Display instance")
	}

	if app.camera == nil {
		return nil, errors.New("the camera is mandatory in order to build a Display instance")
	}

	if app.mat != nil {
		return createDisplayWithMaterial(app.id, app.index, app.viewport, app.camera, app.mat), nil
	}

	return createDisplay(app.id, app.index, app.viewport, app.camera), nil
}
