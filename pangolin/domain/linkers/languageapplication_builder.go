package linkers

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/tests"
)

type languageApplicationBuilder struct {
	name    string
	version string
	ins     instructions.Instructions
	tests   tests.Tests
	lbls    labels.Labels
	imports []External
}

func createLanguageApplicationBuilder() LanguageApplicationBuilder {
	out := languageApplicationBuilder{
		name:    "",
		version: "",
		ins:     nil,
		tests:   nil,
		lbls:    nil,
		imports: nil,
	}

	return &out
}

// Create initializes the builder
func (app *languageApplicationBuilder) Create() LanguageApplicationBuilder {
	return createLanguageApplicationBuilder()
}

// WithName adds a name to the builder
func (app *languageApplicationBuilder) WithName(name string) LanguageApplicationBuilder {
	app.name = name
	return app
}

// WithInstructions add instructions to the builder
func (app *languageApplicationBuilder) WithInstructions(ins instructions.Instructions) LanguageApplicationBuilder {
	app.ins = ins
	return app
}

// WithTests add tests to the builder
func (app *languageApplicationBuilder) WithTests(tests tests.Tests) LanguageApplicationBuilder {
	app.tests = tests
	return app
}

// WithLabels add labels to the builder
func (app *languageApplicationBuilder) WithLabels(lbls labels.Labels) LanguageApplicationBuilder {
	app.lbls = lbls
	return app
}

// WithImports add external imports to the builder
func (app *languageApplicationBuilder) WithImports(imports []External) LanguageApplicationBuilder {
	app.imports = imports
	return app
}

// WithVersion adds a version to the builder
func (app *languageApplicationBuilder) WithVersion(version string) LanguageApplicationBuilder {
	app.version = version
	return app
}

// Now builds a new Application instance
func (app *languageApplicationBuilder) Now() (LanguageApplication, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a LanguageApplication instance")
	}

	if app.version == "" {
		return nil, errors.New("the version is mandatory in order to build a LanguageApplication instance")
	}

	if app.ins == nil {
		return nil, errors.New("the instructions are mandatory in order to build a LanguageApplication instance")
	}

	if app.tests == nil {
		return nil, errors.New("the tests are mandatory in order to build a LanguageApplication instance")
	}

	if app.lbls == nil {
		return nil, errors.New("the labels are mandatory in order to build a LanguageApplication instance")
	}

	if app.imports != nil {
		mp := map[string]Application{}
		for _, oneImport := range app.imports {
			mp[oneImport.Name()] = oneImport.Application()
		}

		return createLanguageApplicationWithImports(
			app.name,
			app.version,
			app.ins,
			app.tests,
			app.lbls,
			app.imports,
			mp,
		), nil
	}

	return createLanguageApplication(
		app.name,
		app.version,
		app.ins,
		app.tests,
		app.lbls,
	), nil
}
