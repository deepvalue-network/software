package languages

import "github.com/deepvalue-network/software/pangolin/domain/middle/languages/definitions"

type value struct {
	root           string
	tokensPath     string
	rulesPath      string
	logicsPath     string
	patternMatches []definitions.PatternMatch
	inputVariable  string
	channelsPath   string
	extends        []string
}

func createValueWithRoot(root string) Value {
	return createValueInternally(root, "", "", "", nil, "", "", nil)
}

func createValueWithTokensPath(tokensPath string) Value {
	return createValueInternally("", tokensPath, "", "", nil, "", "", nil)
}

func createValueWithRulesPath(rulesPath string) Value {
	return createValueInternally("", "", rulesPath, "", nil, "", "", nil)
}

func createValueWithLogicsPath(logicsPath string) Value {
	return createValueInternally("", "", "", logicsPath, nil, "", "", nil)
}

func createValueWithPatternMatches(patternMatches []definitions.PatternMatch) Value {
	return createValueInternally("", "", "", "", patternMatches, "", "", nil)
}

func createValueWithInputVariable(inputVariable string) Value {
	return createValueInternally("", "", "", "", nil, inputVariable, "", nil)
}

func createValueWithChannelsPath(channelsPath string) Value {
	return createValueInternally("", "", "", "", nil, "", channelsPath, nil)
}

func createValueWithExtends(extends []string) Value {
	return createValueInternally("", "", "", "", nil, "", "", extends)
}

func createValueInternally(
	root string,
	tokensPath string,
	rulesPath string,
	logicsPath string,
	patternMatches []definitions.PatternMatch,
	inputVariable string,
	channelsPath string,
	extends []string,
) Value {
	out := value{
		root:           root,
		tokensPath:     tokensPath,
		rulesPath:      rulesPath,
		logicsPath:     logicsPath,
		patternMatches: patternMatches,
		inputVariable:  inputVariable,
		channelsPath:   channelsPath,
		extends:        extends,
	}

	return &out
}

// IsRoot returns true if there is a root, false otherwise
func (obj *value) IsRoot() bool {
	return obj.root != ""
}

// Root returns the root, if any
func (obj *value) Root() string {
	return obj.root
}

// IsTokensPath returns true if there is a tokensPath, false otherwise
func (obj *value) IsTokensPath() bool {
	return obj.tokensPath != ""
}

// TokensPath returns the tokensPath, if any
func (obj *value) TokensPath() string {
	return obj.tokensPath
}

// IsRulesPath returns true if there is a rulesPath, false otherwise
func (obj *value) IsRulesPath() bool {
	return obj.rulesPath != ""
}

// RulesPath returns the rulesPath, if any
func (obj *value) RulesPath() string {
	return obj.rulesPath
}

// IsLogicsPath returns true if there is a logicsPath, false otherwise
func (obj *value) IsLogicsPath() bool {
	return obj.logicsPath != ""
}

// LogicsPath returns the logicsPath, if any
func (obj *value) LogicsPath() string {
	return obj.logicsPath
}

// IsPatternMatches returns true if there is pattern matches, false otherwise
func (obj *value) IsPatternMatches() bool {
	return obj.patternMatches != nil
}

// PatternMatches returns the pattern matches, if any
func (obj *value) PatternMatches() []definitions.PatternMatch {
	return obj.patternMatches
}

// IsInputVariable returns true if there is an input variable, false otherwise
func (obj *value) IsInputVariable() bool {
	return obj.inputVariable != ""
}

// InputVariable returns the input variable, if any
func (obj *value) InputVariable() string {
	return obj.inputVariable
}

// IsChannelsPath returns true if there is a channelsPath, false otherwise
func (obj *value) IsChannelsPath() bool {
	return obj.channelsPath != ""
}

// ChannelsPath returns the channelsPath, if any
func (obj *value) ChannelsPath() string {
	return obj.channelsPath
}

// IsExtends returns true if there is an extends, false otherwise
func (obj *value) IsExtends() bool {
	return obj.extends != nil
}

// Extends returns the extends, if any
func (obj *value) Extends() []string {
	return obj.extends
}
