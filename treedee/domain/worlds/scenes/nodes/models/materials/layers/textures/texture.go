package textures

import (
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/math/ints"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/cameras"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials/layers/textures/pixels"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials/layers/textures/shaders"
)

type texture struct {
	id       *uuid.UUID
	dim      ints.Vec2
	variable string
	pix      []pixels.Pixel
	cam      cameras.Camera
	shader   shaders.Shader
}

func createTextureWithPixels(
	id *uuid.UUID,
	dim ints.Vec2,
	variable string,
	pix []pixels.Pixel,
) Texture {
	return createTextureInternally(id, dim, variable, pix, nil, nil)
}

func createTextureWithCamera(
	id *uuid.UUID,
	dim ints.Vec2,
	variable string,
	cam cameras.Camera,
) Texture {
	return createTextureInternally(id, dim, variable, nil, cam, nil)
}

func createTextureWithShader(
	id *uuid.UUID,
	dim ints.Vec2,
	variable string,
	shader shaders.Shader,
) Texture {
	return createTextureInternally(id, dim, variable, nil, nil, shader)
}

func createTextureInternally(
	id *uuid.UUID,
	dim ints.Vec2,
	variable string,
	pix []pixels.Pixel,
	cam cameras.Camera,
	shader shaders.Shader,
) Texture {
	out := texture{
		id:       id,
		dim:      dim,
		variable: variable,
		pix:      pix,
		cam:      cam,
		shader:   shader,
	}

	return &out
}

// ID returns the id
func (obj *texture) ID() *uuid.UUID {
	return obj.id
}

// Dimension returns the dimension
func (obj *texture) Dimension() ints.Vec2 {
	return obj.dim
}

// Variable returns the variable
func (obj *texture) Variable() string {
	return obj.variable
}

// IsPixels returns true if there is pixels, false otherwise
func (obj *texture) IsPixels() bool {
	return obj.pix != nil
}

// Pixels returns the pixels, if any
func (obj *texture) Pixels() []pixels.Pixel {
	return obj.pix
}

// IsCamera returns true if there a camera, false otherwise
func (obj *texture) IsCamera() bool {
	return obj.cam != nil
}

// Camera returns a camera, if any
func (obj *texture) Camera() cameras.Camera {
	return obj.cam
}

// IsShader returns true if there a shader, false otherwise
func (obj *texture) IsShader() bool {
	return obj.shader != nil
}

// Shader returns the shader, if any
func (obj *texture) Shader() shaders.Shader {
	return obj.shader
}
