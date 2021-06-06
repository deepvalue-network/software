package applications

import (
	"errors"

	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/heads"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/tests"
)

type builder struct {
	head   heads.Head
	labels labels.Labels
	main   instructions.Instructions
	tests  tests.Tests
}

func createBuilder() Builder {
	out := builder{
		head:   nil,
		labels: nil,
		main:   nil,
		tests:  nil,
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

// WithLabels add labels to the builder
func (app *builder) WithLabels(labels labels.Labels) Builder {
	app.labels = labels
	return app
}

// WithMain add main instructions to the builder
func (app *builder) WithMain(main instructions.Instructions) Builder {
	app.main = main
	return app
}

// WithTests add tests instructions to the builder
func (app *builder) WithTests(tests tests.Tests) Builder {
	app.tests = tests
	return app
}

// Now builds a new Application instance
func (app *builder) Now() (Application, error) {
	if app.head == nil {
		return nil, errors.New("the head is mandatory in order to build an Application instance")
	}

	if app.labels == nil {
		return nil, errors.New("the labels is mandatory in order to build an Application instance")
	}

	if app.main == nil {
		return nil, errors.New("the main instructions is mandatory in order to build an Application instance")
	}

	if app.tests != nil {
		return createApplicationWithTests(app.head, app.labels, app.main, app.tests), nil
	}

	return createApplication(app.head, app.labels, app.main), nil
}
