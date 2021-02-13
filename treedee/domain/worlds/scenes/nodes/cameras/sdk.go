package cameras

import (
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/math/fl32"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a camera builder
type Builder interface {
	Create() Builder
	WithID(id *uuid.UUID) Builder
	WithLookAtVariable(lookAtVariable string) Builder
	WithLookAtEye(eye fl32.Vec3) Builder
	WithLookAtCenter(center fl32.Vec3) Builder
	WithLookAtUp(up fl32.Vec3) Builder
	WithProjectionVariable(projVariable string) Builder
	WithProjectionFieldofView(fov float32) Builder
	WithProjectionAspectRatio(aspectRatio float32) Builder
	WithProjectionNear(near float32) Builder
	WithProjectionFar(far float32) Builder
	WithIndex(index uint) Builder
	Now() (Camera, error)
}

// Camera represents a camera
type Camera interface {
	ID() *uuid.UUID
	Index() uint
	Projection() Projection
	LookAt() LookAt
}

// LookAt represents the direction where the camera looks at
type LookAt interface {
	Variable() string
	Eye() fl32.Vec3
	Center() fl32.Vec3
	Up() fl32.Vec3
}

// Projection represents the camera projection
type Projection interface {
	Variable() string
	FieldOfView() float32
	AspectRatio() float32
	Near() float32
	Far() float32
}
