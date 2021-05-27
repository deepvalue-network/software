package middle

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	testableAdapter := testables.NewAdapter()
	languageAdapter := applications.NewAdapter()
	builder := NewBuilder()
	return createAdapter(testableAdapter, languageAdapter, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents a program adapter
type Adapter interface {
	ToProgram(parsed parsers.Program) (Program, error)
}

// Builder represents a program builder
type Builder interface {
	Create() Builder
	WithTestable(testable testables.Testable) Builder
	WithLanguage(language applications.Application) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	IsTestable() bool
	Testable() testables.Testable
	IsLanguage() bool
	Language() applications.Application
}
