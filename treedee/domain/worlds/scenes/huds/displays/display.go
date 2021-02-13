package displays

import (
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/cameras"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials"
	"github.com/deepvalue-network/software/treedee/domain/worlds/viewports"
)

type display struct {
	id       *uuid.UUID
	index    uint
	viewport viewports.Viewport
	camera   cameras.Camera
	mat      materials.Material
}

func createDisplay(
	id *uuid.UUID,
	index uint,
	viewport viewports.Viewport,
	camera cameras.Camera,
) Display {
	return createDisplayInternally(id, index, viewport, camera, nil)
}

func createDisplayWithMaterial(
	id *uuid.UUID,
	index uint,
	viewport viewports.Viewport,
	camera cameras.Camera,
	mat materials.Material,
) Display {
	return createDisplayInternally(id, index, viewport, camera, mat)
}

func createDisplayInternally(
	id *uuid.UUID,
	index uint,
	viewport viewports.Viewport,
	camera cameras.Camera,
	mat materials.Material,
) Display {
	out := display{
		id:       id,
		index:    index,
		viewport: viewport,
		camera:   camera,
		mat:      mat,
	}

	return &out
}

// ID returns the id
func (obj *display) ID() *uuid.UUID {
	return obj.id
}

// Index returns the index
func (obj *display) Index() uint {
	return obj.index
}

// Viewport returns the viewport
func (obj *display) Viewport() viewports.Viewport {
	return obj.viewport
}

// Camera returns the camera
func (obj *display) Camera() cameras.Camera {
	return obj.camera
}

// HasMaterial returns true if there is material, false otherwise
func (obj *display) HasMaterial() bool {
	return obj.mat != nil
}

// Material returns the material, if any
func (obj *display) Material() materials.Material {
	return obj.mat
}
