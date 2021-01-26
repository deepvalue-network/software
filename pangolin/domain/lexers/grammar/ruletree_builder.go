package grammar

import (
	"errors"
	"fmt"
)

type ruleTreeBuilder struct {
	ruleNodeBuilder RuleNodeBuilder
	tok             Token
	nodes           [][]RepetitiveRuleNode
	grammar         Grammar
	children        map[string]RuleTree
}

func createRuleTreeBuilder(ruleNodeBuilder RuleNodeBuilder) RuleTreeBuilder {
	out := ruleTreeBuilder{
		ruleNodeBuilder: ruleNodeBuilder,
		tok:             nil,
		nodes:           nil,
		grammar:         nil,
		children:        nil,
	}

	return &out
}

// Create initializes the builder
func (app *ruleTreeBuilder) Create() RuleTreeBuilder {
	return createRuleTreeBuilder(app.ruleNodeBuilder)
}

// WithToken adds a token to the builder
func (app *ruleTreeBuilder) WithToken(token Token) RuleTreeBuilder {
	app.tok = token
	return app
}

// WithNodes adds nodes to the builder
func (app *ruleTreeBuilder) WithNodes(nodes [][]RepetitiveRuleNode) RuleTreeBuilder {
	app.nodes = nodes
	return app
}

// WithGrammar adds a grammar to the builder
func (app *ruleTreeBuilder) WithGrammar(grammar Grammar) RuleTreeBuilder {
	app.grammar = grammar
	return app
}

// WithChildren adds a children RuleTree map to the builder
func (app *ruleTreeBuilder) WithChildren(children map[string]RuleTree) RuleTreeBuilder {
	app.children = children
	return app
}

// Now builds a new RuleTree instance
func (app *ruleTreeBuilder) Now() (RuleTree, error) {
	if app.tok == nil {
		return nil, errors.New("the Token is mandatory in order to build a RuleTree instance")
	}

	if app.nodes == nil {
		return nil, errors.New("the []RepetitiveRuleNode are mandatory in order to build a RuleTree instance")
	}

	if app.grammar == nil {
		return nil, errors.New("the Grammar is mandatory in order to build a RuleTree instance")
	}

	// create the ruleTree:
	ruleTree := createRuleTree(app.tok, app.nodes, app.grammar)
	ruleNode, err := app.ruleNodeBuilder.Create().WithLeaf(ruleTree).Now()
	if err != nil {
		return nil, err
	}

	// create stack then replace recursive:
	repRecurRuleTree, err := app.createStackThenReplaceRecursives(ruleTree, ruleNode)
	if err != nil {
		return nil, err
	}

	// replace the children, if any:
	if app.children != nil {
		return app.replaceChildren(repRecurRuleTree, app.children)
	}

	// no childre, so return:
	return repRecurRuleTree, nil
}

func (app *ruleTreeBuilder) createStackThenReplaceRecursives(ruleTree RuleTree, ruleNode RuleNode) (RuleTree, error) {
	// create the stack and add the current ruleTree:
	stack := app.createStack(ruleTree.Nodes())
	keyname := fmt.Sprintf("%s.%s", ruleTree.Grammar().Name(), ruleTree.Token().Name())
	stack[keyname] = ruleNode

	// replace the recursive RuleTree then return:
	return app.replaceRecursive(ruleTree, stack)
}

func (app *ruleTreeBuilder) replaceChildren(tree RuleTree, children map[string]RuleTree) (RuleTree, error) {
	for keyname, oneChild := range children {

		// change the name of the token:
		oneChild.Token().SetName(keyname)

		// build the ruleNode:
		childRuleNode, err := app.ruleNodeBuilder.Create().WithLeaf(oneChild).Now()
		if err != nil {
			return nil, err
		}

		repNodesMatrix := tree.Nodes()
		for i, oneRepNodeList := range repNodesMatrix {
			for j, oneRepNode := range oneRepNodeList {
				if oneRepNode.HasNode() {
					node := oneRepNode.Node()
					if node.HasLeaf() {
						leaf := node.Leaf()
						if leaf.Token().Name() == keyname {
							oneRepNode.SetNode(childRuleNode)
							tree.ResetNodeAtIndex(i, j, oneRepNode)
						}
					}
				}
			}
		}
	}

	return tree, nil
}

func (app *ruleTreeBuilder) createStack(allNodes [][]RepetitiveRuleNode) map[string]RuleNode {
	out := map[string]RuleNode{}
	for _, oneNodeList := range allNodes {
		for _, oneNode := range oneNodeList {
			if !oneNode.HasNode() {
				continue
			}

			if oneNode.HasRecursiveName() {
				continue
			}

			node := oneNode.Node()
			if !node.HasLeaf() {
				continue
			}

			// add the node:
			leaf := node.Leaf()
			keyname := fmt.Sprintf("%s.%s", leaf.Grammar().Name(), leaf.Token().Name())
			if _, ok := out[keyname]; ok {
				continue
			}

			// add the node:
			out[keyname] = node

			// add the sub nodes:
			leafNodes := node.Leaf().Nodes()
			sub := app.createStack(leafNodes)
			for subKeyname, el := range sub {
				out[subKeyname] = el
			}
		}
	}

	return out
}
func (app *ruleTreeBuilder) replaceRecursive(tree RuleTree, stack map[string]RuleNode) (RuleTree, error) {
	nodes := tree.Nodes()
	for i, oneNodeList := range nodes {
		for j, oneNode := range oneNodeList {
			if oneNode.HasRecursiveName() && oneNode.HasNode() {
				continue
			}

			var newNode RuleNode
			if oneNode.HasNode() {
				node := oneNode.Node()
				if node.HasLeaf() {
					leaf := node.Leaf()
					newLeaf, err := app.replaceRecursive(leaf, stack)
					if err != nil {
						return nil, err
					}

					// reset the leaf:
					node.ResetLeaf(newLeaf)
					newNode = node
				}
			}

			if oneNode.HasRecursiveName() {
				recursiveName := oneNode.RecursiveName()
				if stackNode, ok := stack[recursiveName]; ok {
					newNode = stackNode
				}
			}

			if newNode != nil {
				// reset the RuleNode in the RepetitiveRuleNode:
				oneNode.SetNode(newNode)

				// change the node in the tree:
				err := tree.ResetNodeAtIndex(i, j, oneNode)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return tree, nil
}
