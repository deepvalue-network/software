package registry

type fetch struct {
	from  string
	to    string
	index Index
}

func createFetch(
	from string,
	to string,
) Fetch {
	return createFetchInternally(from, to, nil)
}

func createFetchWithIndex(
	from string,
	to string,
	index Index,
) Fetch {
	return createFetchInternally(from, to, index)
}

func createFetchInternally(
	from string,
	to string,
	index Index,
) Fetch {
	out := fetch{
		from:  from,
		to:    to,
		index: index,
	}

	return &out
}

// From returns the from variable
func (obj *fetch) From() string {
	return obj.from
}

// To returns the to variable
func (obj *fetch) To() string {
	return obj.to
}

// HasIndex returns true if there is an index, false otherwise
func (obj *fetch) HasIndex() bool {
	return obj.index != nil
}

// Index returns the index, if any
func (obj *fetch) Index() Index {
	return obj.index
}
