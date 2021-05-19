package languages

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/definitions"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

type value struct {
	root           string
	tokensPath     parsers.RelativePath
	rulesPath      parsers.RelativePath
	logicsPath     parsers.RelativePath
	patternMatches []definitions.PatternMatch
	inputVariable  string
	channelsPath   parsers.RelativePath
	extends        []parsers.RelativePath
}

func createValueWithRoot(root string) Value {
	return createValueInternally(root, nil, nil, nil, nil, "", nil, nil)
}

func createValueWithTokensPath(tokensPath parsers.RelativePath) Value {
	return createValueInternally("", tokensPath, nil, nil, nil, "", nil, nil)
}

func createValueWithRulesPath(rulesPath parsers.RelativePath) Value {
	return createValueInternally("", nil, rulesPath, nil, nil, "", nil, nil)
}

func createValueWithLogicsPath(logicsPath parsers.RelativePath) Value {
	return createValueInternally("", nil, nil, logicsPath, nil, "", nil, nil)
}

func createValueWithPatternMatches(patternMatches []definitions.PatternMatch) Value {
	return createValueInternally("", nil, nil, nil, patternMatches, "", nil, nil)
}

func createValueWithInputVariable(inputVariable string) Value {
	return createValueInternally("", nil, nil, nil, nil, inputVariable, nil, nil)
}

func createValueWithChannelsPath(channelsPath parsers.RelativePath) Value {
	return createValueInternally("", nil, nil, nil, nil, "", channelsPath, nil)
}

func createValueWithExtends(extends []parsers.RelativePath) Value {
	return createValueInternally("", nil, nil, nil, nil, "", nil, extends)
}

func createValueInternally(
	root string,
	tokensPath parsers.RelativePath,
	rulesPath parsers.RelativePath,
	logicsPath parsers.RelativePath,
	patternMatches []definitions.PatternMatch,
	inputVariable string,
	channelsPath parsers.RelativePath,
	extends []parsers.RelativePath,
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
	return obj.tokensPath != nil
}

// TokensPath returns the tokensPath, if any
func (obj *value) TokensPath() parsers.RelativePath {
	return obj.tokensPath
}

// IsRulesPath returns true if there is a rulesPath, false otherwise
func (obj *value) IsRulesPath() bool {
	return obj.rulesPath != nil
}

// RulesPath returns the rulesPath, if any
func (obj *value) RulesPath() parsers.RelativePath {
	return obj.rulesPath
}

// IsLogicsPath returns true if there is a logicsPath, false otherwise
func (obj *value) IsLogicsPath() bool {
	return obj.logicsPath != nil
}

// LogicsPath returns the logicsPath, if any
func (obj *value) LogicsPath() parsers.RelativePath {
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
	return obj.channelsPath != nil
}

// ChannelsPath returns the channelsPath, if any
func (obj *value) ChannelsPath() parsers.RelativePath {
	return obj.channelsPath
}

// IsExtends returns true if there is an extends, false otherwise
func (obj *value) IsExtends() bool {
	return obj.extends != nil
}

// Extends returns the extends, if any
func (obj *value) Extends() []parsers.RelativePath {
	return obj.extends
}
