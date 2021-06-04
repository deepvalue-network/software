package stackframe

type save struct {
	to   string
	from string
}

func createSave(
	to string,
) Save {
	return createSaveInternally(to, "")
}

func createSaveWithFrom(
	to string,
	from string,
) Save {
	return createSaveInternally(to, from)
}

func createSaveInternally(
	to string,
	from string,
) Save {
	out := save{
		from: from,
		to:   to,
	}

	return &out
}

// To returns the to variable
func (obj *save) To() string {
	return obj.to
}

// HasFrom returns true if there is a from variable, false otherwise
func (obj *save) HasFrom() bool {
	return obj.from != ""
}

// From returns the from variable, if any
func (obj *save) From() string {
	return obj.from
}
