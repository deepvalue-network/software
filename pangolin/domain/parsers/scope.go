package parsers

type scope struct {
	isInternal bool
	isExternal bool
}

func createScopeWithInternal() Scope {
	return createScope(true, false)
}

func createScopeWithExternal() Scope {
	return createScope(false, true)
}

func createScope(
	isInternal bool,
	isExternal bool,
) Scope {
	out := scope{
		isInternal: isInternal,
		isExternal: isExternal,
	}

	return &out
}

// IsInternal returns true if the scope is internal, false otherwise
func (obj *scope) IsInternal() bool {
	return obj.isInternal
}

// IsExternal returns true if the scope is external, false otherwise
func (obj *scope) IsExternal() bool {
	return obj.isExternal
}
