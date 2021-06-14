package rules

type amount struct {
	exactly  int
	interval Interval
}

func createAmountWithExactly(
	exactly int,
) Amount {
	return createAmountInternally(exactly, nil)
}

func createAmountWithInterval(
	interval Interval,
) Amount {
	return createAmountInternally(-1, interval)
}

func createAmountInternally(
	exactly int,
	interval Interval,
) Amount {
	out := amount{
		exactly:  exactly,
		interval: interval,
	}

	return &out
}

// IsExactly returns true if there is an exact amount, false otherwise
func (obj *amount) IsExactly() bool {
	return obj.exactly != -1
}

// Exactly returns the exact amount, if any
func (obj *amount) Exactly() int {
	return obj.exactly
}

// IsInterval returns true if there is an interval, false otherwise
func (obj *amount) IsInterval() bool {
	return obj.interval != nil
}

// Interval returns the interval, if any
func (obj *amount) Interval() Interval {
	return obj.interval
}
