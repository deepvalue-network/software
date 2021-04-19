package parsers

type scopes struct {
	list []Scope
}

func createScopes(
	list []Scope,
) Scopes {
	out := scopes{
		list: list,
	}

	return &out
}

// All returns the scopes
func (obj *scopes) All() []Scope {
	return obj.list
}
