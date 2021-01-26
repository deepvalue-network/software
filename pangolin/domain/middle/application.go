package middle

import (
	"github.com/steve-care-software/products/pangolin/domain/middle/instructions"
	"github.com/steve-care-software/products/pangolin/domain/middle/labels"
	"github.com/steve-care-software/products/pangolin/domain/middle/tests"
	"github.com/steve-care-software/products/pangolin/domain/middle/variables"
)

type application struct {
	name    string
	version string
	ins     instructions.Instructions
	tests   tests.Tests
	labels  labels.Labels
	vars    variables.Variables
	imports []External
	extends []External
}

func createApplication(
	name string,
	version string,
	ins instructions.Instructions,
	tests tests.Tests,
	labels labels.Labels,
	vars variables.Variables,
) Application {
	return createApplicationInternally(name, version, ins, tests, labels, vars, nil, nil)
}

func createApplicationWithImports(
	name string,
	version string,
	ins instructions.Instructions,
	tests tests.Tests,
	labels labels.Labels,
	vars variables.Variables,
	imports []External,
) Application {
	return createApplicationInternally(name, version, ins, tests, labels, vars, imports, nil)
}

func createApplicationWithExtends(
	name string,
	version string,
	ins instructions.Instructions,
	tests tests.Tests,
	labels labels.Labels,
	vars variables.Variables,
	extends []External,
) Application {
	return createApplicationInternally(name, version, ins, tests, labels, vars, nil, extends)
}

func createApplicationWithImportsAndExtends(
	name string,
	version string,
	ins instructions.Instructions,
	tests tests.Tests,
	labels labels.Labels,
	vars variables.Variables,
	imports []External,
	extends []External,
) Application {
	return createApplicationInternally(name, version, ins, tests, labels, vars, imports, extends)
}

func createApplicationInternally(
	name string,
	version string,
	ins instructions.Instructions,
	tests tests.Tests,
	labels labels.Labels,
	vars variables.Variables,
	imports []External,
	extends []External,
) Application {
	out := application{
		name:    name,
		version: version,
		ins:     ins,
		tests:   tests,
		labels:  labels,
		vars:    vars,
		imports: imports,
		extends: extends,
	}

	return &out
}

// Name returns the name
func (obj *application) Name() string {
	return obj.name
}

// Version returns the version
func (obj *application) Version() string {
	return obj.version
}

// Instructions returns the instructions
func (obj *application) Instructions() instructions.Instructions {
	return obj.ins
}

// Tests returns the tests
func (obj *application) Tests() tests.Tests {
	return obj.tests
}

// Labels returns the labels
func (obj *application) Labels() labels.Labels {
	return obj.labels
}

// Variables returns the variables
func (obj *application) Variables() variables.Variables {
	return obj.vars
}

// HasImports returns true if there is imports, false otherwise
func (obj *application) HasImports() bool {
	return obj.imports != nil
}

// Imports returns the imports, if any
func (obj *application) Imports() []External {
	return obj.imports
}

// HasExtends returns true if there is extends, false otherwise
func (obj *application) HasExtends() bool {
	return obj.extends != nil
}

// Extends returns the extends, if any
func (obj *application) Extends() []External {
	return obj.extends
}
