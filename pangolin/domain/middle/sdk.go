package middle

import (
	"github.com/deepvalue-network/software/libs/hash"
	"github.com/deepvalue-network/software/pangolin/domain/middle/instructions"
	"github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/labels"
	"github.com/deepvalue-network/software/pangolin/domain/middle/tests"
	"github.com/deepvalue-network/software/pangolin/domain/middle/tests/test"
	test_instructions "github.com/deepvalue-network/software/pangolin/domain/middle/tests/test/instructions"
	test_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/tests/test/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/variables"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter(parser parsers.Parser) Adapter {
	instructionAdapter := instruction.NewAdapter()
	instructionsAdapter := instructions.NewAdapter(instructionAdapter)
	labelsAdapter := labels.NewAdapter()
	testInstructionAdapter := test_instruction.NewAdapter(instructionAdapter)
	testInstructionsAdapter := test_instructions.NewAdapter(testInstructionAdapter)
	testAdapter := test.NewAdapter(testInstructionsAdapter)
	testsAdapter := tests.NewAdapter(testAdapter)
	programBuilder := NewBuilder()
	applicationBuilder := NewApplicationBuilder()
	externalBuilder := NewExternalBuilder()
	languageBuilder := NewLanguageBuilder()
	patternMatchBuilder := NewPatternMatchBuilder()
	scriptBuilder := NewScriptBuilder()
	return createAdapter(
		parser,
		instructionsAdapter,
		labelsAdapter,
		testsAdapter,
		programBuilder,
		applicationBuilder,
		externalBuilder,
		languageBuilder,
		patternMatchBuilder,
		scriptBuilder,
	)
}

// NewBuilder creates a new program builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewApplicationBuilder creates a new application builder instance
func NewApplicationBuilder() ApplicationBuilder {
	instructionsBuilder := instructions.NewBuilder()
	labelsBuilder := labels.NewBuilder()
	variablesBuilder := variables.NewBuilder()
	testsBuilder := tests.NewBuilder()
	return createApplicationBuilder(instructionsBuilder, labelsBuilder, variablesBuilder, testsBuilder)
}

// NewExternalBuilder creates a new external builder instance
func NewExternalBuilder() ExternalBuilder {
	hashAdapter := hash.NewAdapter()
	return createExternalBuilder(hashAdapter)
}

// NewLanguageBuilder creates a new language builder
func NewLanguageBuilder() LanguageBuilder {
	return createLanguageBuilder()
}

// NewPatternMatchBuilder creates a new pattern match builder
func NewPatternMatchBuilder() PatternMatchBuilder {
	return createPatternMatchBuilder()
}

// NewScriptBuilder creates a new script builder
func NewScriptBuilder() ScriptBuilder {
	return createScriptBuilder()
}

// Adapter represents a program adapter
type Adapter interface {
	ToProgram(parsed parsers.Program) (Program, error)
}

// Builder represents a program builder
type Builder interface {
	Create() Builder
	WithApplication(app Application) Builder
	WithLanguage(lang Language) Builder
	WithScript(script Script) Builder
	Now() (Program, error)
}

// Program represents a program
type Program interface {
	IsApplication() bool
	Application() Application
	IsLanguage() bool
	Language() Language
	IsScript() bool
	Script() Script
}

// ScriptBuilder represents a script builder
type ScriptBuilder interface {
	Create() ScriptBuilder
	WithName(name string) ScriptBuilder
	WithVersion(version string) ScriptBuilder
	WithLanguagePath(lang string) ScriptBuilder
	WithScriptPath(script string) ScriptBuilder
	Now() (Script, error)
}

// Script represents a script
type Script interface {
	Name() string
	Version() string
	LanguagePath() string
	ScriptPath() string
}

// LanguageBuilder represents the language builder
type LanguageBuilder interface {
	Create() LanguageBuilder
	WithRoot(root string) LanguageBuilder
	WithTokensPath(tokens string) LanguageBuilder
	WithChannelsPath(channels string) LanguageBuilder
	WithRulesPath(rules string) LanguageBuilder
	WithLogicsPath(logics string) LanguageBuilder
	WithPatternMatches(patternMatches []PatternMatch) LanguageBuilder
	WithInputVariable(input string) LanguageBuilder
	WithOutputVariable(output string) LanguageBuilder
	WithExtends(extends []string) LanguageBuilder
	Now() (Language, error)
}

// Language represents the language
type Language interface {
	Root() string
	TokensPath() string
	RulesPath() string
	LogicsPath() string
	PatternMatches() []PatternMatch
	InputVariable() string
	OutputVariable() string
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

// ApplicationBuilder represents an application builder
type ApplicationBuilder interface {
	Create() ApplicationBuilder
	WithName(name string) ApplicationBuilder
	WithVersion(version string) ApplicationBuilder
	WithImports(imports []External) ApplicationBuilder
	WithExtends(extends []External) ApplicationBuilder
	WithInstructions(instructions instructions.Instructions) ApplicationBuilder
	WithTests(tests tests.Tests) ApplicationBuilder
	WithLabels(labels labels.Labels) ApplicationBuilder
	WithVariables(variables variables.Variables) ApplicationBuilder
	Now() (Application, error)
}

// Application represents an application
type Application interface {
	Name() string
	Version() string
	Instructions() instructions.Instructions
	Tests() tests.Tests
	Labels() labels.Labels
	Variables() variables.Variables
	HasImports() bool
	Imports() []External
	HasExtends() bool
	Extends() []External
}

// ExternalBuilder represents the external builder
type ExternalBuilder interface {
	Create() ExternalBuilder
	WithName(name string) ExternalBuilder
	WithPath(path string) ExternalBuilder
	Now() (External, error)
}

// External represents an external
type External interface {
	Name() string
	Path() string
}
