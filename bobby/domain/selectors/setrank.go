package selectors

type setRank struct {
	from *uint
	to   *uint
}

func createSetRankWithFrom(
	from *uint,
) SetRank {
	return createSetRankInternally(from, nil)
}

func createSetRankWithTo(
	to *uint,
) SetRank {
	return createSetRankInternally(nil, to)
}

func createSetRankWithFromAndTo(
	from *uint,
	to *uint,
) SetRank {
	return createSetRankInternally(from, to)
}

func createSetRankInternally(
	from *uint,
	to *uint,
) SetRank {
	out := setRank{
		from: from,
		to:   to,
	}

	return &out
}

// HasFrom retruns true if there is a from, false otherwise
func (obj *setRank) HasFrom() bool {
	return obj.from != nil
}

// From returns the from, if any
func (obj *setRank) From() *uint {
	return obj.from
}

// HasTo retruns true if there is a to, false otherwise
func (obj *setRank) HasTo() bool {
	return obj.to != nil
}

// To returns the to, if any
func (obj *setRank) To() *uint {
	return obj.to
}
