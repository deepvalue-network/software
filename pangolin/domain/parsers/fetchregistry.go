package parsers

type fetchRegistry struct {
	to    string
	from  string
	index IntPointer
}

func createFetchRegister(
	to string,
	from string,
) FetchRegistry {
	return createFetchRegistryInternally(to, from, nil)
}

func createFetchRegisterWithIndex(
	to string,
	from string,
	index IntPointer,
) FetchRegistry {
	return createFetchRegistryInternally(to, from, index)
}

func createFetchRegistryInternally(
	to string,
	from string,
	index IntPointer,
) FetchRegistry {
	out := fetchRegistry{
		to:    to,
		from:  from,
		index: index,
	}

	return &out
}

// To returns the to variable
func (obj *fetchRegistry) To() string {
	return obj.to
}

// From returns the from variable
func (obj *fetchRegistry) From() string {
	return obj.from
}

// HasIndex returns true if there is an index, false otherwise
func (obj *fetchRegistry) HasIndex() bool {
	return obj.index != nil
}

// Index returns the index pointer, if any
func (obj *fetchRegistry) Index() IntPointer {
	return obj.index
}
