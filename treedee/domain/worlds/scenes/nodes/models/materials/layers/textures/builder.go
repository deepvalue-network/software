package textures

import (
	"errors"
	"fmt"
	"math"

	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/math/ints"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/cameras"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials/layers/textures/pixels"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials/layers/textures/shaders"
)

type builder struct {
	id       *uuid.UUID
	dim      uint
	pixels   []pixels.Pixel
	variable string
	camera   cameras.Camera
	shader   shaders.Shader
}

func createBuilder() Builder {
	out := builder{
		id:       nil,
		dim:      0,
		pixels:   nil,
		variable: "",
		camera:   nil,
		shader:   nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithID adds an ID to the builder
func (app *builder) WithID(id *uuid.UUID) Builder {
	app.id = id
	return app
}

// WithDimension adds a dimension to the builder
func (app *builder) WithDimension(dim uint) Builder {
	app.dim = dim
	return app
}

// WithPixels add pixels to the builder
func (app *builder) WithPixels(pixels []pixels.Pixel) Builder {
	app.pixels = pixels
	return app
}

// WithVariable adds a variable to the builder
func (app *builder) WithVariable(variable string) Builder {
	app.variable = variable
	return app
}

// WithCamera adds a camera to the builder
func (app *builder) WithCamera(camera cameras.Camera) Builder {
	app.camera = camera
	return app
}

// WithShader adds a shader to the builder
func (app *builder) WithShader(shader shaders.Shader) Builder {
	app.shader = shader
	return app
}

// Now builds a new Texture instance
func (app *builder) Now() (Texture, error) {
	if app.id == nil {
		return nil, errors.New("the ID is mandatory in order to build a Texture instance")
	}

	if app.variable == "" {
		return nil, errors.New("the variable is mandatory in order to build a Texture instance")
	}

	pow := math.Pow(2.0, float64(app.dim))
	dim := ints.Vec2{
		int(pow),
		int(pow),
	}

	if app.pixels != nil && len(app.pixels) <= 0 {
		app.pixels = nil
	}

	if app.pixels != nil && app.variable != "" {
		amount := dim.X() * dim.Y()
		if amount != len(app.pixels) {
			str := fmt.Sprintf("the texture (%dx%d) was expected to cointain %d pixels, %d provided", dim.X(), dim.Y(), amount, len(app.pixels))
			return nil, errors.New(str)
		}

		return createTextureWithPixels(app.id, dim, app.variable, app.pixels), nil
	}

	if app.camera != nil {
		return createTextureWithCamera(app.id, dim, app.variable, app.camera), nil
	}

	if app.shader != nil {
		return createTextureWithShader(app.id, dim, app.variable, app.shader), nil
	}

	return nil, errors.New("the Texture is invalid")
}
