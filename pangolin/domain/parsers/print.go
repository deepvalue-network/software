package parsers

type print struct {
	value Value
}

func createPrint(value Value) Print {
	out := print{
		value: value,
	}

	return &out
}

// Value returns the value
func (obj *print) Value() Value {
	return obj.value
}
