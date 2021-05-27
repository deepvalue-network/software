package applications

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/heads"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/labels"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/tests"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type builder struct {
	head    heads.Head
	main    instructions.Instructions
	tests   tests.Tests
	labels  labels.Labels
	extends []parsers.ImportSingle
}

func createBuilder() Builder {
	out := builder{
		head:    nil,
		main:    nil,
		tests:   nil,
		labels:  nil,
		extends: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithHead adds an head to the builder
func (app *builder) WithHead(head heads.Head) Builder {
	app.head = head
	return app
}

// WithMain add main instructions to the builder
func (app *builder) WithMain(main instructions.Instructions) Builder {
	app.main = main
	return app
}

// WithTests add tests to the builder
func (app *builder) WithTests(tests tests.Tests) Builder {
	app.tests = tests
	return app
}

// WithLabels add labels to the builder
func (app *builder) WithLabels(labels labels.Labels) Builder {
	app.labels = labels
	return app
}

// WithExtends add extends to the builder
func (app *builder) WithExtends(extends []parsers.ImportSingle) Builder {
	app.extends = extends
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.head == nil {
		return nil, errors.New("the head is mandatory in order to build an Application instance")
	}

	if app.main == nil {
		return nil, errors.New("the main instructions is mandatory in order to build an Application instance")
	}

	if app.tests == nil {
		return nil, errors.New("the tests is mandatory in order to build an Application instance")
	}

	if app.labels == nil {
		return nil, errors.New("the labels is mandatory in order to build an Application instance")
	}

	if app.extends != nil && len(app.extends) <= 0 {
		app.extends = nil
	}

	if app.extends != nil {
		return createApplicationWithExtends(app.head, app.main, app.tests, app.labels, app.extends), nil
	}

	return createApplication(app.head, app.main, app.tests, app.labels), nil
}
