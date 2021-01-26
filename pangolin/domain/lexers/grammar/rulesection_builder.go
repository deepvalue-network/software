package grammar

import "errors"

type ruleSectionBuilder struct {
	constant RawToken
	pattern  RawToken
}

func createRuleSectionBuilder() RuleSectionBuilder {
	out := ruleSectionBuilder{
		constant: nil,
		pattern:  nil,
	}

	return &out
}

// Create initializes the builder
func (app *ruleSectionBuilder) Create() RuleSectionBuilder {
	return createRuleSectionBuilder()
}

// WithConstant adds a constant to the builder
func (app *ruleSectionBuilder) WithConstant(constant RawToken) RuleSectionBuilder {
	app.constant = constant
	return app
}

// WithPattern adds a pattern to the builder
func (app *ruleSectionBuilder) WithPattern(pattern RawToken) RuleSectionBuilder {
	app.pattern = pattern
	return app
}

// Now builds a new RuleSection instance
func (app *ruleSectionBuilder) Now() (RuleSection, error) {
	if app.constant != nil {
		return createRuleSectionWithConstant(app.constant), nil
	}

	if app.pattern != nil {
		return createRuleSectionWithPattern(app.pattern), nil
	}

	return nil, errors.New("the RuleSection is invalid")
}
