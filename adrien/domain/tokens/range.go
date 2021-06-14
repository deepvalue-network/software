package tokens

type rnge struct {
	min uint
	max *uint
}

func createRange(
	min uint,
) Range {
	return createRangeInternally(min, nil)
}

func createRangeWithMaximum(
	min uint,
	max *uint,
) Range {
	return createRangeInternally(min, max)
}

func createRangeInternally(
	min uint,
	max *uint,
) Range {
	out := rnge{
		min: min,
		max: max,
	}

	return &out
}

// Min returns the minimum
func (obj *rnge) Min() uint {
	return obj.min
}

// HasMax returns true if there is a maximum
func (obj *rnge) HasMax() bool {
	return obj.max != nil
}

// Max returns the maximum
func (obj *rnge) Max() *uint {
	return obj.max
}
