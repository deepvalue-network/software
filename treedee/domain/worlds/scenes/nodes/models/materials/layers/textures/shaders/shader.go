package shaders

import uuid "github.com/satori/go.uuid"

type shader struct {
	id        *uuid.UUID
	code      string
	isDynamic bool
}

func createShader(
	id *uuid.UUID,
	code string,
	isDynamic bool,
) Shader {
	out := shader{
		id:        id,
		code:      code,
		isDynamic: isDynamic,
	}

	return &out
}

// ID returns the id
func (obj *shader) ID() *uuid.UUID {
	return obj.id
}

// Code returns the code
func (obj *shader) Code() string {
	return obj.code
}

// IsDynamic returns true if the shader is dynamic, false otherwise
func (obj *shader) IsDynamic() bool {
	return obj.isDynamic
}
