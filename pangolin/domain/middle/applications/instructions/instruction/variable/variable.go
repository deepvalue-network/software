package variable

import (
	var_value "github.com/deepvalue-network/software/pangolin/domain/middle/applications/instructions/instruction/variable/value"
)

type variable struct {
	isIncoming bool
	isOutgoing bool
	name       string
	value      var_value.Value
}

func createVariable(
	isIncoming bool,
	isOutgoing bool,
	name string,
	value var_value.Value,
) Variable {
	out := variable{
		isIncoming: isIncoming,
		isOutgoing: isOutgoing,
		name:       name,
		value:      value,
	}

	return &out
}

// IsIncoming returns true if the variable is incoming, false otherwise
func (obj *variable) IsIncoming() bool {
	return obj.isIncoming
}

// IsOutgoing returns true if the variable is outgoing, false otherwise
func (obj *variable) IsOutgoing() bool {
	return obj.isOutgoing
}

// Name returns the name
func (obj *variable) Name() string {
	return obj.name
}

// Value returns the value
func (obj *variable) Value() var_value.Value {
	return obj.value
}
