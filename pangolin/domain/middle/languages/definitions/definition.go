package definitions

type definition struct {
	root           string
	tokens         string
	rules          string
	logics         string
	patternMatches []PatternMatch
	input          string
	channels       string
	extends        []string
}

func createDefinition(
	root string,
	tokens string,
	rules string,
	logics string,
	patternMatches []PatternMatch,
	input string,
) Definition {
	return createDefinitionInternally(root, tokens, rules, logics, patternMatches, input, "", nil)
}

func createDefinitionWithExtends(
	root string,
	tokens string,
	rules string,
	logics string,
	patternMatches []PatternMatch,
	input string,
	extends []string,
) Definition {
	return createDefinitionInternally(root, tokens, rules, logics, patternMatches, input, "", extends)
}

func createDefinitionWithChannels(
	root string,
	tokens string,
	rules string,
	logics string,
	patternMatches []PatternMatch,
	input string,
	channels string,
) Definition {
	return createDefinitionInternally(root, tokens, rules, logics, patternMatches, input, channels, nil)
}

func createDefinitionWithChannelsAndExtends(
	root string,
	tokens string,
	rules string,
	logics string,
	patternMatches []PatternMatch,
	input string,
	channels string,
	extends []string,
) Definition {
	return createDefinitionInternally(root, tokens, rules, logics, patternMatches, input, channels, extends)
}

func createDefinitionInternally(
	root string,
	tokens string,
	rules string,
	logics string,
	patternMatches []PatternMatch,
	input string,
	channels string,
	extends []string,
) Definition {
	out := definition{
		root:           root,
		tokens:         tokens,
		rules:          rules,
		logics:         logics,
		patternMatches: patternMatches,
		input:          input,
		channels:       channels,
		extends:        extends,
	}

	return &out
}

// Root returns the root
func (obj *definition) Root() string {
	return obj.root
}

// TokensPath returns the tokens path
func (obj *definition) TokensPath() string {
	return obj.tokens
}

// RulesPath returns the rules path
func (obj *definition) RulesPath() string {
	return obj.rules
}

// LogicsPath returns the logics path
func (obj *definition) LogicsPath() string {
	return obj.logics
}

// PatternMatches returns the pattern matches
func (obj *definition) PatternMatches() []PatternMatch {
	return obj.patternMatches
}

// InputVariable returns the input variable
func (obj *definition) InputVariable() string {
	return obj.input
}

// HasChannelsPath returns true if there is a channels path, false otherwise
func (obj *definition) HasChannelsPath() bool {
	return obj.channels != ""
}

// ChannelsPath returns the channels path, if any
func (obj *definition) ChannelsPath() string {
	return obj.channels
}

// HasExtends returns true if there is extends, false otherwise
func (obj *definition) HasExtends() bool {
	return obj.extends != nil
}

// Extends returns the extends, if any
func (obj *definition) Extends() []string {
	return obj.extends
}
