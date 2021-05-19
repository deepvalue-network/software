package languages

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/definitions"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	patternMatchAdapter := definitions.NewPatternMatchAdapter()
	builder := NewBuilder()
	valueBuilder := NewValueBuilder()
	return createAdapter(patternMatchAdapter, builder, valueBuilder)
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
	WithTokensPath(tokensPath parsers.RelativePath) ValueBuilder
	WithRulesPath(rulesPath parsers.RelativePath) ValueBuilder
	WithLogicsPath(logicsPath parsers.RelativePath) ValueBuilder
	WithPatternMatches(patternMatches []definitions.PatternMatch) ValueBuilder
	WithInputVariable(inputVariable string) ValueBuilder
	WithChannelsPath(channelsPath parsers.RelativePath) ValueBuilder
	WithExtends(extends []parsers.RelativePath) ValueBuilder
	Now() (Value, error)
}

// Value represents a language value
type Value interface {
	IsRoot() bool
	Root() string
	IsTokensPath() bool
	TokensPath() parsers.RelativePath
	IsRulesPath() bool
	RulesPath() parsers.RelativePath
	IsLogicsPath() bool
	LogicsPath() parsers.RelativePath
	IsPatternMatches() bool
	PatternMatches() []definitions.PatternMatch
	IsInputVariable() bool
	InputVariable() string
	IsChannelsPath() bool
	ChannelsPath() parsers.RelativePath
	IsExtends() bool
	Extends() []parsers.RelativePath
}
