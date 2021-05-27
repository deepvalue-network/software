package executables

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications"
	"github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/scripts"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	applicationAdapter := applications.NewAdapter()
	scriptAdapter := scripts.NewAdapter()
	builder := NewBuilder()
	return createAdapter(applicationAdapter, scriptAdapter, builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents an executable adapter
type Adapter interface {
	ToExecutable(parsed parsers.Executable) (Executable, error)
}

// Builder represents an executable builder
type Builder interface {
	Create() Builder
	WithApplication(application applications.Application) Builder
	WithScript(script scripts.Script) Builder
	Now() (Executable, error)
}

// Executable represents an executable
type Executable interface {
	IsApplication() bool
	Application() applications.Application
	IsScript() bool
	Script() scripts.Script
}
