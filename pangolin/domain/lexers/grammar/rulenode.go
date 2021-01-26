package grammar

type ruleNode struct {
	rule Rule
	leaf RuleTree
}

func createRuleNodeWithRule(rule Rule) RuleNode {
	return createRuleNodeInternally(rule, nil)
}

func createRuleNodeWithLeaf(leaf RuleTree) RuleNode {
	return createRuleNodeInternally(nil, leaf)
}

func createRuleNodeInternally(rule Rule, leaf RuleTree) RuleNode {
	out := ruleNode{
		rule: rule,
		leaf: leaf,
	}

	return &out
}

// HasRule returns true if there is a Rule, false otherwise
func (obj *ruleNode) HasRule() bool {
	return obj.rule != nil
}

// Rule returns the Rule if any
func (obj *ruleNode) Rule() Rule {
	return obj.rule
}

// HasLeaf returns true if there is a leaf, false otherwise
func (obj *ruleNode) HasLeaf() bool {
	return obj.leaf != nil
}

// Leaf return the leaf if any
func (obj *ruleNode) Leaf() RuleTree {
	return obj.leaf
}

// ResetLeaf reset the RuleTree (leaf)
func (obj *ruleNode) ResetLeaf(leaf RuleTree) {
	obj.leaf = leaf
}
