package nodes

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/math/fl32"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/huds/displays"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models"
)

type builder struct {
	id        *uuid.UUID
	pos       *fl32.Vec2
	posVar    string
	angle     float32
	direction *fl32.Vec3
	oriVar    string
	model     models.Model
	display   displays.Display
	children  []Node
}

func createBuilder() Builder {
	out := builder{
		id:        nil,
		pos:       nil,
		posVar:    "",
		angle:     0.0,
		direction: nil,
		oriVar:    "",
		model:     nil,
		display:   nil,
		children:  nil,
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

// WithPosition adds a position to the builder
func (app *builder) WithPosition(pos fl32.Vec2) Builder {
	app.pos = &pos
	return app
}

// WithPositionVariable adds a position variable to the builder
func (app *builder) WithPositionVariable(posVar string) Builder {
	app.posVar = posVar
	return app
}

// WithOrientationAngle adds an orientation angle to the buWithChildren adds children nodes to the builderilder
func (app *builder) WithOrientationAngle(angle float32) Builder {
	app.angle = angle
	return app
}

// WithOrientationDirection adds an orientation direction to the builder
func (app *builder) WithOrientationDirection(direction fl32.Vec3) Builder {
	app.direction = &direction
	return app
}

// WithOrientationVariable adds an orientation variable to the builder
func (app *builder) WithOrientationVariable(oriVar string) Builder {
	app.oriVar = oriVar
	return app
}

// WithOrWithModelientationDirection adds a model direction to the builder
func (app *builder) WithModel(model models.Model) Builder {
	app.model = model
	return app
}

// WithDisplay adds a display to the builder
func (app *builder) WithDisplay(display displays.Display) Builder {
	app.display = display
	return app
}

// WithChildren adds children nodes to the builder
func (app *builder) WithChildren(children []Node) Builder {
	app.children = children
	return app
}

// Now builds a new Node instance
func (app *builder) Now() (Node, error) {
	if app.id == nil {
		return nil, errors.New("the ID is mandatory in order to build a Node instance")
	}

	if app.pos == nil {
		return nil, errors.New("the position Vec2 is mandatory in order to build a Node instance")
	}

	if app.posVar == "" {
		return nil, errors.New("the position variable is mandatory in order to build a Node instance")
	}

	if app.angle >= 360 {
		return nil, errors.New("the angle must be smaller than 360")
	}

	if app.direction == nil {
		return nil, errors.New("the direction Vec3 is mandatory in order to build a Node instance")
	}

	if app.oriVar == "" {
		return nil, errors.New("the orientation variable is mandatory in order to build a Node instance")
	}

	if app.children != nil && len(app.children) <= 0 {
		app.children = nil
	}

	var content Content
	if app.model != nil {
		content = createContentWithModel(app.model)
	}

	if app.display != nil {
		content = createContentWithDisplay(app.display)
	}

	pos := createPosition(*app.pos, app.posVar)
	orientation := createOrientation(app.angle, *app.direction, app.oriVar)
	if content != nil {
		if app.children != nil {
			return createNodeWithContentAndChildren(app.id, pos, orientation, content, app.children), nil
		}

		return createNodeWithContent(app.id, pos, orientation, content), nil
	}

	if app.children != nil {
		return createNodeWithChildren(app.id, pos, orientation, app.children), nil
	}

	return createNode(app.id, pos, orientation), nil
}
