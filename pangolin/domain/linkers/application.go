package linkers

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/labels"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/tests"
)

type application struct {
	name    string
	version string
	ins     instructions.Instructions
	tests   tests.Tests
	lbls    labels.Labels
	imports []External
	mp      map[string]Executable
}

func createApplication(
	name string,
	version string,
	ins instructions.Instructions,
	tests tests.Tests,
	lbls labels.Labels,
) Application {
	return createApplicationInternally(name, version, ins, tests, lbls, nil, map[string]Executable{})
}

func createApplicationWithImports(
	name string,
	version string,
	ins instructions.Instructions,
	tests tests.Tests,
	lbls labels.Labels,
	imports []External,
	mp map[string]Executable,
) Application {
	return createApplicationInternally(name, version, ins, tests, lbls, imports, mp)
}

func createApplicationInternally(
	name string,
	version string,
	ins instructions.Instructions,
	tests tests.Tests,
	lbls labels.Labels,
	imports []External,
	mp map[string]Executable,
) Application {
	out := application{
		name:    name,
		version: version,
		ins:     ins,
		tests:   tests,
		lbls:    lbls,
		imports: imports,
		mp:      mp,
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

// HasImports returns true if there is imports, false otherwise
func (obj *application) HasImports() bool {
	return obj.imports != nil
}

// Imports returns the imports, if any
func (obj *application) Imports() []External {
	return obj.imports
}

// Import returns an imported application by name, if any
func (obj *application) Import(name string) (Executable, error) {
	if app, ok := obj.mp[name]; ok {
		return app, nil
	}

	str := fmt.Sprintf("the name (%s) is not a valid imported executable", name)
	return nil, errors.New(str)
}
