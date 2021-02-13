package opengl

import (
	"time"

	"github.com/go-gl/gl/v4.6-core/gl"
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/math/ints"
)

type texture struct {
	id        *uuid.UUID
	dimension ints.Vec2
	variable  string
	res       uint32
	cam       Camera
	shader    TextureShader
}

func createTextureWithResource(
	id *uuid.UUID,
	dimension ints.Vec2,
	variable string,
	res uint32,
) Texture {
	return createTextureInternally(id, dimension, variable, res, nil, nil)
}

func createTextureWithCamera(
	id *uuid.UUID,
	dimension ints.Vec2,
	variable string,
	cam Camera,
) Texture {
	return createTextureInternally(id, dimension, variable, 0, cam, nil)
}

func createTextureWithShader(
	id *uuid.UUID,
	dimension ints.Vec2,
	variable string,
	shader TextureShader,
) Texture {
	return createTextureInternally(id, dimension, variable, 0, nil, shader)
}

func createTextureInternally(
	id *uuid.UUID,
	dimension ints.Vec2,
	variable string,
	res uint32,
	cam Camera,
	shader TextureShader,
) Texture {
	out := texture{
		id:        id,
		dimension: dimension,
		variable:  variable,
		res:       res,
		cam:       cam,
		shader:    shader,
	}

	return &out
}

// ID returns the id
func (obj *texture) ID() *uuid.UUID {
	return obj.id
}

// Dimension returns the dimension
func (obj *texture) Dimension() ints.Vec2 {
	return obj.dimension
}

// Variable returns the variable
func (obj *texture) Variable() string {
	return obj.variable
}

// IsResource returns true if there is a resource, false otherwise
func (obj *texture) IsResource() bool {
	return obj.res != 0
}

// Resource returns the resource, if any
func (obj *texture) Resource() uint32 {
	return obj.res
}

// IsCamera returns true if there is a camera, false otherwise
func (obj *texture) IsCamera() bool {
	return obj.cam != nil
}

// Camera returns the camera, if any
func (obj *texture) Camera() Camera {
	return obj.cam
}

// IsShader returns true if there is a shader, false otherwise
func (obj *texture) IsShader() bool {
	return obj.shader != nil
}

// Shader returns the shader, if any
func (obj *texture) Shader() TextureShader {
	return obj.shader
}

// Render renders a texture
func (obj *texture) Render(
	delta time.Duration,
	pos Position,
	orientation Orientation,
	activeScene Scene,
	program uint32,
) error {
	if obj.IsResource() {
		// use the texture:
		gl.ActiveTexture(gl.TEXTURE0)
		gl.BindTexture(gl.TEXTURE_2D, obj.Resource())
		return nil
	}

	if obj.IsShader() {
		return obj.Shader().Render(delta, pos, orientation, program)
	}

	return obj.Camera().Render(delta, pos, orientation, activeScene, program)
}
