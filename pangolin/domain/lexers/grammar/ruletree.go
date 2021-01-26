package grammar

import (
	"errors"
	"fmt"
)

type ruleTree struct {
	tok     Token
	nodes   [][]RepetitiveRuleNode
	grammar Grammar
}

func createRuleTree(tok Token, nodes [][]RepetitiveRuleNode, grammar Grammar) RuleTree {
	out := ruleTree{
		tok:     tok,
		nodes:   nodes,
		grammar: grammar,
	}

	return &out
}

// Token returns the token
func (obj *ruleTree) Token() Token {
	return obj.tok
}

// Nodes returns the nodes
func (obj *ruleTree) Nodes() [][]RepetitiveRuleNode {
	return obj.nodes
}

// Grammar returns the grammar
func (obj *ruleTree) Grammar() Grammar {
	return obj.grammar
}

// ResetNodeAtIndex reset the node at index
func (obj *ruleTree) ResetNodeAtIndex(i int, j int, newNode RepetitiveRuleNode) error {
	if i < 0 {
		return errors.New("the i cannot be smaller than 0")
	}

	if j < 0 {
		return errors.New("the j cannot be smaller than 0")
	}

	length := len(obj.nodes)
	if i >= length {
		str := fmt.Sprintf("the i (%d) must be smaller than the []RepetitiveRuleNode length (%d)", i, length)
		return errors.New(str)
	}

	secondLength := len(obj.nodes[i])
	if j >= secondLength {
		str := fmt.Sprintf("the j (%d) must be smaller than the []RepetitiveRuleNode length (%d) at index (%d)", j, secondLength, i)
		return errors.New(str)
	}

	obj.nodes[i][j] = newNode
	return nil
}
