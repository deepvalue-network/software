package opengl

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type textureShader struct {
	id        *uuid.UUID
	program   uint32
	isDynamic bool
}

func createTextureShader(
	id *uuid.UUID,
	program uint32,
	isDynamic bool,
) TextureShader {
	out := textureShader{
		id:        id,
		program:   program,
		isDynamic: isDynamic,
	}

	return &out
}

// ID returns the id
func (obj *textureShader) ID() *uuid.UUID {
	return obj.id
}

// Program returns the program
func (obj *textureShader) Program() uint32 {
	return obj.program
}

// IsDynamic returns true if the textureShader is dynamic, false otehrwise
func (obj *textureShader) IsDynamic() bool {
	return obj.isDynamic
}

// Render renders a texture shader
func (obj *textureShader) Render(
	delta time.Duration,
	pos Position,
	orientation Orientation,
	program uint32,
) error {
	return nil
}
