package parsers

type constantDeclaration struct {
	constant string
	typ      Type
	value    Value
}

func createConstantDeclaration(constant string, typ Type, value Value) ConstantDeclaration {
	out := constantDeclaration{
		constant: constant,
		typ:      typ,
		value:    value,
	}

	return &out
}

// Constant returns the constant
func (obj *constantDeclaration) Constant() string {
	return obj.constant
}

// Type returns the type
func (obj *constantDeclaration) Type() Type {
	return obj.typ
}

// Value returns the value
func (obj *constantDeclaration) Value() Value {
	return obj.value
}
