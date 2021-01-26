package middle

import (
	"errors"

	"github.com/steve-care-software/products/pangolin/domain/middle/instructions"
	"github.com/steve-care-software/products/pangolin/domain/middle/labels"
	"github.com/steve-care-software/products/pangolin/domain/middle/tests"
	"github.com/steve-care-software/products/pangolin/domain/middle/variables"
)

type applicationBuilder struct {
	instructionsBuilder instructions.Builder
	labelsBuilder       labels.Builder
	variablesBuilder    variables.Builder
	testsBuilder        tests.Builder
	name                string
	version             string
	ins                 instructions.Instructions
	tests               tests.Tests
	labels              labels.Labels
	vars                variables.Variables
	imports             []External
	extends             []External
}

func createApplicationBuilder(
	instructionsBuilder instructions.Builder,
	labelsBuilder labels.Builder,
	variablesBuilder variables.Builder,
	testsBuilder tests.Builder,
) ApplicationBuilder {
	out := applicationBuilder{
		instructionsBuilder: instructionsBuilder,
		labelsBuilder:       labelsBuilder,
		variablesBuilder:    variablesBuilder,
		testsBuilder:        testsBuilder,
		name:                "",
		version:             "",
		ins:                 nil,
		tests:               nil,
		labels:              nil,
		vars:                nil,
		imports:             nil,
		extends:             nil,
	}

	return &out
}

// Create initializes the applicationBuilder
func (app *applicationBuilder) Create() ApplicationBuilder {
	return createApplicationBuilder(
		app.instructionsBuilder,
		app.labelsBuilder,
		app.variablesBuilder,
		app.testsBuilder,
	)
}

// WithName adds a name to the applicationBuilder
func (app *applicationBuilder) WithName(name string) ApplicationBuilder {
	app.name = name
	return app
}

// WithVersion adds a version to the applicationBuilder
func (app *applicationBuilder) WithVersion(version string) ApplicationBuilder {
	app.version = version
	return app
}

// WithImports add imports to the applicationBuilder
func (app *applicationBuilder) WithImports(imports []External) ApplicationBuilder {
	app.imports = imports
	return app
}

// WithExtends add extends to the applicationBuilder
func (app *applicationBuilder) WithExtends(extends []External) ApplicationBuilder {
	app.extends = extends
	return app
}

// WithInstructions add instructions to the applicationBuilder
func (app *applicationBuilder) WithInstructions(instructions instructions.Instructions) ApplicationBuilder {
	app.ins = instructions
	return app
}

// WithTests add tests to the applicationBuilder
func (app *applicationBuilder) WithTests(tests tests.Tests) ApplicationBuilder {
	app.tests = tests
	return app
}

// WithLabels add labels to the applicationBuilder
func (app *applicationBuilder) WithLabels(labels labels.Labels) ApplicationBuilder {
	app.labels = labels
	return app
}

// WithVariables adds a variables to the applicationBuilder
func (app *applicationBuilder) WithVariables(variables variables.Variables) ApplicationBuilder {
	app.vars = variables
	return app
}

// Now builds a new Application instance
func (app *applicationBuilder) Now() (Application, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Application instance")
	}

	if app.version == "" {
		return nil, errors.New("the version is mandatory in order to build a Application instance")
	}

	if app.ins == nil {
		ins, err := app.instructionsBuilder.Create().Now()
		if err != nil {
			return nil, err
		}

		app.ins = ins
	}

	if app.labels == nil {
		labels, err := app.labelsBuilder.Create().Now()
		if err != nil {
			return nil, err
		}

		app.labels = labels
	}

	if app.vars == nil {
		vars, err := app.variablesBuilder.Create().Now()
		if err != nil {
			return nil, err
		}

		app.vars = vars
	}

	if app.tests == nil {
		tests, err := app.testsBuilder.Create().Now()
		if err != nil {
			return nil, err
		}

		app.tests = tests
	}

	if app.imports != nil && app.extends != nil {
		return createApplicationWithImportsAndExtends(
			app.name,
			app.version,
			app.ins,
			app.tests,
			app.labels,
			app.vars,
			app.imports,
			app.extends,
		), nil
	}

	if app.imports != nil {
		return createApplicationWithImports(
			app.name,
			app.version,
			app.ins,
			app.tests,
			app.labels,
			app.vars,
			app.imports,
		), nil
	}

	if app.extends != nil {
		return createApplicationWithExtends(
			app.name,
			app.version,
			app.ins,
			app.tests,
			app.labels,
			app.vars,
			app.extends,
		), nil
	}

	return createApplication(
		app.name,
		app.version,
		app.ins,
		app.tests,
		app.labels,
		app.vars,
	), nil
}
