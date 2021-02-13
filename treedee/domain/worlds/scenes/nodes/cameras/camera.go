package cameras

import uuid "github.com/satori/go.uuid"

type camera struct {
	id         *uuid.UUID
	index      uint
	projection Projection
	lookAt     LookAt
}

func createCamera(
	id *uuid.UUID,
	index uint,
	projection Projection,
	lookAt LookAt,
) Camera {
	out := camera{
		id:         id,
		index:      index,
		projection: projection,
		lookAt:     lookAt,
	}

	return &out
}

// ID returns the camera id
func (obj *camera) ID() *uuid.UUID {
	return obj.id
}

// Index returns the camera index
func (obj *camera) Index() uint {
	return obj.index
}

// Projection returns the camera projection
func (obj *camera) Projection() Projection {
	return obj.projection
}

// LookAt returns the camera lookAt
func (obj *camera) LookAt() LookAt {
	return obj.lookAt
}
