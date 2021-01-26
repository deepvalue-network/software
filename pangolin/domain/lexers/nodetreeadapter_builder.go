package lexers

import (
	"errors"

	"github.com/steve-care-software/products/pangolin/domain/lexers/grammar"
)

type nodeTreeAdapterBuilder struct {
	ruleTreeAdapter grammar.RuleTreeAdapter
	nodeTreeBuilder NodeTreeBuilder
	grammar         grammar.Grammar
}

func createNodeTreeAdapterBuilder(
	ruleTreeAdapter grammar.RuleTreeAdapter,
	nodeTreeBuilder NodeTreeBuilder,
) NodeTreeAdapterBuilder {
	out := nodeTreeAdapterBuilder{
		ruleTreeAdapter: ruleTreeAdapter,
		nodeTreeBuilder: nodeTreeBuilder,
		grammar:         nil,
	}

	return &out
}

// Create initializes the builder
func (app *nodeTreeAdapterBuilder) Create() NodeTreeAdapterBuilder {
	return createNodeTreeAdapterBuilder(app.ruleTreeAdapter, app.nodeTreeBuilder)
}

// WithGrammar adds a Grammar instance to the builder
func (app *nodeTreeAdapterBuilder) WithGrammar(grammar grammar.Grammar) NodeTreeAdapterBuilder {
	app.grammar = grammar
	return app
}

// Now builds a new NodeTreeAdapter instance
func (app *nodeTreeAdapterBuilder) Now() (NodeTreeAdapter, error) {
	if app.grammar == nil {
		return nil, errors.New("the Grammar is mandatory in order to build a NodeTreeAdapter instance")
	}

	ruleTree, err := app.ruleTreeAdapter.ToRuleTree(app.grammar)
	if err != nil {
		return nil, err
	}

	return createNodeTreeAdapter(app.nodeTreeBuilder, ruleTree), nil
}
