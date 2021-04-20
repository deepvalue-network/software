package commands

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/commands/heads"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/commands/labels"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/commands/mains"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/commands/tests"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/definitions"
)

// Builder represents a command builder
type Builder interface {
	Create() Builder
	WithLanguage(lang Language) Builder
	WithScript(script Script) Builder
	WithHead(head heads.Head) Builder
	WithMain(main mains.Main) Builder
	WithTest(test tests.Test) Builder
	WithLabel(label labels.Label) Builder
	Now() (Command, error)
}

// Command represents a command
type Command interface {
	IsLanguage() bool
	Language() Language
	IsScript() bool
	Script() Script
	IsHead() bool
	Head() heads.Head
	IsMain() bool
	Main() mains.Main
	IsTest() bool
	Test() tests.Test
	IsLabel() bool
	Label() labels.Label
}

// LanguageBuilder represents a language builder
type LanguageBuilder interface {
	Create() LanguageBuilder
	WithVariable(variable string) LanguageBuilder
	WithValues(values []LanguageValue) LanguageBuilder
	Now() (Language, error)
}

// Language represents a language command
type Language interface {
	Variable() string
	Values() []LanguageValue
}

// LanguageValueBuilder represents a language value builder
type LanguageValueBuilder interface {
	Create() LanguageValueBuilder
	WithRoot(root string) LanguageValueBuilder
	WithTokensPath(tokensPath string) LanguageValueBuilder
	WithRulesPath(rulesPath string) LanguageValueBuilder
	WithLogicsPath(logicsPath string) LanguageValueBuilder
	WithPatternMatches(patternMatches []definitions.PatternMatch) LanguageValueBuilder
	WithInputVariable(inputVariable string) LanguageValueBuilder
	WithChannelsPath(channelsPath string) LanguageValueBuilder
	WithExtends(extends []string) LanguageValueBuilder
	Now() (LanguageValue, error)
}

// LanguageValue represents a language value
type LanguageValue interface {
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

// ScriptBuilder represents a script builder
type ScriptBuilder interface {
	Create() ScriptBuilder
	WithVariable(variable string) ScriptBuilder
	WithValues(values []ScriptValue) ScriptBuilder
	Now() (Script, error)
}

// Script represents a script command
type Script interface {
	Variable() string
	Values() []ScriptValue
}

// ScriptValueBuilder represents a script value builder
type ScriptValueBuilder interface {
	Create() ScriptValueBuilder
	WithName(name string) ScriptValueBuilder
	WithVersion(version string) ScriptValueBuilder
	WithLanguagePath(langPath string) ScriptValueBuilder
	WithScriptPath(scriptPath string) ScriptValueBuilder
	Now() (ScriptValue, error)
}

// ScriptValue represents a script value
type ScriptValue interface {
	IsName() bool
	Name() string
	IsVersion() bool
	Version() string
	IsLanguagePath() bool
	LanguagePath() string
	IsScriptPath() bool
	ScriptPath() string
}
