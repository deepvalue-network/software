package linkers

import (
	"errors"
	"fmt"

	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/tests"
)

type languageApplication struct {
	name    string
	version string
	ins     instructions.Instructions
	tests   tests.Tests
	lbls    labels.Labels
	imports []External
	mp      map[string]Application
}

func createLanguageApplication(
	name string,
	version string,
	ins instructions.Instructions,
	tests tests.Tests,
	lbls labels.Labels,
) LanguageApplication {
	return createLanguageApplicationInternally(name, version, ins, tests, lbls, nil, map[string]Application{})
}

func createLanguageApplicationWithImports(
	name string,
	version string,
	ins instructions.Instructions,
	tests tests.Tests,
	lbls labels.Labels,
	imports []External,
	mp map[string]Application,
) LanguageApplication {
	return createLanguageApplicationInternally(name, version, ins, tests, lbls, imports, mp)
}

func createLanguageApplicationInternally(
	name string,
	version string,
	ins instructions.Instructions,
	tests tests.Tests,
	lbls labels.Labels,
	imports []External,
	mp map[string]Application,
) LanguageApplication {
	out := languageApplication{
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
func (obj *languageApplication) Name() string {
	return obj.name
}

// Version returns the version
func (obj *languageApplication) Version() string {
	return obj.version
}

// Instructions returns the instructions
func (obj *languageApplication) Instructions() instructions.Instructions {
	return obj.ins
}

// Tests returns the tests
func (obj *languageApplication) Tests() tests.Tests {
	return obj.tests
}

// Labels returns the labels
func (obj *languageApplication) Labels() labels.Labels {
	return obj.lbls
}

// HasImports returns true if there is imports, false otherwise
func (obj *languageApplication) HasImports() bool {
	return obj.imports != nil
}

// Imports returns the imports, if any
func (obj *languageApplication) Imports() []External {
	return obj.imports
}

// Import returns an imported languageApplication by name, if any
func (obj *languageApplication) Import(name string) (Application, error) {
	if app, ok := obj.mp[name]; ok {
		return app, nil
	}

	str := fmt.Sprintf("the name (%s) is not a valid imported application", name)
	return nil, errors.New(str)
}
