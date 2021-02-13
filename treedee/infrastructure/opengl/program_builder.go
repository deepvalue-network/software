package opengl

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-gl/gl/v2.1/gl"
	uuid "github.com/satori/go.uuid"
	vertex_shaders "github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/geometries/shaders"
	texture_shaders "github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials/layers/textures/shaders"
	fragment_shader "github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models/materials/shaders"
)

type programBuilder struct {
	texShader  texture_shaders.Shader
	verShader  vertex_shaders.Shader
	fragShader fragment_shader.Shader
}

func createProgramBuilder() ProgramBuilder {
	out := programBuilder{
		texShader:  nil,
		verShader:  nil,
		fragShader: nil,
	}

	return &out
}

// Create initializes the builder
func (app *programBuilder) Create() ProgramBuilder {
	return createProgramBuilder()
}

// WithTextureShader adds a texture shader to the builder
func (app *programBuilder) WithTextureShader(texShader texture_shaders.Shader) ProgramBuilder {
	app.texShader = texShader
	return app
}

// WithVertexShader adds a vertex shader to the builder
func (app *programBuilder) WithVertexShader(verShader vertex_shaders.Shader) ProgramBuilder {
	app.verShader = verShader
	return app
}

// WithFragmentShader adds a fragment shader to the builder
func (app *programBuilder) WithFragmentShader(fragShader fragment_shader.Shader) ProgramBuilder {
	app.fragShader = fragShader
	return app
}

// Now builds a new Program instance
func (app *programBuilder) Now() (uint32, error) {
	// create program:
	program := gl.CreateProgram()

	// if there is a texture shader:
	if app.texShader != nil {
		shaderID := app.texShader.ID()
		code := app.texShader.Code()
		shader, err := app.compileAny(shaderID, code, false, true)
		if err != nil {
			return 0, err
		}

		return app.linkThenClean(program, []uint32{
			shader,
		})
	}

	if app.verShader != nil && app.fragShader != nil {
		verShaderID := app.verShader.ID()
		verShaderCode := app.verShader.Code()
		verShader, err := app.compileAny(verShaderID, verShaderCode, true, false)
		if err != nil {
			return 0, err
		}

		fragShaderID := app.fragShader.ID()
		fragShaderCode := app.fragShader.Code()
		fragShader, err := app.compileAny(fragShaderID, fragShaderCode, false, true)
		if err != nil {
			return 0, err
		}

		return app.linkThenClean(program, []uint32{
			verShader,
			fragShader,
		})
	}

	// return the program:
	return program, nil
}

func (app *programBuilder) linkThenClean(program uint32, shaders []uint32) (uint32, error) {
	// attach all compiled shaders:
	for _, oneShader := range shaders {
		gl.AttachShader(program, oneShader)
	}

	// link the program:
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		str := fmt.Sprintf("there was a problem while linking the program: %s", log)
		return 0, errors.New(str)
	}

	// delete compiled shaders:
	for _, oneShader := range shaders {
		gl.DeleteShader(oneShader)
	}

	return program, nil
}

func (app *programBuilder) compileAny(id *uuid.UUID, code string, isVertex bool, isFragment bool) (uint32, error) {
	if isVertex {
		return app.compile(id, code, gl.VERTEX_SHADER)
	}

	if isFragment {
		return app.compile(id, code, gl.FRAGMENT_SHADER)
	}

	str := fmt.Sprintf("the shader (ID: %s) must be a vertex or fragment shader in order to build a program", id.String())
	return 0, errors.New(str)
}

func (app *programBuilder) compile(id *uuid.UUID, source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)
	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		str := fmt.Sprintf("the shader (ID: %s) could not compile\nsource: %v\n---\n log: %v", id.String(), source, log)
		return 0, errors.New(str)
	}

	return shader, nil
}
