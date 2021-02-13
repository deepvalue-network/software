package opengl

import (
	"fmt"
	"time"

	"github.com/go-gl/gl/v2.1/gl"
	uuid "github.com/satori/go.uuid"
	"github.com/deepvalue-network/software/treedee/domain/worlds/viewports"
)

type material struct {
	id       *uuid.UUID
	alpha    Alpha
	viewport viewports.Viewport
	layers   []Layer
}

func createMaterial(
	id *uuid.UUID,
	alpha Alpha,
	viewport viewports.Viewport,
	layers []Layer,
) Material {
	out := material{
		id:       id,
		alpha:    alpha,
		viewport: viewport,
		layers:   layers,
	}

	return &out
}

// ID returns the id
func (obj *material) ID() *uuid.UUID {
	return obj.id
}

// Alpha returns the alpha
func (obj *material) Alpha() Alpha {
	return obj.alpha
}

// Viewport returns the viewport
func (obj *material) Viewport() viewports.Viewport {
	return obj.viewport
}

// Layers returns the layers
func (obj *material) Layers() []Layer {
	return obj.layers
}

// Render renders a material
func (obj *material) Render(
	delta time.Duration,
	pos Position,
	orientation Orientation,
	activeScene Scene,
	program uint32,
) error {
	// use the program:
	gl.UseProgram(program)

	// loop the layers:
	for _, oneLayer := range obj.layers {
		err := oneLayer.Render(delta, pos, orientation, activeScene, program)
		if err != nil {
			return err
		}
	}

	// fetch the unform variable on the alpha, and update it:
	alphaValue := obj.alpha.Value()
	alphaVar := obj.alpha.Variable()
	alphaVarName := fmt.Sprintf(glStrPattern, alphaVar)
	alphaVarUniform := gl.GetUniformLocation(program, gl.Str(alphaVarName))
	gl.Uniform1f(alphaVarUniform, alphaValue)

	// update the viewport:
	rect := obj.viewport.Rectangle()
	rectPos := rect.Position()
	rectDim := rect.Dimension()
	viewportVar := obj.viewport.Variable()
	viewportVarname := fmt.Sprintf(glStrPattern, viewportVar)
	viewportUiform := gl.GetUniformLocation(program, gl.Str(viewportVarname))
	gl.Uniform4i(viewportUiform, int32(rectPos.X()), int32(rectPos.Y()), int32(rectDim.X()), int32(rectDim.Y()))

	return nil
}
