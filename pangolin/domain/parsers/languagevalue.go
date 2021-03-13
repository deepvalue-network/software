package parsers

type languageValue struct {
	root           string
	tokens         RelativePath
	channels       RelativePath
	rules          RelativePath
	logic          RelativePath
	input          string
	output         string
	extends        []RelativePath
	patternMatches []PatternMatch
	targets        []Target
}

func createLanguageValueWithRoot(root string) LanguageValue {
	return createLanguageValueInternally(root, nil, nil, nil, nil, "", "", nil, nil, nil)
}

func createLanguageValueWithTokens(tokens RelativePath) LanguageValue {
	return createLanguageValueInternally("", tokens, nil, nil, nil, "", "", nil, nil, nil)
}

func createLanguageValueWithChannels(channels RelativePath) LanguageValue {
	return createLanguageValueInternally("", nil, channels, nil, nil, "", "", nil, nil, nil)
}

func createLanguageValueWithRules(rules RelativePath) LanguageValue {
	return createLanguageValueInternally("", nil, nil, rules, nil, "", "", nil, nil, nil)
}

func createLanguageValueWithLogic(logic RelativePath) LanguageValue {
	return createLanguageValueInternally("", nil, nil, nil, logic, "", "", nil, nil, nil)
}

func createLanguageValueWithInput(input string) LanguageValue {
	return createLanguageValueInternally("", nil, nil, nil, nil, input, "", nil, nil, nil)
}

func createLanguageValueWithOutput(output string) LanguageValue {
	return createLanguageValueInternally("", nil, nil, nil, nil, "", output, nil, nil, nil)
}

func createLanguageValueWithExtends(extends []RelativePath) LanguageValue {
	return createLanguageValueInternally("", nil, nil, nil, nil, "", "", extends, nil, nil)
}

func createLanguageValueWithPatternMatches(patternMatches []PatternMatch) LanguageValue {
	return createLanguageValueInternally("", nil, nil, nil, nil, "", "", nil, patternMatches, nil)
}

func createLanguageValueWithTargets(targets []Target) LanguageValue {
	return createLanguageValueInternally("", nil, nil, nil, nil, "", "", nil, nil, targets)
}

func createLanguageValueInternally(
	root string,
	tokens RelativePath,
	channels RelativePath,
	rules RelativePath,
	logic RelativePath,
	input string,
	output string,
	extends []RelativePath,
	patternMatches []PatternMatch,
	targets []Target,
) LanguageValue {
	out := languageValue{
		root:           root,
		tokens:         tokens,
		channels:       channels,
		rules:          rules,
		logic:          logic,
		input:          input,
		output:         output,
		extends:        extends,
		patternMatches: patternMatches,
		targets:        targets,
	}

	return &out
}

// IsRoot returns true if there is a root, false otherwise
func (obj *languageValue) IsRoot() bool {
	return obj.root != ""
}

// Root returns the root, if any
func (obj *languageValue) Root() string {
	return obj.root
}

// IsTokens returns true if there is a tokens, false otherwise
func (obj *languageValue) IsTokens() bool {
	return obj.tokens != nil
}

// Tokens returns the tokens, if any
func (obj *languageValue) Tokens() RelativePath {
	return obj.tokens
}

// IsChannels returns true if there is a channels, false otherwise
func (obj *languageValue) IsChannels() bool {
	return obj.channels != nil
}

// Channels returns the channels, if any
func (obj *languageValue) Channels() RelativePath {
	return obj.channels
}

// IsRules returns true if there is a rules, false otherwise
func (obj *languageValue) IsRules() bool {
	return obj.rules != nil
}

// Rules returns the rules, if any
func (obj *languageValue) Rules() RelativePath {
	return obj.rules
}

// IsLogic returns true if there is a logic, false otherwise
func (obj *languageValue) IsLogic() bool {
	return obj.logic != nil
}

// Logic returns the logic, if any
func (obj *languageValue) Logic() RelativePath {
	return obj.logic
}

// IsInputVariable returns true if there is an inputVariable, false otherwise
func (obj *languageValue) IsInputVariable() bool {
	return obj.input != ""
}

// IsInputVariable returns the input variable, if any
func (obj *languageValue) InputVariable() string {
	return obj.input
}

// IsOutputVariable returns true if there is an outputVariable, false otherwise
func (obj *languageValue) IsOutputVariable() bool {
	return obj.output != ""
}

// OutputVariable returns the output variable, if any
func (obj *languageValue) OutputVariable() string {
	return obj.output
}

// IsExtends returns true if there is extends, false otherwise
func (obj *languageValue) IsExtends() bool {
	return obj.extends != nil
}

// Extends returns the extends, if any
func (obj *languageValue) Extends() []RelativePath {
	return obj.extends
}

// IsPatternMatches returns true if there is pattern matches, false otherwise
func (obj *languageValue) IsPatternMatches() bool {
	return obj.patternMatches != nil
}

// PatternMatches returns the pattern matches, if any
func (obj *languageValue) PatternMatches() []PatternMatch {
	return obj.patternMatches
}

// IsTargets returns true if there is targets, false otherwise
func (obj *languageValue) IsTargets() bool {
	return obj.targets != nil
}

// Targets returns the targets, if any
func (obj *languageValue) Targets() []Target {
	return obj.targets
}
