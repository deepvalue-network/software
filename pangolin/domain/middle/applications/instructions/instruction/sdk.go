package instruction

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/call"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/condition"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/exit"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/remaining"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/stackframe"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/standard"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/transform"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/value"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variablename"
	var_variable "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable"
	var_value "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value"
	"github.com/deepvalue-network/software/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	stackframeBuilder := stackframe.NewBuilder()
	skipBuilder := stackframe.NewSkipBuilder()
	transformBuilder := transform.NewBuilder()
	variableNameBuilder := variablename.NewBuilder()
	conditionBuilder := condition.NewBuilder()
	propositionBuilder := condition.NewPropositionBuilder()
	remainingBuilder := remaining.NewBuilder()
	standardBuilder := standard.NewBuilder()
	valueBuilder := value.NewBuilder()
	varValueAdapter := var_value.NewAdapter()
	varValueFactory := var_value.NewFactory()
	varVariableBuilder := var_variable.NewBuilder()
	callBuilder := call.NewBuilder()
	exitBuilder := exit.NewBuilder()
	builder := NewBuilder()
	return createAdapter(
		stackframeBuilder,
		skipBuilder,
		transformBuilder,
		variableNameBuilder,
		conditionBuilder,
		propositionBuilder,
		remainingBuilder,
		standardBuilder,
		valueBuilder,
		varValueAdapter,
		varValueFactory,
		varVariableBuilder,
		callBuilder,
		exitBuilder,
		builder,
	)
}

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// Adapter represents the instruction adapter
type Adapter interface {
	ToInstruction(instruction parsers.Instruction) (Instruction, error)
}

// Builder represents an instruction builder
type Builder interface {
	Create() Builder
	WithStackframe(stackframe stackframe.Stackframe) Builder
	WithTransform(transform transform.Transform) Builder
	WithVariableName(variableName variablename.VariableName) Builder
	WithCondition(condition condition.Condition) Builder
	WithStandard(standard standard.Standard) Builder
	WithRemaining(remaining remaining.Remaining) Builder
	WithValue(value value.Value) Builder
	WithInsert(insert var_variable.Variable) Builder
	WithSave(save var_variable.Variable) Builder
	WithDelete(del string) Builder
	WithCall(call call.Call) Builder
	WithExit(exit exit.Exit) Builder
	Now() (Instruction, error)
}

// Instruction represents an instruction
type Instruction interface {
	IsStackframe() bool
	Stackframe() stackframe.Stackframe
	IsTransform() bool
	Transform() transform.Transform
	IsVariableName() bool
	VariableName() variablename.VariableName
	IsCondition() bool
	Condition() condition.Condition
	IsStandard() bool
	Standard() standard.Standard
	IsRemaining() bool
	Remaining() remaining.Remaining
	IsValue() bool
	Value() value.Value
	IsInsert() bool
	Insert() var_variable.Variable
	IsSave() bool
	Save() var_variable.Variable
	IsDelete() bool
	Delete() string
	IsCall() bool
	Call() call.Call
	IsExit() bool
	Exit() exit.Exit
}
