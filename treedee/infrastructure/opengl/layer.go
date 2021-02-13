package opengl

import (
	"fmt"
	"time"

	"github.com/go-gl/gl/v2.1/gl"
	uuid "github.com/satori/go.uuid"
)

type layer struct {
	id    *uuid.UUID
	index uint
	alpha Alpha
	tex   Texture
}

func createLayer(
	id *uuid.UUID,
	index uint,
	alpha Alpha,
	tex Texture,
) Layer {
	out := layer{
		id:    id,
		index: index,
		alpha: alpha,
		tex:   tex,
	}

	return &out
}

// ID returns the id
func (obj *layer) ID() *uuid.UUID {
	return obj.id
}

// Index returns the index
func (obj *layer) Index() uint {
	return obj.index
}

// Alpha returns the alpha
func (obj *layer) Alpha() Alpha {
	return obj.alpha
}

// Texture returns the texture
func (obj *layer) Texture() Texture {
	return obj.tex
}

// Render renders a texture
func (obj *layer) Render(
	delta time.Duration,
	pos Position,
	orientation Orientation,
	activeScene Scene,
	program uint32,
) error {
	// render the texture:
	err := obj.Texture().Render(delta, pos, orientation, activeScene, program)
	if err != nil {
		return err
	}

	// fetch the unform variable on the alpha, and update it:
	alphaValue := obj.alpha.Value()
	alphaVar := obj.alpha.Variable()
	alphaVarName := fmt.Sprintf(glStrPattern, alphaVar)
	alphaVarUniform := gl.GetUniformLocation(program, gl.Str(alphaVarName))
	gl.Uniform1f(alphaVarUniform, alphaValue)

	return nil
}
