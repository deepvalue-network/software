package linkers

import (
	"github.com/deepvalue-network/software/pangolin/domain/lexers/grammar"
	"github.com/deepvalue-network/software/pangolin/domain/middle"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/tests"
	language_instructions "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions"
	language_labels "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/labels"
	language_tests "github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/tests"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/definitions"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// ParseFileFn parses a path and returns the parsed program
type ParseFileFn func(filePath string) (parsers.Program, error)

const scriptName = "default"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	middleAdapter := middle.NewAdapter()
	grammarRetrieverCriteriaBuilder := grammar.NewRetrieverCriteriaBuilder()
	applicationBuilder := NewApplicationBuilder()
	languageBuilder := NewLanguageBuilder()
	executableBuilder := NewExecutableBuilder()
	programBuilder := NewProgramBuilder()
	languageDefinitionBuilder := NewLanguageDefinitionBuilder()
	pathsBuilder := NewPathsBuilder()
	scriptBuilder := NewScriptBuilder()
	testBuilder := NewTestBuilder()
	languageReferenceBuilder := NewLanguageReferenceBuilder()
	languageApplicationBuilder := NewLanguageApplicationBuilder()
	return createBuilder(
		middleAdapter,
		grammarRetrieverCriteriaBuilder,
		applicationBuilder,
		languageBuilder,
		executableBuilder,
		programBuilder,
		languageDefinitionBuilder,
		pathsBuilder,
		scriptBuilder,
		testBuilder,
		languageReferenceBuilder,
		languageApplicationBuilder,
	)
}

// NewExecutableBuilder creates a new executable builder
func NewExecutableBuilder() ExecutableBuilder {
	return createExecutableBuilder()
}

// NewProgramBuilder creates a new program builder instance
func NewProgramBuilder() ProgramBuilder {
	return createProgramBuilder()
}

// NewLanguageBuilder creates a new language builder instance
func NewLanguageBuilder() LanguageBuilder {
	return createLanguageBuilder()
}

// NewApplicationBuilder creates a new application builder
func NewApplicationBuilder() ApplicationBuilder {
	return createApplicationBuilder()
}

// NewExternalBuilder creates a new external builder instance
func NewExternalBuilder() ExternalBuilder {
	return createExternalBuilder()
}

// NewScriptBuilder creates a new script builder instance
func NewScriptBuilder() ScriptBuilder {
	return createScriptBuilder()
}

// NewTestBuilder creates a new test buildeer instance
func NewTestBuilder() TestBuilder {
	return createTestBuilder()
}

// NewLanguageReferenceBuilder creates a new language reference builder
func NewLanguageReferenceBuilder() LanguageReferenceBuilder {
	return createLanguageReferenceBuilder()
}

// NewLanguageDefinitionBuilder creates a new language builder
func NewLanguageDefinitionBuilder() LanguageDefinitionBuilder {
	return createLanguageDefinitionBuilder()
}

// NewLanguageApplicationBuilder creates a new language application builder
func NewLanguageApplicationBuilder() LanguageApplicationBuilder {
	return createLanguageApplicationBuilder()
}

// NewPathsBuilder creates a new paths builder
func NewPathsBuilder() PathsBuilder {
	return createPathsBuilder()
}

// Builder represents a linker builder
type Builder interface {
	Create() Builder
	WithParser(parser parsers.Parser) Builder
	WithDirPath(dirPath string) Builder
	Now() (Linker, error)
}

// Linker represents a linker application
type Linker interface {
	Execute(parsed parsers.Program) (Executable, error)
}

// ExecutableBuilder represents an executable builder
type ExecutableBuilder interface {
	Create() ExecutableBuilder
	WithApplication(app Application) ExecutableBuilder
	WithScript(script Script) ExecutableBuilder
	Now() (Executable, error)
}

// Executable represents an executable
type Executable interface {
	IsApplication() bool
	Application() Application
	IsScript() bool
	Script() Script
}

// ProgramBuilder represents a program builder
type ProgramBuilder interface {
	Create() ProgramBuilder
	WithApplication(app Application) ProgramBuilder
	WithLanguage(lang Language) ProgramBuilder
	WithScript(script Script) ProgramBuilder
	Now() (Program, error)
}

// Program represents a linked program
type Program interface {
	IsApplication() bool
	Application() Application
	IsLanguage() bool
	Language() Language
	IsScript() bool
	Script() Script
}

// LanguageBuilder represents a language builder
type LanguageBuilder interface {
	Create() LanguageBuilder
	WithReference(ref LanguageReference) LanguageBuilder
	WithApplication(app LanguageApplication) LanguageBuilder
	Now() (Language, error)
}

