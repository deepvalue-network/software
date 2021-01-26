package grammar

import (
	"errors"
	"sort"
)

type ruleBuilder struct {
	name     string
	sections []RuleSection
}

func createRuleBuilder() RuleBuilder {
	out := ruleBuilder{
		name:     "",
		sections: nil,
	}

	return &out
}

// Create initializes the builder
func (app *ruleBuilder) Create() RuleBuilder {
	return createRuleBuilder()
}

// WithName adds a name to tjhe builder
func (app *ruleBuilder) WithName(name string) RuleBuilder {
	app.name = name
	return app
}

// WithSections adds sections to tjhe builder
func (app *ruleBuilder) WithSections(sections []RuleSection) RuleBuilder {
	app.sections = sections
	return app
}

// Now builds a new Rule instance
func (app *ruleBuilder) Now() (Rule, error) {
	if app.name == "" {
		return nil, errors.New("the name is mandatory in order to build a Rule instance")
	}

	if app.sections == nil {
		return nil, errors.New("the name is mandatory in order to build a Rule instance")
	}

	keys := []int{}
	sections := map[int]RuleSection{}
	for _, oneSection := range app.sections {
		index := -1
		if oneSection.HasConstant() {
			index = oneSection.Constant().Index()
		}

		if oneSection.HasPattern() {
			index = oneSection.Pattern().Index()
		}

		keys = append(keys, index)
		sections[index] = oneSection
	}

	sort.Ints(keys)
	orderedSections := []RuleSection{}
	for _, index := range keys {
		orderedSections = append(orderedSections, sections[index])
	}

	return createRule(app.name, orderedSections, keys[0]), nil
}
