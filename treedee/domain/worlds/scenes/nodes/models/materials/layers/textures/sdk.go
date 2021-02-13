package textures

import (
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/math/ints"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/cameras"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials/layers/textures/pixels"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials/layers/textures/shaders"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a texture builder
type Builder interface {
	Create() Builder
	WithID(id *uuid.UUID) Builder
	WithDimension(dim uint) Builder
	WithPixels(pixels []pixels.Pixel) Builder
	WithVariable(variable string) Builder
	WithCamera(camera cameras.Camera) Builder
	WithShader(shader shaders.Shader) Builder
	Now() (Texture, error)
}

// Texture represents a texture
type Texture interface {
	ID() *uuid.UUID
	Dimension() ints.Vec2
	Variable() string
	IsPixels() bool
	Pixels() []pixels.Pixel
	IsCamera() bool
	Camera() cameras.Camera
	IsShader() bool
	Shader() shaders.Shader
}
