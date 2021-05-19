package variable

import (
	var_value "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value"
)

type variable struct {
	name  string
	value var_value.Value
}

func createVariable(
	name string,
	value var_value.Value,
) Variable {
	out := variable{
		name:  name,
		value: value,
	}

	return &out
}

// Name returns the name
func (obj *variable) Name() string {
	return obj.name
}

// Value returns the value
func (obj *variable) Value() var_value.Value {
	return obj.value
}
