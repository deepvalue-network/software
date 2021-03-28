package parsers

type match struct {
	input   string
	pattern string
}

func createMatch(
	input string,
) Match {
	return createMatchInternally(input, "")
}

func createMatchWithPattern(
	input string,
	pattern string,
) Match {
	return createMatchInternally(input, pattern)
}

func createMatchInternally(
	input string,
	pattern string,
) Match {
	out := match{
		input:   input,
		pattern: pattern,
	}

	return &out
}

// Input returns the input identifier
func (obj *match) Input() string {
	return obj.input
}

// HasPattern returns true if there is a pattern, false otherwise
func (obj *match) HasPattern() bool {
	return obj.pattern != ""
}

// Pattern returns the pattern string
func (obj *match) Pattern() string {
	return obj.pattern
}
