package scripts

import "github.com/deepvalue-network/software/pangolin/domain/parsers"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents an adapter
type Adapter interface {
	ToScript(parsed parsers.ScriptCommand) (Script, error)
}

// Builder represents a script builder
type Builder interface {
	Create() Builder
	WithVariable(variable string) Builder
	WithValues(values []Value) Builder
	Now() (Script, error)
}

// Script represents a script command
type Script interface {
	Variable() string
	Values() []Value
}

// ValueBuilder represents a script value builder
type ValueBuilder interface {
	Create() ValueBuilder
	WithName(name string) ValueBuilder
	WithVersion(version string) ValueBuilder
	WithLanguagePath(langPath string) ValueBuilder
	WithScriptPath(scriptPath string) ValueBuilder
	Now() (Value, error)
}

// Value represents a script value
type Value interface {
	IsName() bool
	Name() string
	IsVersion() bool
	Version() string
	IsLanguagePath() bool
	LanguagePath() string
	IsScriptPath() bool
	ScriptPath() string
}
