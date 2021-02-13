package opengl

import (
	"github.com/go-gl/gl/v2.1/gl"
	uuid "github.com/satori/go.uuid"
)

type geometry struct {
	id           *uuid.UUID
	prog         uint32
	vao          uint32
	vertexAmount int32
	vertexType   VertexType
	shader       VertexShader
}

func createGeometry(
	id *uuid.UUID,
	prog uint32,
	vao uint32,
	vertexAmount int32,
	vertexType VertexType,
	shader VertexShader,
) Geometry {
	out := geometry{
		id:           id,
		prog:         prog,
		vao:          vao,
		vertexAmount: vertexAmount,
		vertexType:   vertexType,
		shader:       shader,
	}

	return &out
}

// ID returns the id
func (obj *geometry) ID() *uuid.UUID {
	return obj.id
}

// Program returns the program
func (obj *geometry) Program() uint32 {
	return obj.prog
}

// VAO returns the vao
func (obj *geometry) VAO() uint32 {
	return obj.vao
}

// VertexAmount returns the vertex amount
func (obj *geometry) VertexAmount() int32 {
	return obj.vertexAmount
}

// VertexType returns the vertex type
func (obj *geometry) VertexType() VertexType {
	return obj.vertexType
}

// Shader returns the shader
func (obj *geometry) Shader() VertexShader {
	return obj.shader
}

// Prepare prepares the geometry to be rendered
func (obj *geometry) Prepare() error {
	// vao:
	vao := obj.VAO()
	gl.BindVertexArray(vao)

	// returns:
	return nil
}

// Render renders the geometry
func (obj *geometry) Render() error {
	// draw:
	amount := obj.VertexAmount()
	gl.DrawArrays(gl.TRIANGLES, 0, amount)

	// returns:
	return nil
}
