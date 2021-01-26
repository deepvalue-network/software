package parsers

type patternMatch struct {
	pattern  string
	labels   PatternLabels
	variable string
}

func createPatternMatch(
	pattern string,
	labels PatternLabels,
	variable string,
) PatternMatch {
	out := patternMatch{
		pattern:  pattern,
		labels:   labels,
		variable: variable,
	}

	return &out
}

// Pattern returns the pattern
func (obj *patternMatch) Pattern() string {
	return obj.pattern
}

// Labels returns the pattern labels
func (obj *patternMatch) Labels() PatternLabels {
	return obj.labels
}

// Variable returns the variable name
func (obj *patternMatch) Variable() string {
	return obj.variable
}
