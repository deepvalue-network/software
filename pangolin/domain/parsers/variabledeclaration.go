package parsers

type variableDeclaration struct {
	typ       Type
	variable  string
	direction VariableDirection
}

func createVariableDeclaration(typ Type, variable string) VariableDeclaration {
	return createVariableDeclarationInternally(typ, variable, nil)
}

func createVariableDeclarationWithDirection(typ Type, variable string, direction VariableDirection) VariableDeclaration {
	return createVariableDeclarationInternally(typ, variable, direction)
}

func createVariableDeclarationInternally(
	typ Type,
	variable string,
	direction VariableDirection,
) VariableDeclaration {
	out := variableDeclaration{
		typ:       typ,
		variable:  variable,
		direction: direction,
	}

	return &out
}

// Type returns the type
func (obj *variableDeclaration) Type() Type {
	return obj.typ
}

// Variable returns the variable
func (obj *variableDeclaration) Variable() string {
	return obj.variable
}

// HasDirection returns true if there is a direction, false otherwise
func (obj *variableDeclaration) HasDirection() bool {
	return obj.direction != nil
}

// Direction returns the direction, if any
func (obj *variableDeclaration) Direction() VariableDirection {
	return obj.direction
}
