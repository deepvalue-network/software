package parsers

type declaration struct {
	name string
	typ  Type
}

func createDeclaration(name string, typ Type) Declaration {
	out := declaration{
		name: name,
		typ:  typ,
	}

	return &out
}

// Variable returns the variable name
func (obj *declaration) Variable() string {
	return obj.name
}

// Type returns the type
func (obj *declaration) Type() Type {
	return obj.typ
}
