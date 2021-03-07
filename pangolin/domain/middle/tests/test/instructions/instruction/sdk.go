package instruction

import (
	ins "github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter(
	instructionAdapter ins.Adapter,
) Adapter {
	builder := NewBuilder()
	assertBuilder := NewAssertBuilder()
	readFileBuilder := NewReadFileBuilder()
	return createAdapter(builder, assertBuilder, readFileBuilder, instructionAdapter)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewReadFileBuilder creates a new readFile builder
func NewReadFileBuilder() ReadFileBuilder {
	return createReadFileBuilder()
}

// NewAssertBuilder creates a new assert builder instance
func NewAssertBuilder() AssertBuilder {
	return createAssertBuilder()
}

// Adapter represents an instruction adapter
type Adapter interface {
	ToInstruction(testInstruction parsers.TestInstruction) (Instruction, error)
}

// Builder represents an instruction builder
type Builder interface {
	Create() Builder
	WithAssert(assert Assert) Builder
	WithReadFile(readFile ReadFile) Builder
	WithInstruction(ins ins.Instruction) Builder
	Now() (Instruction, error)
}

// Instruction represents a test instruction
type Instruction interface {
	IsAssert() bool
	Assert() Assert
	IsReadFile() bool
	ReadFile() ReadFile
	IsInstruction() bool
	Instruction() ins.Instruction
}

// ReadFileBuilder represents the readFile builder
type ReadFileBuilder interface {
	Create() ReadFileBuilder
	WithVariable(variable string) ReadFileBuilder
	WithPath(path string) ReadFileBuilder
	Now() (ReadFile, error)
}

// ReadFile represents a readFile test instruction
type ReadFile interface {
	Variable() string
	Path() string
}

// AssertBuilder represents an assert builder
type AssertBuilder interface {
	Create() AssertBuilder
	WithCondition(condition string) AssertBuilder
	Now() (Assert, error)
}

// Assert represents an assert
type Assert interface {
	HasCondition() bool
	Condition() string
}
