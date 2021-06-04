package value

import "github.com/deepvalue-network/software/pangolin/domain/middle/testables/executables/applications/instructions/instruction/variable/value/computable"

type value struct {
	isStackFrame bool
	comp         computable.Value
	variable     string
}

func createValueWithStackFrame() Value {
	return createValueInternally(true, nil, "")
}

func createValueWithComputable(comp computable.Value) Value {
	return createValueInternally(false, comp, "")
}

func createValueWithVariable(variable string) Value {
	return createValueInternally(false, nil, variable)
}

func createValueInternally(
	isStackFrame bool,
	comp computable.Value,
	variable string,
) Value {
	out := value{
		isStackFrame: isStackFrame,
		comp:         comp,
		variable:     variable,
	}

	return &out
}

// IsStackFrame returns true if the value is a stackframe, false otherwise
func (obj *value) IsStackFrame() bool {
	return obj.isStackFrame
}

// IsComputable returns true if the value is computable, false otherwise
func (obj *value) IsComputable() bool {
	return obj.comp != nil
}

// Computable returns the computable value, if any
func (obj *value) Computable() computable.Value {
	return obj.comp
}

// IsVariable returns true if the value is variable, false otherwise
func (obj *value) IsVariable() bool {
	return obj.variable != ""
}

// Variable returns the variable, if any
func (obj *value) Variable() string {
	return obj.variable
}
