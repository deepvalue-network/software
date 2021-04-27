package linkers

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/tests"
)

type application struct {
	name      string
	version   string
	ins       instructions.Instructions
	tests     tests.Tests
	lbls      labels.Labels
	variables []variable.Variable
	imports   []External
	mp        map[string]Application
}

func createApplication(
	name string,
	version string,
	ins instructions.Instructions,
	tests tests.Tests,
	lbls labels.Labels,
	variables []variable.Variable,
) Application {
	return createApplicationInternally(name, version, ins, tests, lbls, variables, nil, map[string]Application{})
}

func createApplicationWithImports(
	name string,
	version string,
	ins instructions.Instructions,
	tests tests.Tests,
	lbls labels.Labels,
	variables []variable.Variable,
	imports []External,
	mp map[string]Application,
) Application {
	return createApplicationInternally(name, version, ins, tests, lbls, variables, imports, mp)
}

func createApplicationInternally(
	name string,
	version string,
	ins instructions.Instructions,
	tests tests.Tests,
	lbls labels.Labels,
	variables []variable.Variable,
	imports []External,
	mp map[string]Application,
) Application {
	out := application{
		name:      name,
		version:   version,
		ins:       ins,
		tests:     tests,
		lbls:      lbls,
		variables: variables,
		imports:   imports,
		mp:        mp,
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
	return obj.lbls
}

// Variables returns the variables
func (obj *application) Variables() []variable.Variable {
	return obj.variables
}

// HasImports returns true if there is imports, false otherwise
func (obj *application) HasImports() bool {
	return obj.imports != nil
}

// Imports returns the imports, if any
func (obj *application) Imports() []External {
	return obj.imports
}

// Import returns an imported application by name, if any
func (obj *application) Import(name string) (Application, error) {
	if app, ok := obj.mp[name]; ok {
		return app, nil
	}

	str := fmt.Sprintf("the name (%s) is not a valid imported application", name)
	return nil, errors.New(str)
}
