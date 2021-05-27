package testables

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/languages/definitions"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	executableAdapter := executables.NewAdapter()
	languageAdapter := definitions.NewAdapter()
	builder := NewBuilder()
	return createAdapter(executableAdapter, languageAdapter, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents a testable adapter
type Adapter interface {
	ToTestable(parsed parsers.Testable) (Testable, error)
}

// Builder represents a testable builder
type Builder interface {
	Create() Builder
	WithExecutable(executable executables.Executable) Builder
	WithLanguage(language definitions.Definition) Builder
	Now() (Testable, error)
}

// Testable represents a testable
type Testable interface {
	IsExecutable() bool
	Executable() executables.Executable
	IsLanguage() bool
	Language() definitions.Definition
}
