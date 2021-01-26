package match

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

// Input returns the input
func (obj *match) Input() string {
	return obj.input
}

// HasPattern returns true if there is a patern, false otherwise
func (obj *match) HasPattern() bool {
	return obj.pattern != ""
}

// Pattern returns the pattern
func (obj *match) Pattern() string {
	return obj.pattern
}
