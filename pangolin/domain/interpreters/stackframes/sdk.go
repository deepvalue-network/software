package stackframes

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/remaining"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/standard"
	var_variable "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable"
	"github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value/computable"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	frameBuilder := NewFrameBuilder()
	return createBuilder(frameBuilder)
}

// NewFrameBuilder creates a new frame builder
func NewFrameBuilder() FrameBuilder {
	computableBuilder := computable.NewBuilder()
	computer := createComputer(computableBuilder)
	return createFrameBuilder(computer, computableBuilder)
}

// Builder represents a stackframe builder
type Builder interface {
	Create() Builder
	WithVariables(variables map[string]computable.Value) Builder
	Now() StackFrame
}

// StackFrame represents a stackframe
type StackFrame interface {
	Push()
	Pop() error
	Index() int
	Add(frame Frame)
	Skip(index int) error
	Current() Frame
	Registry() Registry
}

// FrameBuilder represents a frame builder
type FrameBuilder interface {
	Create() FrameBuilder
	WithVariables(variables map[string]computable.Value) FrameBuilder
	Now() Frame
}

// Frame represents a frame
type Frame interface {
	Standard(first string, second string, result string, operation standard.Operation) error
	Remaining(first string, second string, result string, remaining string, operation remaining.Operation) error
	Insert(operation var_variable.Variable) error
	Update(operation var_variable.Variable) error
	UpdateValue(name string, val computable.Value) error
	Delete(name string) error
	Fetch(name string) (computable.Value, error)
	Stop()
	IsStopped() bool
}

// Registry represents a registry
type Registry interface {
	All() map[string]computable.Value
	Fetch(name string) (computable.Value, error)
	Insert(name string, val computable.Value) error
	Delete(name string) error
}

// Computer represents a computer
type Computer interface {
	Add(first computable.Value, second computable.Value) (computable.Value, error)
	Substract(first computable.Value, second computable.Value) (computable.Value, error)
	Multiply(first computable.Value, second computable.Value) (computable.Value, error)
	Divide(first computable.Value, second computable.Value) (computable.Value, computable.Value, error)
	IsLessThan(first computable.Value, second computable.Value) (computable.Value, error)
	IsEqual(first computable.Value, second computable.Value) (computable.Value, error)
	IsNotEqual(first computable.Value, second computable.Value) (computable.Value, error)
	And(first computable.Value, second computable.Value) (computable.Value, error)
	Or(first computable.Value, second computable.Value) (computable.Value, error)
	Concat(first computable.Value, second computable.Value) (computable.Value, error)
	Match(pattern computable.Value, value computable.Value) ([]computable.Value, computable.Value, error)
}
