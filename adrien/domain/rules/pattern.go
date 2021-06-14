package rules

import "regexp"

type pattern struct {
	name        string
	code        string
	ptrn        *regexp.Regexp
	possibility Possibility
}

func createPattern(
	name string,
	code string,
	ptrn *regexp.Regexp,
	possibility Possibility,
) Pattern {
	out := pattern{
		name:        name,
		code:        code,
		ptrn:        ptrn,
		possibility: possibility,
	}

	return &out
}

// Name returns the name
func (obj *pattern) Name() string {
	return obj.name
}

// Code returns the code
func (obj *pattern) Code() string {
	return obj.code
}

// Pattern returns the pattern
func (obj *pattern) Pattern() *regexp.Regexp {
	return obj.ptrn
}

// Possibility returns the possibility
func (obj *pattern) Possibility() Possibility {
	return obj.possibility
}
