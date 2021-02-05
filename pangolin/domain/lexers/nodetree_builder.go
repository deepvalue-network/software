package lexers

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/deepvalue-network/software/pangolin/domain/lexers/grammar"
)

type nodeTreeBuilder struct {
	elementBuilder ElementBuilder
	nodeBuilder    NodeBuilder
	tree           grammar.RuleTree
	script         string
}

func createNodeTreeBuilder(elementBuilder ElementBuilder, nodeBuilder NodeBuilder) NodeTreeBuilder {
	out := nodeTreeBuilder{
		elementBuilder: elementBuilder,
		nodeBuilder:    nodeBuilder,
		tree:           nil,
		script:         "",
	}

	return &out
}

// Create initializes the builder
func (app *nodeTreeBuilder) Create() NodeTreeBuilder {
	return createNodeTreeBuilder(app.elementBuilder, app.nodeBuilder)
}

// WithRuleTree add a RuleTree instances to the builder
func (app *nodeTreeBuilder) WithRuleTree(tree grammar.RuleTree) NodeTreeBuilder {
	app.tree = tree
	return app
}

// WithScript adds a script to the builder
func (app *nodeTreeBuilder) WithScript(script string) NodeTreeBuilder {
	app.script = script
	return app
}

// Now builds a new NodeTree instance
func (app *nodeTreeBuilder) Now() (NodeTree, error) {
	if app.tree == nil {
		return nil, errors.New("the RuleTree are mandatory in order to build a Node instance")
	}

	if app.script == "" {
		return nil, errors.New("the script is mandatory in ordser to build a Node instance")
	}

	tree, remaining, err := app.parseTree(app.tree, app.script, map[string]string{})
	if err != nil {
		return nil, err
	}

	remaining = strings.TrimSpace(remaining)
	if remaining != "" {
		str := fmt.Sprintf("there is a remaining script after executing the lexer: %s", remaining)
		return nil, errors.New(str)
	}

	return tree, nil
}

func (app *nodeTreeBuilder) rule(rule grammar.Rule, script string, isMultipleMandatory bool, isMultipleOptional bool, isOptional bool) ([]Element, string, error) {
	if isMultipleMandatory || isMultipleOptional {
		return app.ruleMultiple(rule, isMultipleMandatory)
	}

	element, remaining, err := app.ruleSingle(rule, script)
	if err != nil {
		return nil, "", err
	}

	if element == nil && !isOptional {
		str := fmt.Sprintf("the Rule (name: %s) was mandatory but could not match the given script: %s", rule.Name(), script)
		return nil, "", errors.New(str)
	}

	return []Element{element}, remaining, nil
}

func (app *nodeTreeBuilder) ruleSingle(rule grammar.Rule, script string) (Element, string, error) {
	// find the matches:
	code, matches, err := rule.FindFirst(script)
	if err != nil {
		return nil, "", err
	}

	// if there is no match
	if !matches {
		return nil, script, nil
	}

	// there is no error and no code was found, therefore return:
	if code == "" {
		return nil, script, nil
	}

	el, err := app.elementBuilder.Create().WithRule(rule).WithCode(code).Now()
	if err != nil {
		return nil, "", err
	}

	// remove the script token from the script:
	script = strings.Replace(script, code, "", 1)

	// return:
	return el, script, nil
}

func (app *nodeTreeBuilder) ruleMultiple(rule grammar.Rule, isMandatory bool) ([]Element, string, error) {
	return nil, "", errors.New("finish the ruleMultiple inside the NodeTreeBuilder")
}

func (app *nodeTreeBuilder) ruleTree(ruleTree grammar.RuleTree, script string, isMultipleMandatory bool, isMultipleOptional bool, isOptional bool, stack map[string]string) ([]NodeTree, string, error) {
	if isMultipleMandatory || isMultipleOptional {
		out := []NodeTree{}
		for {
			tree, remaining, err := app.parseTree(ruleTree, script, stack)
			if err != nil {
				break
			}

			// append:
			out = append(out, tree)

			// replace the script with the remaining:
			script = remaining
			if script == "" {
				break
			}
		}

		if isMultipleMandatory && len(out) <= 0 {
			return nil, "", errors.New("the RuleTree was expected to have at least 1 occurence, none found")
		}

		return out, script, nil
	}

	tree, remaining, err := app.parseTree(ruleTree, script, stack)
	if err != nil {
		if isOptional {
			return nil, "", nil
		}

		return nil, "", err
	}

	if tree == nil && !isOptional {
		return nil, "", errors.New("the RuleTree was mandatory but could not match the given script")
	}

	return []NodeTree{
		tree,
	}, remaining, nil
}

