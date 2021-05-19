package interpreters

import (
	"github.com/deepvalue-network/software/pangolin/domain/interpreters/stackframes"
	"github.com/deepvalue-network/software/pangolin/domain/linkers"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
)

// Program represents a program interpreter
type Program interface {
	Script() Script
	Language() Language
	Application() Application
}

// Script represents a script interpreter
type Script interface {
	Execute(script linkers.Script) (linkers.Application, error)
}

// Language represents a language interpreter
type Language interface {
	TestByNames(names []string) error
	TestsAll() error
}

// Application represents an application interpreter
type Application interface {
	Execute(input map[string]computable.Value) (stackframes.StackFrame, error)
	TestsAll() error
	TestByNames(names []string) error
}
