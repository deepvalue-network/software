package instruction

import (
	standard_instruction "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/commands"
	"github.com/deepvalue-network/software/pangolin/domain/middle/languages/applications/instructions/instruction/match"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	instructionAdapter := standard_instruction.NewAdapter()
	commandAdapter := commands.NewAdapter()
	matchAdapter := match.NewAdapter()
	commonInstructionBuilder := NewCommonInstructionBuilder()
	builder := NewBuilder()
	return createAdapter(
		instructionAdapter,
		commandAdapter,
		matchAdapter,
		commonInstructionBuilder,
		builder,
	)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewCommonInstructionBuilder creates a new common instruction builder
func NewCommonInstructionBuilder() CommonInstructionBuilder {
	return createCommonInstructionBuilder()
}

// Adapter represents instructions adapter
type Adapter interface {
	ToInstruction(parsed parsers.LanguageInstruction) (Instruction, error)
	ToCommonInstruction(parsed parsers.LanguageInstructionCommon) (CommonInstruction, error)
}

// Builder represents an instruction builder
type Builder interface {
	Create() Builder
	WithInstruction(ins CommonInstruction) Builder
	WithCommand(command commands.Command) Builder
	Now() (Instruction, error)
}

// Instruction represents a language application instruction
type Instruction interface {
	IsInstruction() bool
	Instruction() CommonInstruction
	IsCommand() bool
	Command() commands.Command
}

// CommonInstructionBuilder represents a common instruction builder
type CommonInstructionBuilder interface {
	Create() CommonInstructionBuilder
	WithInstruction(ins standard_instruction.Instruction) CommonInstructionBuilder
	WithMatch(match match.Match) CommonInstructionBuilder
	Now() (CommonInstruction, error)
}

// CommonInstruction represents a common instruction
type CommonInstruction interface {
	IsInstruction() bool
	Instruction() standard_instruction.Instruction
	IsMatch() bool
	Match() match.Match
}
