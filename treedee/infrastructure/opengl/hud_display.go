package opengl

import (
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/viewports"
)

type hudDisplay struct {
	id       *uuid.UUID
	index    uint
	viewport viewports.Viewport
	cam      Camera
	mat      Material
}

func createHudDisplay(
	id *uuid.UUID,
	index uint,
	viewport viewports.Viewport,
	cam Camera,
) HudDisplay {
	return createHudDisplayInternally(id, index, viewport, cam, nil)
}

func createHudDisplayWithMaterial(
	id *uuid.UUID,
	index uint,
	viewport viewports.Viewport,
	cam Camera,
	mat Material,
) HudDisplay {
	return createHudDisplayInternally(id, index, viewport, cam, mat)
}

func createHudDisplayInternally(
	id *uuid.UUID,
	index uint,
	viewport viewports.Viewport,
	cam Camera,
	mat Material,
) HudDisplay {
	out := hudDisplay{
		id:       id,
		index:    index,
		viewport: viewport,
		cam:      cam,
		mat:      mat,
	}

	return &out
}

// ID returns the id
func (obj *hudDisplay) ID() *uuid.UUID {
	return obj.id
}

// Index returns the index
func (obj *hudDisplay) Index() uint {
	return obj.index
}

// Viewport returns the viewport
func (obj *hudDisplay) Viewport() viewports.Viewport {
	return obj.viewport
}

// Camera returns the camera
func (obj *hudDisplay) Camera() Camera {
	return obj.cam
}

// HasMaterial returns true if there is material, false otherwise
func (obj *hudDisplay) HasMaterial() bool {
	return obj.mat != nil
}

// Material returns the material, if any
func (obj *hudDisplay) Material() Material {
	return obj.mat
}
