package languages

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/definitions"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a language builder
type Builder interface {
	Create() Builder
	WithDefinition(def definitions.Definition) Builder
	WithApplication(app applications.Application) Builder
	Now() (Language, error)
}

// Language represents a language
type Language interface {
	IsDefinition() bool
	Definition() definitions.Definition
	IsApplication() bool
	Application() applications.Application
}
