package middle

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages"
	"github.com/deepvalue-network/software/pangolin/domain/middle/scripts"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

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
	WithApplication(app applications.Application) Builder
	WithLanguage(lang languages.Language) Builder
	WithScript(script scripts.Script) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	IsApplication() bool
	Application() applications.Application
	IsLanguage() bool
	Language() languages.Language
	IsScript() bool
	Script() scripts.Script
}
