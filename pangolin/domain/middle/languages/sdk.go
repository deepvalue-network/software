package languages

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/definitions"
)

// Language represents a language
type Language interface {
	IsDefinition() bool
	Definition() definitions.Definition
	IsApplication() bool
	Application() applications.Application
}
