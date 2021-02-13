package opengl

import (
	"errors"

	"github.com/deepvalue-network/software/treedee/domain/worlds/scenes/nodes/models"
)

type modelBuilder struct {
	programBuilder  ProgramBuilder
	materialBuilder MaterialBuilder
	geometryBuilder GeometryBuilder
	model           models.Model
}

func createModelBuilder(
	programBuilder ProgramBuilder,
	materialBuilder MaterialBuilder,
	geometryBuilder GeometryBuilder,
) ModelBuilder {
	out := modelBuilder{
		programBuilder:  programBuilder,
		materialBuilder: materialBuilder,
		geometryBuilder: geometryBuilder,
		model:           nil,
	}

	return &out
}

// Create initializes the builder
func (app *modelBuilder) Create() ModelBuilder {
	return createModelBuilder(
		app.programBuilder,
		app.materialBuilder,
		app.geometryBuilder,
	)
}

// WithModel adds a model to the builder
func (app *modelBuilder) WithModel(model models.Model) ModelBuilder {
	app.model = model
	return app
}

// Now builds a new Model instance
func (app *modelBuilder) Now() (Model, error) {
	if app.model == nil {
		return nil, errors.New("the model is mandatory in order to build a Model instance")
	}

	domainMat := app.model.Material()
	fragmentShader := domainMat.Shader()

	domainGeo := app.model.Geometry()
	vertexShader := domainGeo.Shader()
	program, err := app.programBuilder.Create().WithVertexShader(vertexShader).WithFragmentShader(fragmentShader).Now()
	if err != nil {
		return nil, err
	}

	mat, err := app.materialBuilder.Create().WithMaterial(domainMat).Now()
	if err != nil {
		return nil, err
	}

	geo, err := app.geometryBuilder.Create().WithProgram(program).WithGeometry(domainGeo).Now()
	if err != nil {
		return nil, err
	}

	id := app.model.ID()
	return createModel(id, program, geo, mat), nil
}
