package nodes

import (
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/math/fl32"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/huds/displays"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents the node builder
type Builder interface {
	Create() Builder
	WithID(id *uuid.UUID) Builder
	WithPosition(pos fl32.Vec2) Builder
	WithPositionVariable(posVar string) Builder
	WithOrientationAngle(angle float32) Builder
	WithOrientationDirection(direction fl32.Vec3) Builder
	WithOrientationVariable(oriVar string) Builder
	WithModel(model models.Model) Builder
	WithDisplay(display displays.Display) Builder
	WithChildren(children []Node) Builder
	Now() (Node, error)
}

// Node represents a node
type Node interface {
	ID() *uuid.UUID
	Position() Position
	Orientation() Orientation
	HasContent() bool
	Content() Content
	HasChildren() bool
	Children() []Node
}

// Content represents the node content
type Content interface {
	IsModel() bool
	Model() models.Model
	IsDisplay() bool
	Display() displays.Display
}

// Position represents a position
type Position interface {
	Vector() fl32.Vec2
	Variable() string
}

// Orientation represents an orientation
type Orientation interface {
	Angle() float32
	Direction() fl32.Vec3
	Variable() string
}
