package languages

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/definitions"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents a language adapter
type Adapter interface {
	ToLanguage(parsed parsers.Language) (Language, error)
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
