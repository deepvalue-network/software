package opengl

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type camera struct {
	id         *uuid.UUID
	index      uint
	projection CameraProjection
	lookAt     CameraLookAt
}

func createCamera(
	id *uuid.UUID,
	index uint,
	projection CameraProjection,
	lookAt CameraLookAt,
) Camera {
	out := camera{
		id:         id,
		index:      index,
		projection: projection,
		lookAt:     lookAt,
	}

	return &out
}

// ID returns the id
func (obj *camera) ID() *uuid.UUID {
	return obj.id
}

// Index returns the index
func (obj *camera) Index() uint {
	return obj.index
}

// Projection returns the projection
func (obj *camera) Projection() CameraProjection {
	return obj.projection
}

// LookAt returns the lookAt
func (obj *camera) LookAt() CameraLookAt {
	return obj.lookAt
}

// Render the camera
func (obj *camera) Render(
	delta time.Duration,
	pos Position,
	orientation Orientation,
	activeScene Scene,
	program uint32,
) error {
	return nil
}
