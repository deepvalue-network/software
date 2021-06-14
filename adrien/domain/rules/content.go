package rules

type content struct {
	constant string
	pattern  Pattern
}

func createContentWithConstant(
	constant string,
) Content {
	return createContentInternally(constant, nil)
}

func createContentWithPattern(
	pattern Pattern,
) Content {
	return createContentInternally("", pattern)
}

func createContentInternally(
	constant string,
	pattern Pattern,
) Content {
	out := content{
		constant: constant,
		pattern:  pattern,
	}

	return &out
}

// IsConstant returns true if there is a constant, false otherwise
func (obj *content) IsConstant() bool {
	return obj.constant != ""
}

// IConstant returns the constant, if any
func (obj *content) Constant() string {
	return obj.constant
}

// IsPattern returns true if there is a pattern, false otherwise
func (obj *content) IsPattern() bool {
	return obj.pattern != nil
}

// Pattern returns the pattern, if any
func (obj *content) Pattern() Pattern {
	return obj.pattern
}
