package shaders

import uuid "github.com/satori/go.uuid"

type shader struct {
	id   *uuid.UUID
	code string
}

func createShader(
	id *uuid.UUID,
	code string,
) Shader {
	out := shader{
		id:   id,
		code: code,
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
