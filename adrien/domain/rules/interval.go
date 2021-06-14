package rules

type interval struct {
	min int
	max int
}

func createInterval(
	min int,
) Interval {
	return createIntervalInternally(min, -1)
}

func createIntervalWithMax(
	min int,
	max int,
) Interval {
	return createIntervalInternally(min, max)
}

func createIntervalInternally(
	min int,
	max int,
) Interval {
	out := interval{
		min: min,
		max: max,
	}

	return &out
}

// Min returns the mininum
func (obj *interval) Min() int {
	return obj.min
}

// HasMax returns true if there is a maximum, false otherwise
func (obj *interval) HasMax() bool {
	return obj.max != -1
}

// Max returns the maximum, if any
func (obj *interval) Max() int {
	return obj.max
}
