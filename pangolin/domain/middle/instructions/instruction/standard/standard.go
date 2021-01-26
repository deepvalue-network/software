package standard

type standard struct {
	op     Operation
	res    string
	first  string
	second string
}

func createStandard(op Operation, res string, first string, second string) Standard {
	out := standard{
		op:     op,
		res:    res,
		first:  first,
		second: second,
	}

	return &out
}

// Operation returns the operation
func (obj *standard) Operation() Operation {
	return obj.op
}

// Result returns the result
func (obj *standard) Result() string {
	return obj.res
}

// First returns the first
func (obj *standard) First() string {
	return obj.first
}

// Second returns the second
func (obj *standard) Second() string {
	return obj.second
}