func (app *nodeTreeBuilder) parseTree(tree grammar.RuleTree, script string, stack map[string]string) (NodeTree, string, error) {
	bestLength := math.MaxInt64
	best := []Node{}
	bestRemaining := ""

	remainings := []string{}
	errs := []error{}
	nodes := tree.Nodes()
	for _, oneNodeList := range nodes {
		nodes, remaining, err := app.parseRepetitiveRuleNodes(oneNodeList, script, stack)
		if err != nil {
			remainings = append(remainings, remaining)
			errs = append(errs, err)
			continue
		}

		remainingLength := len(remaining)
		if remainingLength < bestLength {
			bestLength = remainingLength
			best = nodes
			bestRemaining = remaining
		}
	}

	if len(best) <= 0 {
		str := fmt.Sprintf("the RuleTree could not find a matching []Node on the given script.  The NodeTree contained %d  choice(s), here are the errors, script and remaining script for each: \n %v \n %v \n", len(nodes), remainings, errs)
		return nil, "", errors.New(str)
	}

	nodeTree := createNodeTree(tree.Grammar(), tree.Token(), best)
	return nodeTree, bestRemaining, nil
}

func (app *nodeTreeBuilder) parseRepetitiveRuleNodes(ruleNodes []grammar.RepetitiveRuleNode, script string, stack map[string]string) ([]Node, string, error) {
	originalScript := script
	originalStack := stack

	nodes := []Node{}
	cpt := 0
	for _, oneRepetitiveRuleNode := range ruleNodes {
		cpt++

		isMultipleMandatory := oneRepetitiveRuleNode.IsMultipleMandatory()
		isMultipleOptional := oneRepetitiveRuleNode.IsMultipleOptional()
		isOptional := oneRepetitiveRuleNode.IsOptional()

		remaining := ""
		var node Node
		builder := app.nodeBuilder.Create()
		if !oneRepetitiveRuleNode.HasNode() {
			continue
		}

		subNode := oneRepetitiveRuleNode.Node()

		// if the recursive was called previously with exactly the same code, skip it:
		if oneRepetitiveRuleNode.HasNode() && oneRepetitiveRuleNode.HasRecursiveName() {
			recursiveName := oneRepetitiveRuleNode.RecursiveName()
			if stackScript, ok := stack[recursiveName]; ok {
				if stackScript == script {
					continue
				}
			}

			stack[recursiveName] = script
		}

		if subNode.HasRule() {
			rule := subNode.Rule()
			elements, rem, err := app.rule(rule, script, isMultipleMandatory, isMultipleOptional, isOptional)
			if err != nil {
				return nil, "", err
			}

			remaining = rem
			builder.WithElements(elements)
		}

		if subNode.HasLeaf() {
			leaf := subNode.Leaf()
			nodeTrees, rem, err := app.ruleTree(leaf, script, isMultipleMandatory, isMultipleOptional, isOptional, stack)
			if err != nil {
				return nil, "", err
			}

			remaining = rem
			builder.WithTrees(nodeTrees)
		}

		// build:
		node, err := builder.Now()
		if err != nil {
			continue
		}

		// append:
		nodes = append(nodes, node)

		// replace the script with the remaining, and reset the stack:
		if script != remaining {
			script = remaining
			stack = map[string]string{}
		}

		// if there is no more script, end the loop:
		if script == "" {
			break
		}
	}

	if cpt < len(ruleNodes) {
		remainingRuleNodes := ruleNodes[cpt:]
		if app.containsMandatory(remainingRuleNodes) {
			newRuleNodes := app.removeFirstOptionalIfPossible(ruleNodes)
			if len(newRuleNodes) == len(ruleNodes) {
				return nil, "", errors.New("the Grammar contained mandatory elements but the script was empty, and no optional prefix token could be removed to make it pass")
			}

			return app.parseRepetitiveRuleNodes(newRuleNodes, originalScript, originalStack)
		}
	}

	// return
	return nodes, script, nil
}

func (app *nodeTreeBuilder) containsMandatory(ruleNodes []grammar.RepetitiveRuleNode) bool {
	for _, oneRuleNode := range ruleNodes {
		if oneRuleNode.IsMultipleOptional() {
			continue
		}

		if !oneRuleNode.IsOptional() || (oneRuleNode.IsMultipleMandatory()) {
			return true
		}
	}

	return false
}

func (app *nodeTreeBuilder) removeFirstOptionalIfPossible(ruleNodes []grammar.RepetitiveRuleNode) []grammar.RepetitiveRuleNode {
	isFirstRemoved := false
	out := []grammar.RepetitiveRuleNode{}
	for _, oneRuleNode := range ruleNodes {
		if !isFirstRemoved {
			if oneRuleNode.IsMultipleOptional() || oneRuleNode.IsOptional() {
				isFirstRemoved = true
				continue
			}
		}

		out = append(out, oneRuleNode)
	}

	return out
}
