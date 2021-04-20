package commands

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/heads"
	language_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/commands/labels"
	test_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/tests/test/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/definitions"
)

// Builder represents a command builder
type Builder interface {
	Create() Builder
	WithLanguage(lang Language) Builder
	WithScript(script Script) Builder
	WithHead(head Head) Builder
	WithMain(main Main) Builder
	WithTest(test Test) Builder
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
	Head() Head
	IsMain() bool
	Main() Main
	IsTest() bool
	Test() Test
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

// HeadBuilder represents an head builder
type HeadBuilder interface {
	Create() HeadBuilder
	WithVariable(variable string) HeadBuilder
	WithValues(values []HeadValue) HeadBuilder
	Now() (Head, error)
}

// Head represents a head command
type Head interface {
	Variable() string
	Values() []HeadValue
}

// HeadValueBuilder represents an headValue builder
type HeadValueBuilder interface {
	Create() HeadValueBuilder
	WithName(name string) HeadValueBuilder
	WithVersion(version string) HeadValueBuilder
	WithImports(imports []string) HeadValueBuilder
	Now() (HeadValue, error)
}

// HeadValue represents an head value
type HeadValue interface {
	IsName() bool
	Name() string
	IsVersion() bool
	Version() string
	IsImports() bool
	Imports() []heads.External
}

// MainBuilder represents a main builder
type MainBuilder interface {
	Create() MainBuilder
	WithVariable(variable string) MainBuilder
	WithInstructions(ins []MainInstruction) MainBuilder
	Now() (Main, error)
}

// Main represents a main command
type Main interface {
	Variable() string
	Instructions() []MainInstruction
}

// MainInstructionBuilder represents a main instruction builder
type MainInstructionBuilder interface {
	Create() MainInstructionBuilder
	WithInstruction(ins language_instruction.Instruction) MainInstructionBuilder
	WithExternals(externals []bool) MainInstructionBuilder
	Now() (MainInstruction, error)
}

// MainInstruction represents a main instruction
type MainInstruction interface {
	Instruction() language_instruction.Instruction
	HasExternals() bool
	Externals() []bool
}

// TestBuilder represents a test builder
type TestBuilder interface {
	Create() TestBuilder
	WithName(name string) TestBuilder
	WithVariable(variable string) TestBuilder
	WithInstructions(ins []TestInstruction) TestBuilder
	Now() (Test, error)
}

// Test represents a test command
type Test interface {
	Name() string
	Variable() string
	Instructions() []TestInstruction
}

// TestInstructionBuilder represents a test instruction builder
type TestInstructionBuilder interface {
	Create() TestInstructionBuilder
	WithInstruction(ins test_instruction.Instruction) TestInstructionBuilder
	WithExternals(externals []bool) TestInstructionBuilder
	Now() (TestInstruction, error)
}

// TestInstruction represents a test instruction
type TestInstruction interface {
	Instruction() test_instruction.Instruction
	HasExternals() bool
	Externals() []bool
}
