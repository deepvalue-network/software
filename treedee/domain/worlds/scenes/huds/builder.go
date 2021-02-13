package huds

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/huds/nodes"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials"
)

type builder struct {
	id    *uuid.UUID
	nodes []nodes.Node
	mat   materials.Material
}

func createBuilder() Builder {
	out := builder{
		id:    nil,
		nodes: nil,
		mat:   nil,
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

// WithNodes add nodes to the builder
func (app *builder) WithNodes(nodes []nodes.Node) Builder {
	app.nodes = nodes
	return app
}

// WithMaterial adds a material to the builder
func (app *builder) WithMaterial(mat materials.Material) Builder {
	app.mat = mat
	return app
}

// Now builds a new Hud instance
func (app *builder) Now() (Hud, error) {
	if app.id == nil {
		return nil, errors.New("the ID is mandatory in order to build a Hud instance")
	}

	if app.nodes != nil && len(app.nodes) <= 0 {
		app.nodes = nil
	}

	if app.nodes != nil && app.mat != nil {
		return createHudWithNodesAndMaterial(app.id, app.nodes, app.mat), nil
	}

	if app.nodes != nil {
		return createHudWithNodes(app.id, app.nodes), nil
	}

	if app.mat != nil {
		return createHudWithMaterial(app.id, app.mat), nil
	}

	return nil, errors.New("the HUD is invalid")
}
