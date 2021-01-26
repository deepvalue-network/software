package instruction

import (
	"github.com/steve-care-software/products/pangolin/domain/middle/instructions/instruction/condition"
	"github.com/steve-care-software/products/pangolin/domain/middle/instructions/instruction/match"
	"github.com/steve-care-software/products/pangolin/domain/middle/instructions/instruction/remaining"
	"github.com/steve-care-software/products/pangolin/domain/middle/instructions/instruction/stackframe"
	"github.com/steve-care-software/products/pangolin/domain/middle/instructions/instruction/standard"
	"github.com/steve-care-software/products/pangolin/domain/middle/instructions/instruction/token"
	"github.com/steve-care-software/products/pangolin/domain/middle/instructions/instruction/transform"
	"github.com/steve-care-software/products/pangolin/domain/middle/instructions/instruction/value"
	"github.com/steve-care-software/products/pangolin/domain/middle/instructions/instruction/variablename"
	var_variable "github.com/steve-care-software/products/pangolin/domain/middle/variables/variable"
	var_value "github.com/steve-care-software/products/pangolin/domain/middle/variables/variable/value"
	"github.com/steve-care-software/products/pangolin/domain/parsers"
)

// NewAdapter creates a new adapter instance
func NewAdapter() Adapter {
	stackframeBuilder := stackframe.NewBuilder()
	transformBuilder := transform.NewBuilder()
	variableNameBuilder := variablename.NewBuilder()
	conditionBuilder := condition.NewBuilder()
	propositionBuilder := condition.NewPropositionBuilder()
	remainingBuilder := remaining.NewBuilder()
	standardBuilder := standard.NewBuilder()
	matchBuilder := match.NewBuilder()
	valueBuilder := value.NewBuilder()
	varValueAdapter := var_value.NewAdapter()
	varValueFactory := var_value.NewFactory()
	varVariableBuilder := var_variable.NewBuilder()
	tokenCodeBuilder := token.NewCodeBuilder()
	tokenCodeMatchBuilder := token.NewCodeMatchBuilder()
	tokenBuilder := token.NewBuilder()
	builder := NewBuilder()
	return createAdapter(
		stackframeBuilder,
		transformBuilder,
		variableNameBuilder,
		conditionBuilder,
		propositionBuilder,
		remainingBuilder,
		standardBuilder,
		matchBuilder,
		valueBuilder,
		varValueAdapter,
		varValueFactory,
		varVariableBuilder,
		tokenCodeBuilder,
		tokenCodeMatchBuilder,
		tokenBuilder,
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
	WithMatch(match match.Match) Builder
	WithToken(token token.Token) Builder
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
	IsMatch() bool
	Match() match.Match
	IsToken() bool
	Token() token.Token
}
