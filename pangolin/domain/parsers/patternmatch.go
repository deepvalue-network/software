package parsers

type patternMatch struct {
	pattern string
	labels  PatternLabels
}

func createPatternMatch(
	pattern string,
	labels PatternLabels,
) PatternMatch {
	out := patternMatch{
		pattern: pattern,
		labels:  labels,
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
