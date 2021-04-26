package languages

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/definitions"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter(
	patternMatchadapter definitions.PatternMatchAdapter,
) Adapter {
	builder := NewBuilder()
	valueBuilder := NewValueBuilder()
	return createAdapter(patternMatchadapter, builder, valueBuilder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewValueBuilder creates a new value builder
func NewValueBuilder() ValueBuilder {
	return createValueBuilder()
}

// Adapter represents an adapter
type Adapter interface {
	ToLanguage(parsed parsers.LanguageCommand) (Language, error)
}

// Builder represents a language builder
type Builder interface {
	Create() Builder
	WithVariable(variable string) Builder
	WithValues(values []Value) Builder
	Now() (Language, error)
}

// Language represents a language command
type Language interface {
	Variable() string
	Values() []Value
}

// ValueBuilder represents a language value builder
type ValueBuilder interface {
	Create() ValueBuilder
	WithRoot(root string) ValueBuilder
	WithTokensPath(tokensPath string) ValueBuilder
	WithRulesPath(rulesPath string) ValueBuilder
	WithLogicsPath(logicsPath string) ValueBuilder
	WithPatternMatches(patternMatches []definitions.PatternMatch) ValueBuilder
	WithInputVariable(inputVariable string) ValueBuilder
	WithChannelsPath(channelsPath string) ValueBuilder
	WithExtends(extends []string) ValueBuilder
	Now() (Value, error)
}

// Value represents a language value
type Value interface {
	IsRoot() bool
	Root() string
	IsTokensPath() bool
	TokensPath() string
	IsRulesPath() bool
	RulesPath() string
	IsLogicsPath() bool
	LogicsPath() string
	IsPatternMatches() bool
	PatternMatches() []definitions.PatternMatch
	IsInputVariable() bool
	InputVariable() string
	IsChannelsPath() bool
	ChannelsPath() string
	IsExtends() bool
	Extends() []string
}
