package grammar

import "errors"

type ruleNodeBuilder struct {
	rule Rule
	leaf RuleTree
}

func createRuleNodeBuilder() RuleNodeBuilder {
	out := ruleNodeBuilder{
		rule: nil,
		leaf: nil,
	}

	return &out
}

// Create initializes the builder
func (app *ruleNodeBuilder) Create() RuleNodeBuilder {
	return createRuleNodeBuilder()
}

// WithRule add a Rule to the builder
func (app *ruleNodeBuilder) WithRule(rule Rule) RuleNodeBuilder {
	app.rule = rule
	return app
}

// WithLeaf adds a leaf RuleTree to the builder
func (app *ruleNodeBuilder) WithLeaf(leaf RuleTree) RuleNodeBuilder {
	app.leaf = leaf
	return app
}

// Now builds a new RuleNode instance
func (app *ruleNodeBuilder) Now() (RuleNode, error) {
	if app.rule != nil {
		return createRuleNodeWithRule(app.rule), nil
	}

	if app.leaf != nil {
		return createRuleNodeWithLeaf(app.leaf), nil
	}

	return nil, errors.New("the RuleNode is invalid")
}
