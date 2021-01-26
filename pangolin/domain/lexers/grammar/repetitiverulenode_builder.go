package grammar

import "errors"

type repetitiveRuleNodeBuilder struct {
	recursiveName       string
	node                RuleNode
	isMultipleMandatory bool
	isMultipleOptional  bool
	isOptional          bool
}

func createRepetitiveRuleNodeBuilder() RepetitiveRuleNodeBuilder {
	out := repetitiveRuleNodeBuilder{
		recursiveName:       "",
		node:                nil,
		isMultipleMandatory: false,
		isMultipleOptional:  false,
		isOptional:          false,
	}

	return &out
}

// Create initialies the builder
func (app *repetitiveRuleNodeBuilder) Create() RepetitiveRuleNodeBuilder {
	return createRepetitiveRuleNodeBuilder()
}

// WithRecursiveName adds a recursiveName to the builder
func (app *repetitiveRuleNodeBuilder) WithRecursiveName(recursiveName string) RepetitiveRuleNodeBuilder {
	app.recursiveName = recursiveName
	return app
}

// WithNodes adds a node to the builder
func (app *repetitiveRuleNodeBuilder) WithNode(node RuleNode) RepetitiveRuleNodeBuilder {
	app.node = node
	return app
}

// IsMultipleMandatory flags the builder as multiple mandatory
func (app *repetitiveRuleNodeBuilder) IsMultipleMandatory() RepetitiveRuleNodeBuilder {
	app.isMultipleMandatory = true
	return app
}

// IsMultipleOptional flags the builder as multiple optional
func (app *repetitiveRuleNodeBuilder) IsMultipleOptional() RepetitiveRuleNodeBuilder {
	app.isMultipleOptional = true
	return app
}

// IsOptional flags the builder as optional
func (app *repetitiveRuleNodeBuilder) IsOptional() RepetitiveRuleNodeBuilder {
	app.isOptional = true
	return app
}

// Now builds a new RepetitiveRuleNode instance
func (app *repetitiveRuleNodeBuilder) Now() (RepetitiveRuleNode, error) {
	if app.recursiveName != "" {
		if app.isMultipleMandatory {
			return createRepetitiveRuleNodeRecursivelyWithMultipleMandatory(app.recursiveName), nil
		}

		if app.isMultipleOptional {
			return createRepetitiveRuleNodeRecursivelyWithMultipleOptional(app.recursiveName), nil
		}

		if app.isOptional {
			return createRepetitiveRuleNodeRecursivelyWithOptional(app.recursiveName), nil
		}

		return createRepetitiveRuleNodeRecursively(app.recursiveName), nil
	}

	if app.node != nil {
		if app.isMultipleMandatory {
			return createRepetitiveRuleNodeWithMultipleMandatory(app.node), nil
		}

		if app.isMultipleOptional {
			return createRepetitiveRuleNodeWithMultipleOptional(app.node), nil
		}

		if app.isOptional {
			return createRepetitiveRuleNodeWithOptional(app.node), nil
		}

		return createRepetitiveRuleNode(app.node), nil
	}

	return nil, errors.New("the RepetitiveRuleNode is invalid")
}
