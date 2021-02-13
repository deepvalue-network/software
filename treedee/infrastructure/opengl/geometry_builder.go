package opengl

import (
	"errors"
	"fmt"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/geometries"
)

type geometryBuilder struct {
	prog *uint32
	geo  geometries.Geometry
}

func createGeometryBuilder() GeometryBuilder {
	out := geometryBuilder{
		prog: nil,
		geo:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *geometryBuilder) Create() GeometryBuilder {
	return createGeometryBuilder()
}

// WithProgram adds a program to the builder
func (app *geometryBuilder) WithProgram(prog uint32) GeometryBuilder {
	app.prog = &prog
	return app
}

// WithGeometry adds a geometry to the builder
func (app *geometryBuilder) WithGeometry(geo geometries.Geometry) GeometryBuilder {
	app.geo = geo
	return app
}

// Now builds a new Geometry instance
func (app *geometryBuilder) Now() (Geometry, error) {
	if app.prog == nil {
		return nil, errors.New("the program is mandatory in order to build a Geometry instance")
	}

	if app.geo == nil {
		return nil, errors.New("the geometry is mandatory in order to build a Geometry instance")
	}

	// vertex shader variables:
	shader := app.geo.Shader()
	variables := shader.Variables()
	vertexCoordinatesVariable := variables.VertexCoordinates()
	textureCoordinatesVariable := variables.TextureCoordinates()
	vertexShaderVariables := createVertexShaderVariables(
		vertexCoordinatesVariable,
		textureCoordinatesVariable,
	)

	// vertex shader:
	shaderID := shader.ID()
	vertexShader := createVertexShader(shaderID, vertexShaderVariables)

	list := []float32{}
	vertices := app.geo.Vertices()
	for _, oneVertice := range vertices {
		pos := oneVertice.Space()
		tex := oneVertice.Texture()
		list = append(list, []float32{
			pos.X(),
			pos.Y(),
			pos.Z(),
			tex.X(),
			tex.Y(),
		}...)
	}

	verticesType := app.geo.Type()
	if verticesType.IsTriangle() {
		triSize := int32(3)
		texSize := int32(2)
		vertexSize := triSize + texSize
		stride := int32(vertexSize * float32SizeInBytes)

		var vao uint32
		gl.GenVertexArrays(1, &vao)
		gl.BindVertexArray(vao)

		var vbo uint32
		gl.GenBuffers(1, &vbo)
		gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
		gl.BufferData(gl.ARRAY_BUFFER, len(list)*float32SizeInBytes, gl.Ptr(list), gl.STATIC_DRAW)

		verOffset := int(0)
		vertexVarName := fmt.Sprintf(glStrPattern, vertexCoordinatesVariable)
		vertAttrib := uint32(gl.GetAttribLocation(*app.prog, gl.Str(vertexVarName)))
		gl.EnableVertexAttribArray(vertAttrib)
		gl.VertexAttribPointer(vertAttrib, triSize, gl.FLOAT, false, stride, gl.PtrOffset(verOffset))

		texOffset := int(triSize * float32SizeInBytes)
		texVarName := fmt.Sprintf(glStrPattern, textureCoordinatesVariable)
		texCoordAttrib := uint32(gl.GetAttribLocation(*app.prog, gl.Str(texVarName)))
		gl.EnableVertexAttribArray(texCoordAttrib)
		gl.VertexAttribPointer(texCoordAttrib, texSize, gl.FLOAT, false, stride, gl.PtrOffset(texOffset))

		// vertex type:
		vertexType := createVertexTypeWithTriangle()

		// geometry:
		id := app.geo.ID()
		vertexAmount := int32(len(vertices))
		return createGeometry(id, *app.prog, vao, vertexAmount, vertexType, vertexShader), nil
	}

	return nil, errors.New("the vertex type is invalid in order to build a Geometry instance")
}
