package parsers

type languageDefinition struct {
	root           string
	patternMatches []PatternMatch
	tokens         RelativePath
	rules          RelativePath
	logic          RelativePath
	input          string
	channels       RelativePath
	extends        []RelativePath
}

func createLanguageDefinition(
	root string,
	patternMatches []PatternMatch,
	tokens RelativePath,
	rules RelativePath,
	logic RelativePath,
	input string,
) LanguageDefinition {
	return createLanguageDefinitioninternally(root, patternMatches, tokens, rules, logic, input, nil, nil)
}

func createLanguageDefinitionWithExtends(
	root string,
	patternMatches []PatternMatch,
	tokens RelativePath,
	rules RelativePath,
	logic RelativePath,
	input string,
	extends []RelativePath,
) LanguageDefinition {
	return createLanguageDefinitioninternally(root, patternMatches, tokens, rules, logic, input, nil, extends)
}

func createLanguageDefinitionWithChannels(
	root string,
	patternMatches []PatternMatch,
	tokens RelativePath,
	rules RelativePath,
	logic RelativePath,
	input string,
	channels RelativePath,
) LanguageDefinition {
	return createLanguageDefinitioninternally(root, patternMatches, tokens, rules, logic, input, channels, nil)
}

func createLanguageDefinitionWithChannelsAndExtends(
	root string,
	patternMatches []PatternMatch,
	tokens RelativePath,
	rules RelativePath,
	logic RelativePath,
	input string,
	channels RelativePath,
	extends []RelativePath,
) LanguageDefinition {
	return createLanguageDefinitioninternally(root, patternMatches, tokens, rules, logic, input, channels, extends)
}

func createLanguageDefinitioninternally(
	root string,
	patternMatches []PatternMatch,
	tokens RelativePath,
	rules RelativePath,
	logic RelativePath,
	input string,
	channels RelativePath,
	extends []RelativePath,
) LanguageDefinition {
	out := languageDefinition{
		root:           root,
		patternMatches: patternMatches,
		tokens:         tokens,
		rules:          rules,
		logic:          logic,
		input:          input,
		channels:       channels,
		extends:        extends,
	}

	return &out
}

// Root returns the root
func (obj *languageDefinition) Root() string {
	return obj.root
}

// Tokens returns the tokens
func (obj *languageDefinition) Tokens() RelativePath {
	return obj.tokens
}

// Rules returns the rules
func (obj *languageDefinition) Rules() RelativePath {
	return obj.rules
}

// Logic returns the logic
func (obj *languageDefinition) Logic() RelativePath {
	return obj.logic
}

// Input returns the input variable
func (obj *languageDefinition) Input() string {
	return obj.input
}

// HasChannels returns true if there is channels, false otherwise
func (obj *languageDefinition) HasChannels() bool {
	return obj.channels != nil
}

// Channels returns the channels, if any
func (obj *languageDefinition) Channels() RelativePath {
	return obj.channels
}

// HasExtends returns true if there is extends, false otherwise
func (obj *languageDefinition) HasExtends() bool {
	return obj.extends != nil
}

// Extends returns the extends, if any
func (obj *languageDefinition) Extends() []RelativePath {
	return obj.extends
}

// HasPatternMatches returns true if there is pattern matches, false otherwise
func (obj *languageDefinition) HasPatternMatches() bool {
	return obj.patternMatches != nil
}

// PatternMatches returns the pattern matches, if any
func (obj *languageDefinition) PatternMatches() []PatternMatch {
	return obj.patternMatches
}
