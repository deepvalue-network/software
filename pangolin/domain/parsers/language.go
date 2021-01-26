package parsers

type language struct {
	root           string
	patternMatches []PatternMatch
	tokens         RelativePath
	rules          RelativePath
	logic          RelativePath
	input          string
	output         string
	channels       RelativePath
	extends        []RelativePath
}

func createLanguage(
	root string,
	patternMatches []PatternMatch,
	tokens RelativePath,
	rules RelativePath,
	logic RelativePath,
	input string,
	output string,
) Language {
	return createLanguageinternally(root, patternMatches, tokens, rules, logic, input, output, nil, nil)
}

func createLanguageWithExtends(
	root string,
	patternMatches []PatternMatch,
	tokens RelativePath,
	rules RelativePath,
	logic RelativePath,
	input string,
	output string,
	extends []RelativePath,
) Language {
	return createLanguageinternally(root, patternMatches, tokens, rules, logic, input, output, nil, extends)
}

func createLanguageWithChannels(
	root string,
	patternMatches []PatternMatch,
	tokens RelativePath,
	rules RelativePath,
	logic RelativePath,
	input string,
	output string,
	channels RelativePath,
) Language {
	return createLanguageinternally(root, patternMatches, tokens, rules, logic, input, output, channels, nil)
}

func createLanguageWithChannelsAndExtends(
	root string,
	patternMatches []PatternMatch,
	tokens RelativePath,
	rules RelativePath,
	logic RelativePath,
	input string,
	output string,
	channels RelativePath,
	extends []RelativePath,
) Language {
	return createLanguageinternally(root, patternMatches, tokens, rules, logic, input, output, channels, extends)
}

func createLanguageinternally(
	root string,
	patternMatches []PatternMatch,
	tokens RelativePath,
	rules RelativePath,
	logic RelativePath,
	input string,
	output string,
	channels RelativePath,
	extends []RelativePath,
) Language {
	out := language{
		root:           root,
		patternMatches: patternMatches,
		tokens:         tokens,
		rules:          rules,
		logic:          logic,
		input:          input,
		output:         output,
		channels:       channels,
		extends:        extends,
	}

	return &out
}

// Root returns the root
func (obj *language) Root() string {
	return obj.root
}

// Tokens returns the tokens
func (obj *language) Tokens() RelativePath {
	return obj.tokens
}

// Rules returns the rules
func (obj *language) Rules() RelativePath {
	return obj.rules
}

// Logic returns the logic
func (obj *language) Logic() RelativePath {
	return obj.logic
}

// Input returns the input variable
func (obj *language) Input() string {
	return obj.input
}

// Output returns the output variable
func (obj *language) Output() string {
	return obj.output
}

// HasChannels returns true if there is channels, false otherwise
func (obj *language) HasChannels() bool {
	return obj.channels != nil
}

// Channels returns the channels, if any
func (obj *language) Channels() RelativePath {
	return obj.channels
}

// HasExtends returns true if there is extends, false otherwise
func (obj *language) HasExtends() bool {
	return obj.extends != nil
}

// Extends returns the extends, if any
func (obj *language) Extends() []RelativePath {
	return obj.extends
}

// HasPatternMatches returns true if there is pattern matches, false otherwise
func (obj *language) HasPatternMatches() bool {
	return obj.patternMatches != nil
}

// PatternMatches returns the pattern matches, if any
func (obj *language) PatternMatches() []PatternMatch {
	return obj.patternMatches
}
