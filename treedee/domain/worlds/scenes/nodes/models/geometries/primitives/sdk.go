package primitives

import (
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/geometries/primitives/cubes"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/geometries/primitives/planes"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/geometries/primitives/prisms"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/geometries/primitives/spheres"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/geometries/vertices"
)

// Adapter represents a primitive adapter
type Adapter interface {
	ToVertices(prim Primitive) ([]vertices.Vertex, error)
}

// Builder represents a primitive builder
type Builder interface {
	Create() Builder
	WithID(id *uuid.UUID) Builder
	WithPlane(plane planes.Plane) Builder
	WithPrism(prism prisms.Prism) Builder
	WithCube(cube cubes.Cube) Builder
	WithSphere(sphere spheres.Sphere) Builder
	Now() (Primitive, error)
}

// Primitive represents a primitive geometry
type Primitive interface {
	ID() *uuid.UUID
	IsPlane() bool
	Plane() planes.Plane
	IsPrism() bool
	Prism() prisms.Prism
	IsCube() bool
	Cube() cubes.Cube
	IsSphere() bool
	Sphere() spheres.Sphere
}
