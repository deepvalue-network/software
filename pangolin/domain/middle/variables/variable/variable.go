package variable

import (
	var_value "github.com/steve-care-software/products/pangolin/domain/middle/variables/variable/value"
)

type variable struct {
	isGlobal    bool
	isImmutable bool
	isMandatory bool
	isIncoming  bool
	isOutgoing  bool
	name        string
	value       var_value.Value
}

func createVariable(
	isGlobal bool,
	isImmutable bool,
	isMandatory bool,
	isIncoming bool,
	isOutgoing bool,
	name string,
	value var_value.Value,
) Variable {
	out := variable{
		isGlobal:    isGlobal,
		isImmutable: isImmutable,
		isMandatory: isMandatory,
		isIncoming:  isIncoming,
		isOutgoing:  isOutgoing,
		name:        name,
		value:       value,
	}

	return &out
}

// IsGlobal returns true if the variable is global, false otherwise
func (obj *variable) IsGlobal() bool {
	return obj.isGlobal
}

// IsImmutable returns true if the variable is immutable, false otherwise
func (obj *variable) IsImmutable() bool {
	return obj.isImmutable
}

// IsMandatory returns true if the variable is mandatory, false otherwise
func (obj *variable) IsMandatory() bool {
	return obj.isMandatory
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
