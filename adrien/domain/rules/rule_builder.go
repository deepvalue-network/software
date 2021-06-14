package rules

import "errors"

type ruleBuilder struct {
	name     string
	code     string
	elements []Element
}

func createRuleBuilder() RuleBuilder {
	out := ruleBuilder{
		name:     "",
		code:     "",
		elements: nil,
	}

	return &out
}

// Create initializes the builder
func (app *ruleBuilder) Create() RuleBuilder {
	return createRuleBuilder()
}

// WithName adds a name to the builder
func (app *ruleBuilder) WithName(name string) RuleBuilder {
	app.name = name
	return app
}

// WithCode adds a code to the builder
func (app *ruleBuilder) WithCode(code string) RuleBuilder {
	app.code = code
	return app
}

// WithElements add elements to the builder
func (app *ruleBuilder) WithElements(elements []Element) RuleBuilder {
	app.elements = elements
	return app
}

// Now builds a new Rule instance
func (app *ruleBuilder) Now() (Rule, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Rule instance")
	}

	if app.code == "" {
		return nil, errors.New("the code is mandatory in order to build a Rule instance")
	}

	if app.elements != nil && len(app.elements) <= 0 {
		app.elements = nil
	}

	if app.elements == nil {
		return nil, errors.New("there must be at least 1 Element in order to build a Rule instance")
	}

	return createRule(app.name, app.code, app.elements), nil
}
