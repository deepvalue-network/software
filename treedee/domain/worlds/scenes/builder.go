package scenes

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/huds"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes"
)

type builder struct {
	id    *uuid.UUID
	index uint
	hud   huds.Hud
	nodes []nodes.Node
}

func createBuilder() Builder {
	out := builder{
		id:    nil,
		index: 0,
		hud:   nil,
		nodes: nil,
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

// WithHud adds an heads-up display to the builder
func (app *builder) WithHud(hud huds.Hud) Builder {
	app.hud = hud
	return app
}

// WithNodes add nodes to the builder
func (app *builder) WithNodes(nodes []nodes.Node) Builder {
	app.nodes = nodes
	return app
}

// Now builds a new Scene instance
func (app *builder) Now() (Scene, error) {
	if app.id == nil {
		return nil, errors.New("the ID is mandatory in order to build a Scene instance")
	}

	if app.hud == nil {
		return nil, errors.New("the HUD is mandatory in order to build a Scene instance")
	}

	if app.nodes != nil && len(app.nodes) <= 0 {
		app.nodes = nil
	}

	if app.nodes == nil {
		return nil, errors.New("there must be at least 1 Node in order to build a Scene instance")
	}

	if !app.hasCamera(app.nodes) {
		return nil, errors.New("the nodes must contain at least 1 camera in order to build a valid Scene instance")
	}

	return createScene(app.id, app.index, app.hud, app.nodes), nil
}

func (app *builder) hasCamera(nodes []nodes.Node) bool {
	for _, oneNode := range app.nodes {
		if !oneNode.HasContent() {
			continue
		}

		content := oneNode.Content()
		if !content.IsCamera() {
			continue
		}

		if oneNode.HasChildren() {
			children := oneNode.Children()
			childHasCamera := app.hasCamera(children)
			if childHasCamera {
				return true
			}

			continue
		}

		return true
	}

	return false
}
