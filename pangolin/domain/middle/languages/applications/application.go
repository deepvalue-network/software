package applications

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/heads"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/tests"
)

type application struct {
	head   heads.Head
	labels labels.Labels
	main   instructions.Instructions
	tests  tests.Tests
}

func createApplication(
	head heads.Head,
	labels labels.Labels,
	main instructions.Instructions,
) Application {
	return createApplicationInternally(head, labels, main, nil)
}

func createApplicationWithTests(
	head heads.Head,
	labels labels.Labels,
	main instructions.Instructions,
	tests tests.Tests,
) Application {
	return createApplicationInternally(head, labels, main, tests)
}

func createApplicationInternally(
	head heads.Head,
	labels labels.Labels,
	main instructions.Instructions,
	tests tests.Tests,
) Application {
	out := application{
		head:   head,
		labels: labels,
		main:   main,
		tests:  tests,
	}

	return &out
}

// Head returns the head
func (obj *application) Head() heads.Head {
	return obj.head
}

// Labels returns the labels
func (obj *application) Labels() labels.Labels {
	return obj.labels
}

// Main returns the main
func (obj *application) Main() instructions.Instructions {
	return obj.main
}

// HasTests returns true if there is tests, false otherwise
func (obj *application) HasTests() bool {
	return obj.tests != nil
}

// Tests returns tests, if any
func (obj *application) Tests() tests.Tests {
	return obj.tests
}
