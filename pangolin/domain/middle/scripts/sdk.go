package scripts

import (
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	builder := NewBuilder()
	return createAdapter(builder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the script adapter
type Adapter interface {
	ToScript(parsed parsers.Script) (Script, error)
}

// Builder represents a script builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithVersion(version string) Builder
	WithLanguagePath(lang string) Builder
	WithScriptPath(script string) Builder
	Now() (Script, error)
}

// Script represents a script
type Script interface {
	Name() string
	Version() string
	LanguagePath() string
	ScriptPath() string
}
