package middle

type patternMatch struct {
	pattern  string
	variable string
	enter    string
	exit     string
}

func createPatternMatchWithEnterAndExit(
	pattern string,
	variable string,
	enter string,
	exit string,
) PatternMatch {
	return createPatternMatchInternally(pattern, variable, enter, exit)
}

func createPatternMatchWithEnter(
	pattern string,
	variable string,
	enter string,
) PatternMatch {
	return createPatternMatchInternally(pattern, variable, enter, "")
}

func createPatternMatchWithExit(
	pattern string,
	variable string,
	exit string,
) PatternMatch {
	return createPatternMatchInternally(pattern, variable, "", exit)
}

func createPatternMatchInternally(
	pattern string,
	variable string,
	enter string,
	exit string,
) PatternMatch {
	out := patternMatch{
		pattern:  pattern,
		variable: variable,
		enter:    enter,
		exit:     exit,
	}

	return &out
}

// Pattern returns the pattern
func (obj *patternMatch) Pattern() string {
	return obj.pattern
}

// Variable returns the variable
func (obj *patternMatch) Variable() string {
	return obj.variable
}

// HasEnterLabel returns true if there is an enter label, false otherwise
func (obj *patternMatch) HasEnterLabel() bool {
	return obj.enter != ""
}

// EnterLabel returns the enter label, if any
func (obj *patternMatch) EnterLabel() string {
	return obj.enter
}

// HasExitLabel returns true if there is an exit label, false otherwise
func (obj *patternMatch) HasExitLabel() bool {
	return obj.exit != ""
}

// ExitLabel returns the exit label, if any
func (obj *patternMatch) ExitLabel() string {
	return obj.exit
}
