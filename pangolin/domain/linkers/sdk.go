package linkers

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/labels"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/tests"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/definitions"
)

// NewProgramBuilder creates a new program builder instance
func NewProgramBuilder() ProgramBuilder {
	return createProgramBuilder()
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

// NewLanguageReferenceBuilder creates a new language reference builder
func NewLanguageReferenceBuilder() LanguageReferenceBuilder {
	return createLanguageReferenceBuilder()
}

// NewLanguageBuilder creates a new language builder
func NewLanguageBuilder() LanguageBuilder {
	return createLanguageBuilder()
}

// NewPathsBuilder creates a new paths builder
func NewPathsBuilder() PathsBuilder {
	return createPathsBuilder()
}

// ProgramBuilder represents a program builder
type ProgramBuilder interface {
	Create() ProgramBuilder
	WithApplication(app Application) ProgramBuilder
	WithLanguage(lang LanguageReference) ProgramBuilder
	WithScript(script Script) ProgramBuilder
	Now() (Program, error)
}

// Program represents a linked program
type Program interface {
	IsApplication() bool
	Application() Application
	IsLanguage() bool
	Language() LanguageReference
	IsScript() bool
	Script() Script
}

// ApplicationBuilder represents an application builder
type ApplicationBuilder interface {
	Create() ApplicationBuilder
	WithName(name string) ApplicationBuilder
	WithVersion(version string) ApplicationBuilder
	WithInstructions(ins instructions.Instructions) ApplicationBuilder
	WithTests(tests tests.Tests) ApplicationBuilder
	WithLabels(labels labels.Labels) ApplicationBuilder
	WithVariables(vars []variable.Variable) ApplicationBuilder
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
	Variables() []variable.Variable
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
	Now() (Script, error)
}

// Script represents a script
type Script interface {
	Language() LanguageReference
	Name() string
	Version() string
	Code() string
}

// LanguageReferenceBuilder represents a language reference builder
type LanguageReferenceBuilder interface {
	Create() LanguageReferenceBuilder
	WithLanguage(language Language) LanguageReferenceBuilder
	WithInputVariable(input string) LanguageReferenceBuilder
	WithOutputVariable(output string) LanguageReferenceBuilder
	Now() (LanguageReference, error)
}

// LanguageReference represents a language reference
type LanguageReference interface {
	Language() Language
	Input() string
	Output() string
}

// LanguageBuilder represents a language builder
type LanguageBuilder interface {
	Create() LanguageBuilder
	WithApplication(app Application) LanguageBuilder
	WithPatternMatches(matches []definitions.PatternMatch) LanguageBuilder
	WithPaths(paths Paths) LanguageBuilder
	WithRoot(root string) LanguageBuilder
	Now() (Language, error)
}

// Language represents a language application
type Language interface {
	Application() Application
	PatternMatches() []definitions.PatternMatch
	Paths() Paths
	Root() string
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
