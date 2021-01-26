package grammar

type repetitiveRuleNode struct {
	recursiveName       string
	node                RuleNode
	isMultipleMandatory bool
	isMultipleOptional  bool
	isOptional          bool
}

func createRepetitiveRuleNode(node RuleNode) RepetitiveRuleNode {
	return createRepetitiveRuleNodeInternally("", node, false, false, false)
}

func createRepetitiveRuleNodeRecursively(recursiveName string) RepetitiveRuleNode {
	return createRepetitiveRuleNodeInternally(recursiveName, nil, false, false, false)
}

func createRepetitiveRuleNodeWithMultipleMandatory(node RuleNode) RepetitiveRuleNode {
	return createRepetitiveRuleNodeInternally("", node, true, false, false)
}

func createRepetitiveRuleNodeRecursivelyWithMultipleMandatory(recursiveName string) RepetitiveRuleNode {
	return createRepetitiveRuleNodeInternally(recursiveName, nil, true, false, false)
}

func createRepetitiveRuleNodeWithMultipleOptional(node RuleNode) RepetitiveRuleNode {
	return createRepetitiveRuleNodeInternally("", node, false, true, false)
}

func createRepetitiveRuleNodeRecursivelyWithMultipleOptional(recursiveName string) RepetitiveRuleNode {
	return createRepetitiveRuleNodeInternally(recursiveName, nil, false, true, false)
}

func createRepetitiveRuleNodeWithOptional(node RuleNode) RepetitiveRuleNode {
	return createRepetitiveRuleNodeInternally("", node, false, false, true)
}

func createRepetitiveRuleNodeRecursivelyWithOptional(recursiveName string) RepetitiveRuleNode {
	return createRepetitiveRuleNodeInternally(recursiveName, nil, false, false, true)
}

func createRepetitiveRuleNodeInternally(
	recursiveName string,
	node RuleNode,
	isMultipleMandatory bool,
	isMultipleOptional bool,
	isOptional bool,
) RepetitiveRuleNode {
	out := repetitiveRuleNode{
		recursiveName:       recursiveName,
		node:                node,
		isMultipleMandatory: isMultipleMandatory,
		isMultipleOptional:  isMultipleOptional,
		isOptional:          isOptional,
	}

	return &out
}

// HasRecursiveName returns true if there is a recursiveName, false otherwise
func (obj *repetitiveRuleNode) HasRecursiveName() bool {
	return obj.recursiveName != ""
}

// RecursiveName returns the recursiveName, if any
func (obj *repetitiveRuleNode) RecursiveName() string {
	return obj.recursiveName
}

// HasNode returns true if there is a node, false otherwise
func (obj *repetitiveRuleNode) HasNode() bool {
	return obj.node != nil
}

// Node return the node, if not recursive
func (obj *repetitiveRuleNode) Node() RuleNode {
	return obj.node
}

// IsMultipleMandatory returns true if the RepetitiveRuleNode requires multiple mandatory
func (obj *repetitiveRuleNode) IsMultipleMandatory() bool {
	return obj.isMultipleMandatory
}

// IsMultipleOptional returns true if the RepetitiveRuleNode is multiple optional
func (obj *repetitiveRuleNode) IsMultipleOptional() bool {
	return obj.isMultipleOptional
}

// IsOptional returns true if the RepetitiveRuleNode is optional
func (obj *repetitiveRuleNode) IsOptional() bool {
	return obj.isOptional
}

// Reset sets the node:
func (obj *repetitiveRuleNode) SetNode(node RuleNode) {
	obj.node = node
}
