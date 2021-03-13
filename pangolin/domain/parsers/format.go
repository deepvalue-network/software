package parsers

type format struct {
	results VariableName
	pattern Identifier
	first   Identifier
	second  Identifier
}

func createFormat(
	results VariableName,
	pattern Identifier,
	first Identifier,
	second Identifier,
) Format {
	out := format{
		results: results,
		pattern: pattern,
		first:   first,
		second:  second,
	}

	return &out
}

// Results returns the results variableName
func (obj *format) Results() VariableName {
	return obj.results
}

// Pattern returns the results identifier
func (obj *format) Pattern() Identifier {
	return obj.pattern
}

// First returns the first identifier
func (obj *format) First() Identifier {
	return obj.first
}

// Second returns the second identifier
func (obj *format) Second() Identifier {
	return obj.second
}
