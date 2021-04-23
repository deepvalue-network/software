package applications

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/externals"
	"github.com/deepvalue-network/software/pangolin/domain/middle/heads"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/tests"
)

type application struct {
	head    heads.Head
	main    instructions.Instructions
	tests   tests.Tests
	labels  labels.Labels
	extends []externals.External
}

func createApplication(
	head heads.Head,
	main instructions.Instructions,
	tests tests.Tests,
	labels labels.Labels,
) Application {
	return createApplicationInternally(head, main, tests, labels, nil)
}

func createApplicationWithExtends(
	head heads.Head,
	main instructions.Instructions,
	tests tests.Tests,
	labels labels.Labels,
	extends []externals.External,
) Application {
	return createApplicationInternally(head, main, tests, labels, extends)
}

func createApplicationInternally(
	head heads.Head,
	main instructions.Instructions,
	tests tests.Tests,
	labels labels.Labels,
	extends []externals.External,
) Application {
	out := application{
		head:    head,
		main:    main,
		tests:   tests,
		labels:  labels,
		extends: extends,
	}

	return &out
}

// Head returns the head
func (obj *application) Head() heads.Head {
	return obj.head
}

// Main returns the main instructions
func (obj *application) Main() instructions.Instructions {
	return obj.main
}

// Tests returns the tests
func (obj *application) Tests() tests.Tests {
	return obj.tests
}

// Labels returns the labels
func (obj *application) Labels() labels.Labels {
	return obj.labels
}

// HasExtends returns true if there is an extends, false otherwise
func (obj *application) HasExtends() bool {
	return obj.extends != nil
}

// Extends returns the extends, if any
func (obj *application) Extends() []externals.External {
	return obj.extends
}