// Language represents a language
type Language interface {
	IsReference() bool
	Reference() LanguageReference
	IsApplication() bool
	Application() LanguageApplication
}

// ApplicationBuilder represents an application builder
type ApplicationBuilder interface {
	Create() ApplicationBuilder
	WithName(name string) ApplicationBuilder
	WithVersion(version string) ApplicationBuilder
	WithInstructions(ins instructions.Instructions) ApplicationBuilder
	WithTests(tests tests.Tests) ApplicationBuilder
	WithLabels(labels labels.Labels) ApplicationBuilder
	WithImports(imps []External) ApplicationBuilder
	Now() (Application, error)
}

// Application represents a linked application
type Application interface {
	Name() string
	Version() string
	Instructions() instructions.Instructions
	Tests() tests.Tests
	Labels() labels.Labels
	HasImports() bool
	Imports() []External
	Import(name string) (Application, error)
}

// ExternalBuilder represents an external builder
type ExternalBuilder interface {
	Create() ExternalBuilder
	WithName(name string) ExternalBuilder
	WithApplication(application Application) ExternalBuilder
	WithScript(script Script) ExternalBuilder
	Now() (External, error)
}

// External represents an imported external application
type External interface {
	Name() string
	HasApplication() bool
	Application() Application
	HasScript() bool
	Script() Script
}

// ScriptBuilder represents a script builder
type ScriptBuilder interface {
	Create() ScriptBuilder
	WithLanguage(language LanguageReference) ScriptBuilder
	WithName(name string) ScriptBuilder
	WithVersion(version string) ScriptBuilder
	WithCode(code string) ScriptBuilder
	WithOutput(output string) ScriptBuilder
	WithTests(tests []Test) ScriptBuilder
	Now() (Script, error)
}

// Script represents a script
type Script interface {
	Language() LanguageReference
	Name() string
	Version() string
	Code() string
	Output() string
	HasTests() bool
	Tests() []Test
}

// TestBuilder represents a test builder
type TestBuilder interface {
	Create() TestBuilder
	WithName(name string) TestBuilder
	WithScript(script Script) TestBuilder
	Now() (Test, error)
}

// Test represents a script test
type Test interface {
	Name() string
	Script() Script
}

// LanguageReferenceBuilder represents a language reference builder
type LanguageReferenceBuilder interface {
	Create() LanguageReferenceBuilder
	WithDefinition(def LanguageDefinition) LanguageReferenceBuilder
	WithInputVariable(input string) LanguageReferenceBuilder
	Now() (LanguageReference, error)
}

// LanguageReference represents a language reference
type LanguageReference interface {
	Definition() LanguageDefinition
	Input() string
}

// LanguageDefinitionBuilder represents a language builder
type LanguageDefinitionBuilder interface {
	Create() LanguageDefinitionBuilder
	WithApplication(app LanguageApplication) LanguageDefinitionBuilder
	WithPatternMatches(matches []definitions.PatternMatch) LanguageDefinitionBuilder
	WithPaths(paths Paths) LanguageDefinitionBuilder
	WithRoot(root string) LanguageDefinitionBuilder
	Now() (LanguageDefinition, error)
}

// LanguageDefinition represents a language definition
type LanguageDefinition interface {
	Application() LanguageApplication
	PatternMatches() []definitions.PatternMatch
	Paths() Paths
	Root() string
}

// LanguageApplicationBuilder represents a language application builder
type LanguageApplicationBuilder interface {
	Create() LanguageApplicationBuilder
	WithName(name string) LanguageApplicationBuilder
	WithVersion(version string) LanguageApplicationBuilder
	WithInstructions(ins language_instructions.Instructions) LanguageApplicationBuilder
	WithTests(tests language_tests.Tests) LanguageApplicationBuilder
	WithLabels(labels language_labels.Labels) LanguageApplicationBuilder
	WithImports(imps []External) LanguageApplicationBuilder
	Now() (LanguageApplication, error)
}

// LanguageApplication represents a linked language application
type LanguageApplication interface {
	Name() string
	Version() string
	Instructions() language_instructions.Instructions
	Tests() language_tests.Tests
	Labels() language_labels.Labels
	HasImports() bool
	Imports() []External
	Import(name string) (Application, error)
}

// PathsBuilder represents a paths builder
type PathsBuilder interface {
	Create() PathsBuilder
	WithBaseDir(baseDir string) PathsBuilder
	WithTokens(tokens string) PathsBuilder
	WithRules(rules string) PathsBuilder
	WithLogics(logics string) PathsBuilder
	WithChannels(channels string) PathsBuilder
	Now() (Paths, error)
}

// Paths represents a paths instance
type Paths interface {
	BaseDir() string
	Tokens() string
	Rules() string
	Logics() string
	HasChannels() bool
	Channels() string
}
