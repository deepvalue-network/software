package middle

import "github.com/deepvalue-network/software/pangolin/domain/middle/targets"

type language struct {
	root           string
	tokens         string
	rules          string
	logics         string
	patternMatches []PatternMatch
	input          string
	output         string
	targets        targets.Targets
	channels       string
	extends        []string
}

func createLanguage(
	root string,
	tokens string,
	rules string,
	logics string,
	patternMatches []PatternMatch,
	input string,
	output string,
	targets targets.Targets,
) Language {
	return createLanguageInternally(root, tokens, rules, logics, patternMatches, input, output, targets, "", nil)
}

func createLanguageWithExtends(
	root string,
	tokens string,
	rules string,
	logics string,
	patternMatches []PatternMatch,
	input string,
	output string,
	targets targets.Targets,
	extends []string,
) Language {
	return createLanguageInternally(root, tokens, rules, logics, patternMatches, input, output, targets, "", extends)
}

func createLanguageWithChannels(
	root string,
	tokens string,
	rules string,
	logics string,
	patternMatches []PatternMatch,
	input string,
	output string,
	targets targets.Targets,
	channels string,
) Language {
	return createLanguageInternally(root, tokens, rules, logics, patternMatches, input, output, targets, channels, nil)
}

func createLanguageWithChannelsAndExtends(
	root string,
	tokens string,
	rules string,
	logics string,
	patternMatches []PatternMatch,
	input string,
	output string,
	targets targets.Targets,
	channels string,
	extends []string,
) Language {
	return createLanguageInternally(root, tokens, rules, logics, patternMatches, input, output, targets, channels, extends)
}

func createLanguageInternally(
	root string,
	tokens string,
	rules string,
	logics string,
	patternMatches []PatternMatch,
	input string,
	output string,
	targets targets.Targets,
	channels string,
	extends []string,
) Language {
	out := language{
		root:           root,
		tokens:         tokens,
		rules:          rules,
		logics:         logics,
		patternMatches: patternMatches,
		input:          input,
		output:         output,
		targets:        targets,
		channels:       channels,
		extends:        extends,
	}

	return &out
}

// Root returns the root
func (obj *language) Root() string {
	return obj.root
}

// TokensPath returns the tokens path
func (obj *language) TokensPath() string {
	return obj.tokens
}

// RulesPath returns the rules path
func (obj *language) RulesPath() string {
	return obj.rules
}

// LogicsPath returns the logics path
func (obj *language) LogicsPath() string {
	return obj.logics
}

// PatternMatches returns the pattern matches
func (obj *language) PatternMatches() []PatternMatch {
	return obj.patternMatches
}

// InputVariable returns the input variable
func (obj *language) InputVariable() string {
	return obj.input
}

// OutputVariable returns the output variable
func (obj *language) OutputVariable() string {
	return obj.output
}

// Targets returns the targets
func (obj *language) Targets() targets.Targets {
	return obj.targets
}

// HasChannelsPath returns true if there is a channels path, false otherwise
func (obj *language) HasChannelsPath() bool {
	return obj.channels != ""
}

// ChannelsPath returns the channels path, if any
func (obj *language) ChannelsPath() string {
	return obj.channels
}

// HasExtends returns true if there is extends, false otherwise
func (obj *language) HasExtends() bool {
	return obj.extends != nil
}

// Extends returns the extends, if any
func (obj *language) Extends() []string {
	return obj.extends
}
