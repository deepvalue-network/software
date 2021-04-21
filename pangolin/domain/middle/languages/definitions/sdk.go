package definitions

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Builder represents a language definition builder
type Builder interface {
	Create() Builder
	WithRoot(root string) Builder
	WithTokensPath(tokens string) Builder
	WithChannelsPath(channels string) Builder
	WithRulesPath(rules string) Builder
	WithLogicsPath(logics string) Builder
	WithPatternMatches(patternMatches []PatternMatch) Builder
	WithInputVariable(input string) Builder
	WithExtends(extends []string) Builder
	Now() (Definition, error)
}

// Definition represents a language definition
type Definition interface {
	Root() string
	TokensPath() string
	RulesPath() string
	LogicsPath() string
	PatternMatches() []PatternMatch
	InputVariable() string
	HasChannelsPath() bool
	ChannelsPath() string
	HasExtends() bool
	Extends() []string
}

// PatternMatchBuilder represents a patternMatch builder
type PatternMatchBuilder interface {
	Create() PatternMatchBuilder
	WithPattern(pattern string) PatternMatchBuilder
	WithEnterLabel(enter string) PatternMatchBuilder
	WithExitLabel(exit string) PatternMatchBuilder
	Now() (PatternMatch, error)
}

// PatternMatch represents a pattern match
type PatternMatch interface {
	Pattern() string
	HasEnterLabel() bool
	EnterLabel() string
	HasExitLabel() bool
	ExitLabel() string
}
