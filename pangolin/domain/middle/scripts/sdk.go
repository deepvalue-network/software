package scripts

import (
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	builder := NewBuilder()
	testsBuilder := NewTestsBuilder()
	testBuilder := NewTestBuilder()
	return createAdapter(builder, testsBuilder, testBuilder)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewTestsBuilder creates a new tests builder
func NewTestsBuilder() TestsBuilder {
	return createTestsBuilder()
}

// NewTestBuilder creates a new test builder
func NewTestBuilder() TestBuilder {
	return createTestBuilder()
}

// Adapter represents the script adapter
type Adapter interface {
	ToScript(parsed parsers.Script) (Script, error)
}

// Builder represents a script builder
type Builder interface {
	Create() Builder
	WithName(name string) Builder
	WithVersion(version string) Builder
	WithLanguagePath(lang string) Builder
	WithScriptPath(script string) Builder
	WithOutput(output string) Builder
	WithTests(tests Tests) Builder
	Now() (Script, error)
}

// Script represents a script
type Script interface {
	Name() string
	Version() string
	LanguagePath() string
	ScriptPath() string
	Output() string
	HasTests() bool
	Tests() Tests
}

// TestsBuilder represets a tests builder
type TestsBuilder interface {
	Create() TestsBuilder
	WithTests(tests []Test) TestsBuilder
	Now() (Tests, error)
}

// Tests represents tests
type Tests interface {
	All() []Test
	FetchByName(name string) (Test, error)
}

// TestBuilder represents a test builder
type TestBuilder interface {
	Create() TestBuilder
	WithName(name string) TestBuilder
	WithPath(path string) TestBuilder
	Now() (Test, error)
}

// Test represents a script test
type Test interface {
	Name() string
	Path() string
}
