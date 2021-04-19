package parsers

type print struct {
	value ValueRepresentation
}

func createPrint(value ValueRepresentation) Print {
	out := print{
		value: value,
	}

	return &out
}

// Value returns the value
func (obj *print) Value() ValueRepresentation {
	return obj.value
}
