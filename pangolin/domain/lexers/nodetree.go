package lexers

import (
	"strings"

	"github.com/deepvalue-network/software/pangolin/domain/lexers/grammar"
)

type nodeTree struct {
	grammar grammar.Grammar
	token   grammar.Token
	nodes   []Node
}

func createNodeTree(grammar grammar.Grammar, token grammar.Token, nodes []Node) NodeTree {
	out := nodeTree{
		grammar: grammar,
		token:   token,
		nodes:   nodes,
	}

	return &out
}

// Grammar returns the grammar
func (obj *nodeTree) Grammar() grammar.Grammar {
	return obj.grammar
}

// Token returns the token
func (obj *nodeTree) Token() grammar.Token {
	return obj.token
}

// Nodes returns the []Node
func (obj *nodeTree) Nodes() []Node {
	return obj.nodes
}

// HasSubNodeTrees returns true if rhere is sub []NodeTree
func (obj *nodeTree) HasSubNodeTrees() bool {
	for _, oneNode := range obj.nodes {
		if oneNode.HasTrees() {
			return true
		}
	}

	return false
}

// Code returns the code that the NodeTree matches
func (obj *nodeTree) Code() string {
	out := []string{}
	for _, oneNode := range obj.nodes {
		out = append(out, oneNode.Code())
	}

	return strings.Join(out, "")
}

// CodeFromName returns the code from name, the returned code is empty if the name does not exists
func (obj *nodeTree) CodeFromName(name string) string {
	codes := obj.CodesFromName(name)
	if len(codes) > 0 {
		return codes[0]
	}

	return ""
}

// CodesFromName returns the code from name
func (obj *nodeTree) CodesFromName(name string) []string {
	out := []string{}
	sub := obj.NextNodeTrees()
	for _, oneNodeTree := range sub {
		if oneNodeTree.Token().Name() == name {
			out = append(out, oneNodeTree.Code())
		}
	}

	if len(out) <= 0 {
		for _, oneNode := range obj.nodes {
			nodeCodes := oneNode.CodesFromName(name)
			if len(nodeCodes) > 0 {
				for _, oneCode := range nodeCodes {
					out = append(out, oneCode)
				}
			}
		}
	}

	return out
}

// BestMatchFromNames returns the most important code from the given names, then return the name + its code
func (obj *nodeTree) BestMatchFromNames(names []string) (string, string) {
	name := ""
	code := ""
	for _, oneName := range names {
		cde := obj.CodeFromName(oneName)
		if cde != "" {
			if len(cde) > len(code) {
				name = oneName
				code = cde
			}
		}
	}

	return name, code
}

// NextNodeTree returns the direct next nodeTree
func (obj *nodeTree) NextNodeTree() NodeTree {
	lst := obj.NextNodeTrees()
	if len(lst) > 0 {
		return lst[0]
	}

	return nil
}

// NextNodeTree returns the direct next nodeTree
func (obj *nodeTree) NextNodeTrees() []NodeTree {
	out := []NodeTree{}
	subNodes := obj.Nodes()
	for _, oneNode := range subNodes {
		if oneNode.HasTrees() {
			nodeTrees := oneNode.Trees()
			for _, oneNodeTree := range nodeTrees {
				out = append(out, oneNodeTree)
			}
		}
	}

	return out
}

// SubTreeFromName returns the subTree from name, nil if none found
func (obj *nodeTree) SubTreeFromName(name string) NodeTree {
	for _, oneNode := range obj.nodes {
		if oneNode.HasTrees() {
			subTrees := oneNode.Trees()
			for _, oneSubTree := range subTrees {
				if oneSubTree.Token().Name() == name {
					return oneSubTree
				}
			}
		}
	}

	return nil
}
