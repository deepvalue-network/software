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
	readFileBuilder := NewReadFileBuilder()
	return createAdapter(builder, readFileBuilder, instructionAdapter)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewReadFileBuilder creates a new readFile builder
func NewReadFileBuilder() ReadFileBuilder {
	return createReadFileBuilder()
}

// Adapter represents an instruction adapter
type Adapter interface {
	ToInstruction(testInstruction parsers.TestInstruction) (Instruction, error)
}

// Builder represents an instruction builder
type Builder interface {
	Create() Builder
	IsStart() Builder
	IsStop() Builder
	WithReadFile(readFile ReadFile) Builder
	WithInstruction(ins ins.Instruction) Builder
	Now() (Instruction, error)
}

// Instruction represents a test instruction
type Instruction interface {
	IsStart() bool
	IsStop() bool
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
