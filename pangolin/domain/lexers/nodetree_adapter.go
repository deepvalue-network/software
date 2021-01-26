package lexers

import (
	"github.com/steve-care-software/products/pangolin/domain/lexers/grammar"
)

type nodeTreeAdapter struct {
	nodeTreeBuilder NodeTreeBuilder
	tree            grammar.RuleTree
}

func createNodeTreeAdapter(
	nodeTreeBuilder NodeTreeBuilder,
	tree grammar.RuleTree,
) NodeTreeAdapter {
	out := nodeTreeAdapter{
		nodeTreeBuilder: nodeTreeBuilder,
		tree:            tree,
	}

	return &out
}

// ToNodeTree converts a script to a NodeTree instance
func (app *nodeTreeAdapter) ToNodeTree(script string) (NodeTree, error) {
	return app.nodeTreeBuilder.Create().WithRuleTree(app.tree).WithScript(script).Now()
}
